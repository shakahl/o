package main

import (
	"strings"
	"unicode"

	"github.com/xyproto/vt100"
)

func handleManPageEscape(s string) string {
	var lineRunes []rune
	for _, r := range s {
		if r == 0x8 {
			// Encountered ^H
			// Pop the last appended rune and continue
			lineRunes = lineRunes[:len(lineRunes)-1]
			continue
		}
		lineRunes = append(lineRunes, r)
	}
	return string(lineRunes)
}

func (e *Editor) manPageHighlight(line, programName string, firstLine, lastLine bool) string {
	line = handleManPageEscape(line)
	var (
		normal          = e.Foreground
		off             = vt100.Stop()
		trimmedLine     = strings.TrimSpace(line)
		hasWords        = HasWords(trimmedLine)
		innerSpaceCount = strings.Count(trimmedLine, " ")
	)

	if strings.Count(trimmedLine, "  ") > 10 && (firstLine || lastLine) { // first and last line
		return e.CommentColor.Get(line)
	} else if strings.ToUpper(trimmedLine) == trimmedLine && !strings.HasPrefix(trimmedLine, "-") && hasWords && !strings.HasPrefix(line, " ") { // a sub-section header
		return e.ManSectionColor.Get(line)
	} else if strings.HasPrefix(trimmedLine, "-") { // a flag or parameter
		var rs []rune
		rs = append(rs, []rune(e.MarkdownTextColor.String())...)
		inFlag := false
		spaceCount := 0
		foundLetter := false
		prevR := ' '
		for _, r := range line {
			if strings.HasPrefix(trimmedLine, "-") && strings.Count(line, "-") >= 1 && strings.Count(trimmedLine, " ") <= 1 {
				// One or two command line options, color them differently
				return e.MenuArrowColor.Get(line)
			}

			if !foundLetter && (unicode.IsLetter(r) || r == '_') {
				foundLetter = true
			}
			if r == ' ' {
				spaceCount++
				if innerSpaceCount > 8 {
					inFlag = false
				}
			} else {
				spaceCount = 0
			}

			if r != ' ' && (prevR == ' ' || prevR == '-') && (r == '-' || r == '[' || r == '_') && (prevR == '-' || !inFlag) {
				inFlag = true
				rs = append(rs, []rune(off+e.MenuArrowColor.String())...)
				rs = append(rs, r)
			} else if (prevR == ' ' || prevR == '-') && (r == '-' || r == '[' || r == ']' || r == '_') && inFlag {
				rs = append(rs, r)
			} else if inFlag { // Color the rest of the flag text in the textColor color (LightBlue)
				inFlag = false
				rs = append(rs, []rune(off+e.MarkdownTextColor.String())...)
				rs = append(rs, r)
			} else if foundLetter && spaceCount > 2 { // Color the rest of the line in the foreground color (LightGreen)
				rs = append(rs, []rune(off+normal.String())...)
				rs = append(rs, r)
			} else if r == ']' { // Color the rest of the line in the comment color (DarkGray)
				rs = append(rs, []rune(off+e.CommentColor.String())...)
				rs = append(rs, r)
			} else {
				rs = append(rs, r)
			}
			prevR = r
		}
		rs = append(rs, []rune(off)...)
		return string(rs)
	} else if allUpper(trimmedLine) {
		return e.MarkdownTextColor.Get(line)
	} else { // regular text, but highlight numbers (and hex numbers, if the number starts with a digit) + highlight "@"
		var (
			rs       []rune
			prevRune rune
			inDigits bool
			inWord   bool
			inAngles bool
			nextRune rune
		)

		rs = append(rs, []rune(normal.String())...)
		hasAlpha := strings.Contains(trimmedLine, "@")
		lineRunes := []rune(line)

		for i, r := range line {

			if (i + 1) < len(lineRunes) {
				nextRune = lineRunes[i+1]
			} else {
				nextRune = ' '
			}

			if (unicode.IsLetter(r) || r == '_') && !inWord {
				inWord = true
			} else if inWord && !unicode.IsLetter(r) && !hexDigit(r) {
				inWord = false
			}
			if !inAngles && r == '<' {
				inAngles = true
			} else if inAngles && r == '>' {
				inAngles = false
			}
			if !inWord && unicode.IsDigit(r) && !inDigits {
				inDigits = true
				rs = append(rs, []rune(off+e.ItalicsColor.String())...)
				rs = append(rs, r)
			} else if inDigits && hexDigit(r) {
				rs = append(rs, r)
			} else if !inWord && inDigits {
				inDigits = false
				rs = append(rs, []rune(off+normal.String())...)
				rs = append(rs, r)
			} else if !inWord && (r == '*' || r == '$' || r == '%' || r == '!' || r == '/' || r == '=' || r == '-') {
				rs = append(rs, []rune(off+e.MenuArrowColor.String())...)
				rs = append(rs, r)
			} else if r == '@' { // color @ gray and the rest of the string white
				rs = append(rs, []rune(off+e.CommentColor.String())...)
				rs = append(rs, r)
				rs = append(rs, []rune(off+e.ItalicsColor.String())...)
			} else if hasAlpha && r == '<' { // color < gray and the rest of the string white
				rs = append(rs, []rune(off+e.CommentColor.String())...)
				rs = append(rs, r)
				rs = append(rs, []rune(off+e.ItalicsColor.String())...)
			} else if hasAlpha && r == '>' { // color > gray and the rest of the string normal
				rs = append(rs, []rune(off+e.CommentColor.String())...)
				rs = append(rs, r)
				rs = append(rs, []rune(off+normal.String())...)
			} else if inAngles || r == '>' {
				rs = append(rs, []rune(off+e.ItalicsColor.String())...)
				rs = append(rs, r)
			} else if inWord && unicode.IsUpper(prevRune) && ((unicode.IsUpper(r) && unicode.IsLetter(nextRune)) || (unicode.IsLower(r) && unicode.IsUpper(prevRune) && !unicode.IsLetter(nextRune))) {
				if unicode.IsUpper(r) {
					// This is for the leading and trailing letter of uppercase words
					rs = append(rs, []rune(off+e.ImageColor.String())...)
				} else {
					rs = append(rs, []rune(off+e.MarkdownTextColor.String())...)
				}
				rs = append(rs, r)
			} else if inWord && (unicode.IsUpper(r) || (unicode.IsUpper(prevRune) && unicode.IsLetter(r))) {
				if !unicode.IsLower(r) && (((unicode.IsUpper(nextRune) || nextRune == ' ') && unicode.IsLetter(prevRune)) || unicode.IsUpper(nextRune) || !unicode.IsLetter(nextRune)) {
					// This is for the center letters of uppercase words
					rs = append(rs, []rune(off+e.ImageColor.String())...)
				} else {
					rs = append(rs, []rune(off+e.MarkdownTextColor.String())...)
				}
				rs = append(rs, r)
			} else if inWord && unicode.IsUpper(r) {
				rs = append(rs, []rune(off+e.ImageColor.String())...)
				rs = append(rs, r)
			} else if !inWord || !unicode.IsUpper(r) {
				rs = append(rs, []rune(off+e.MarkdownTextColor.String())...)
				rs = append(rs, r)
			} else {
				rs = append(rs, r)
			}
			prevRune = r
		}
		rs = append(rs, []rune(off)...)
		return string(rs)
	}
}
