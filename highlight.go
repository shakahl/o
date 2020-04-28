package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/xyproto/syntax"
	"github.com/xyproto/textoutput"
	"github.com/xyproto/vt100"
)

// SingleLineComment returns true if the given trimmed line is a single line comment (and not in a multi line comment or in a multi line string)
func SingleLineComment(trimmedLine, singleLineCommentMarker string, inMultiLineComment, inMultiLineString bool) bool {
	return (!inMultiLineComment && !inMultiLineString) && (strings.HasPrefix(trimmedLine, singleLineCommentMarker) || (strings.HasPrefix(trimmedLine, "/*") && strings.HasSuffix(trimmedLine, "*/")))
}

// MultiLineCommentStart returns true if the given trimmed line is the start of a multi line comment
func MultiLineCommentStart(trimmedLine string, inMultiLineComment, inMultiLineString, inSingleLineComment bool) bool {
	return (!inMultiLineComment && !inMultiLineString && !inSingleLineComment) && strings.Contains(trimmedLine, "/*")
}

// MultiLineCommentStop returns true if the given trimmed line is the stop of a multi line comment
func MultiLineCommentStop(trimmedLine string, inMultiLineComment, inMultiLineString, inSingleLineComment bool) bool {
	return (inMultiLineComment && !inMultiLineString && !inSingleLineComment) && strings.Contains(trimmedLine, "*/")
}

// MultiLineStringStart returns true if the given trimmed line is the assumed start of a multi line string
func MultiLineStringStart(trimmedLine string, inMultiLineComment, inMultiLineString, inSingleLineComment bool) bool {
	return (!inMultiLineComment && !inMultiLineString && !inSingleLineComment) && ((strings.Count(trimmedLine, "`") % 2) != 0)
}

// MultiLineStringStop returns true if the given trimmed line is the assumed stop of a multi line string
func MultiLineStringStop(trimmedLine string, inMultiLineComment, inMultiLineString, inSingleLineComment bool) bool {
	return (!inMultiLineComment && inMultiLineString && !inSingleLineComment) && ((strings.Count(trimmedLine, "`") % 2) != 0)
}

// WriteLines will draw editor lines from "fromline" to and up to "toline" to the canvas, at cx, cy
func (e *Editor) WriteLines(c *vt100.Canvas, fromline, toline, cx, cy int) error {
	o := textoutput.NewTextOutput(true, true)
	tabString := " "
	if !e.DrawMode() {
		if e.mode == modeMakefile {
			tabString = strings.Repeat(tabCharacter, e.spacesPerTab)
		} else {
			tabString = strings.Repeat(" ", e.spacesPerTab)
		}
	}
	w := int(c.Width())
	if fromline >= toline {
		return errors.New("fromline >= toline in WriteLines")
	}
	numlines := toline - fromline
	offset := fromline
	inCodeBlock := false // used when highlighting Markdown
	// If in Markdown mode, figure out the current state of block quotes
	if e.mode == modeMarkdown {
		// Figure out if "fromline" is within a markdown code block or not
		for i := 0; i < fromline; i++ {
			// Check if the untrimmed line starts with ~~~ or ```
			contents := e.Line(i)
			if strings.HasPrefix(contents, "~~~") || strings.HasPrefix(contents, "```") {
				// Toggle the flag for if we're in a code block or not
				inCodeBlock = !inCodeBlock
			}
		}
	}
	var (
		noColor                 bool = os.Getenv("NO_COLOR") != ""
		inMultiLineComment      bool // used when highlighting C, Go, C++ etc (using /* and */)
		inMultiLineString       bool // used for multiline strings in Go
		singleLineCommentMarker = e.SingleLineCommentMarker()
		trimmedLine             string
	)
	// First loop from 0 to offset to figure out if we are already in a multiline comment or a multiline string at the current line
	for i := 0; i < offset; i++ {
		trimmedLine = strings.TrimSpace(e.Line(i))
		switch {
		case SingleLineComment(trimmedLine, singleLineCommentMarker, inMultiLineComment, inMultiLineString):
			// A single line comment (including /* ... */)
		case MultiLineCommentStart(trimmedLine, inMultiLineComment, inMultiLineString, false):
			// Multi line comment start
			inMultiLineComment = true
		case MultiLineCommentStop(trimmedLine, inMultiLineComment, inMultiLineString, false):
			// Multi line comment stop
			inMultiLineComment = false
		case MultiLineStringStart(trimmedLine, inMultiLineComment, inMultiLineString, false):
			inMultiLineString = true
		case MultiLineStringStop(trimmedLine, inMultiLineComment, inMultiLineString, false):
			inMultiLineString = false
		}
	}
	var (
		counter                     int
		line                        string
		screenLine                  string
		y                           int
		inSingleLineComment         bool
		newlyStartedMultiLineString bool
		assemblyMode                = e.mode == modeAssembly
	)
	// Then loop from 0 to numlines (used as y+offset in the loop) to draw the text
	for y = 0; y < numlines; y++ {
		counter = 0
		newlyStartedMultiLineString = false
		line = strings.Replace(e.Line(y+offset), "\t", tabString, -1)
		screenLine = strings.TrimRightFunc(line, unicode.IsSpace)
		if len([]rune(screenLine)) >= w {
			screenLine = screenLine[:w]
		}
		if e.syntaxHighlight && !noColor {
			// Output a syntax highlighted line. Escape any tags in the input line.
			// textWithTags must be unescaped if there is not an error.
			if textWithTags, err := syntax.AsText([]byte(Escape(line)), assemblyMode); err != nil {
				// Only output the line up to the width of the canvas
				fmt.Println(screenLine)
				counter += len([]rune(screenLine))
			} else {
				// Color and unescape
				var coloredString string
				if e.mode == modeGit {
					coloredString = e.gitHighlight(line)
				} else if e.mode == modeMarkdown {
					if highlighted, ok, codeBlockFound := markdownHighlight(line, inCodeBlock); ok {
						coloredString = highlighted
						if codeBlockFound {
							inCodeBlock = !inCodeBlock
						}
					} else {
						// Syntax highlight the line if it's not picked up by the markdownHighlight function
						coloredString = UnEscape(o.DarkTags(string(textWithTags)))
					}
				} else {
					trimmedLine = strings.TrimSpace(line)
					inSingleLineComment = false
					if SingleLineComment(trimmedLine, singleLineCommentMarker, inMultiLineComment, inMultiLineString) {
						// A single line comment (including /* ... */)
						inSingleLineComment = true
					}
					if MultiLineCommentStart(trimmedLine, inMultiLineComment, inMultiLineString, inSingleLineComment) {
						// Multi line comment start
						inMultiLineComment = true
					} else if MultiLineCommentStop(trimmedLine, inMultiLineComment, inMultiLineString, inSingleLineComment) {
						// Multi line comment stop
						inMultiLineComment = false
					}
					if MultiLineStringStop(trimmedLine, inMultiLineComment, inMultiLineString, inSingleLineComment) {
						// Multi line string stop
						inMultiLineString = false
					} else if MultiLineStringStart(trimmedLine, inMultiLineComment, inMultiLineString, inSingleLineComment) {
						// Multi line string start
						inMultiLineString = true
						newlyStartedMultiLineString = true
					}

					// TODO: Multiline string start and stop at the same line

					switch {
					case inSingleLineComment:
						// A single line comment (the syntax module did the highlighting)
						coloredString = UnEscape(o.DarkTags(string(textWithTags)))
					case inMultiLineComment:
						// A multi line comment
						coloredString = UnEscape(e.multilineComment.Get(line))
					case !newlyStartedMultiLineString && inMultiLineString:
						// A multi line string
						coloredString = UnEscape(e.multilineString.Get(line))
					default:
						// Regular code
						coloredString = UnEscape(o.DarkTags(string(textWithTags)))
					}

				}

				// Slice of runes and color attributes, while at the same time highlighting search terms
				charactersAndAttributes := o.Extract(coloredString)
				searchTermRunes := []rune(e.searchTerm)
				matchForAnotherN := 0
				for characterIndex, ca := range charactersAndAttributes {
					letter := ca.R
					fg := ca.A
					if letter == ' ' {
						fg = e.fg
					}
					if matchForAnotherN > 0 {
						// Coloring an already found match
						fg = e.searchFg
						matchForAnotherN--
					} else if len(e.searchTerm) > 0 && letter == searchTermRunes[0] {
						// Potential search highlight match
						length := len([]rune(e.searchTerm))
						counter := 0
						match := true
						for i := characterIndex; i < (characterIndex + length); i++ {
							if i >= len(charactersAndAttributes) {
								match = false
								break
							}
							ca2 := charactersAndAttributes[i]
							if ca2.R != []rune(e.searchTerm)[counter] {
								// mismatch, not a hit
								match = false
								break
							}
							counter++
						}
						// match?
						if match {
							fg = e.searchFg
							matchForAnotherN = length - 1
						}
					}
					if letter == '\t' {
						c.Write(uint(cx+counter), uint(cy+y), fg, e.bg, tabString)
						if e.DrawMode() {
							counter++
						} else {
							counter += e.spacesPerTab
						}
					} else {
						c.WriteRune(uint(cx+counter), uint(cy+y), fg, e.bg, letter)
						counter++
					}
				}
			}
		} else {
			// Output a regular line
			c.Write(uint(cx+counter), uint(cy+y), e.fg, e.bg, screenLine)
			counter += len([]rune(screenLine))
		}
		//length := len([]rune(screenLine)) + strings.Count(screenLine, "\t")*(e.spacesPerTab-1)
		// Fill the rest of the line on the canvas with "blanks"
		for x := counter; x < w; x++ {
			c.WriteRune(uint(cx+x), uint(cy+y), e.fg, e.bg, ' ')
		}
	}
	return nil
}
