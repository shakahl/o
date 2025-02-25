// main is the main package for the o editor
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"syscall"

	"github.com/xyproto/vt100"
)

const (
	versionString = "o 2.59.0"
)

func main() {
	var (
		versionFlag = flag.Bool("version", false, "version information")
		helpFlag    = flag.Bool("help", false, "quick overview of hotkeys")
		forceFlag   = flag.Bool("f", false, "open even if already open")
	)

	flag.Parse()

	if *versionFlag {
		fmt.Println(versionString)
		return
	}

	if *helpFlag {
		fmt.Println(versionString + " - simple and limited text editor")
		fmt.Print(`
Hotkeys

ctrl-s     to save
ctrl-q     to quit
ctrl-o     to open the command menu
ctrl-r     to open a portal so that text can be pasted into another file with ctrl-v
ctrl-space to compile programs, render MD to PDF or export adoc/sdoc as man
ctrl-w     for Zig, Rust, V and Go, format with the "... fmt" command
           for C++, format the current file with "clang-format"
           for HTML, format the file with "tidy", for Python: "autopep8"
           for Markdown, toggle checkboxes
           for git interactive rebases, cycle the rebase keywords
ctrl-g     to toggle filename/line/column/unicode/word count status display
ctrl-a     go to start of line, then start of text and then the previous line
ctrl-e     go to end of line and then the next line
ctrl-n     to scroll down 10 lines or go to the next match if a search is active
ctrl-p     to scroll up 10 lines or go to the previous match
ctrl-k     to delete characters to the end of the line, then delete the line
ctrl-j     to join lines
ctrl-d     to delete a single character
ctrl-t     for C and C++, toggle between the header and implementation,
           for Agda, insert a symbol,
           for the rest, record and then play back a macro
ctrl-c     to copy the current line, press twice to copy the current block
ctrl-v     to paste one line, press twice to paste the rest
ctrl-x     to cut the current line, press twice to cut the current block
ctrl-b     to toggle a bookmark for the current line, or jump to a bookmark
ctrl-u     to undo (ctrl-z is also possible, but may background the application)
ctrl-l     to jump to a specific line (press return to jump to the top or bottom)
ctrl-f     to find a string, press Tab after the text to search and replace
ctrl-\     to toggle single-line comments for a block of code
ctrl-~     to jump to matching parenthesis
esc        to redraw the screen and clear the last search

Set NO_COLOR=1 to disable colors.

See the man page for more information.

`)
		return
	}

	traceStart() // if building with -tags trace

	// Check if the executable starts with "g" or "f"
	var executableName string
	if len(os.Args) > 0 {
		executableName = filepath.Base(os.Args[0]) // if os.Args[0] is empty, executableName will be "."
		switch executableName[0] {
		case 'f', 'g':
			// Start the game
			if _, err := Game(); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			return
		}
	}

	var (
		err        error
		fnord      FilenameOrData
		lineNumber LineNumber
		colNumber  ColNumber
	)

	stdinFilename := len(os.Args) == 1 || (len(os.Args) == 2 && (os.Args[1] == "-" || os.Args[1] == "/dev/stdin"))
	// If no regular filename is given, check if data is ready at stdin
	readFromStdin := stdinFilename && dataReadyOnStdin()
	if readFromStdin {
		// TODO: Use a spinner?
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, "could not read from stdin")
			os.Exit(1)
		}
		// Now stop reading further from stdin
		os.Stdin.Close()
		if lendata := len(data); lendata > 0 {
			fnord.filename = "-"
			fnord.data = data
			fnord.length = uint64(lendata)
		}
	} else {
		fnord.filename, lineNumber, colNumber = FilenameAndLineNumberAndColNumber(flag.Arg(0), flag.Arg(1), flag.Arg(2))
	}

	// Check if the given filename contains something
	if fnord.Empty() {
		if fnord.filename == "" {
			fmt.Fprintln(os.Stderr, "please provide a filename")
			os.Exit(1)
		}

		// If the filename starts with "~", then expand it
		fnord.ExpandUser()

		// Check if the given filename exists
		if !exists(fnord.filename) {
			if strings.HasSuffix(fnord.filename, ".") {
				// If the filename ends with "." and the file does not exist, assume this was a result of tab-completion going wrong.
				// If there are multiple files that exist that start with the given filename, open the one first in the alphabet (.cpp before .o)
				matches, err := filepath.Glob(fnord.filename + "*")
				if err == nil && len(matches) > 0 { // no error and at least 1 match
					// Use the first non-binary match of the sorted results
					matches = removeBinaryFiles(matches)
					if len(matches) > 0 {
						sort.Strings(matches)
						fnord.filename = matches[0]
					}
				}
			} else if !strings.Contains(fnord.filename, ".") && allLower(fnord.filename) {
				// The filename has no ".", is written in lowercase and it does not exist,
				// but more than one file that starts with the filename  exists. Assume tab-completion failed.
				matches, err := filepath.Glob(fnord.filename + "*")
				if err == nil && len(matches) > 1 { // no error and more than 1 match
					// Use the first non-binary match of the sorted results
					matches = removeBinaryFiles(matches)
					if len(matches) > 0 {
						sort.Strings(matches)
						fnord.filename = matches[0]
					}
				}
			} else {
				// Also match "PKGBUILD" if just "Pk" was entered
				matches, err := filepath.Glob(strings.ToTitle(fnord.filename) + "*")
				if err == nil && len(matches) >= 1 { // no error and at least 1 match
					// Use the first non-binary match of the sorted results
					matches = removeBinaryFiles(matches)
					if len(matches) > 0 {
						sort.Strings(matches)
						fnord.filename = matches[0]
					}
				}
			}
		}
	}

	// Set the terminal title, if the current terminal emulator supports it, and NO_COLOR is not set
	fnord.SetTitle()

	// If the editor executable has been named "red", use the red/gray theme by default
	// Also use the red/gray theme if $SHELL is /bin/csh (typically BSD)
	theme := NewDefaultTheme()
	syntaxHighlight := true
	if envNoColor {
		theme = NewNoColorDarkBackgroundTheme()
		syntaxHighlight = false
	} else {
		// Check if the executable starts with a specific letter
		if len(executableName) > 0 {
			switch executableName[0] {
			case 'b', 'e': // bo, borland, ed, edit etc.
				theme = NewDarkBlueEditTheme()
				// TODO: Later, when specificLetter is examined, use either NewEditLightTheme or NewEditDarkTheme
				specificLetter = true
				editTheme = true
			case 'l', 'v': // lo, light, vs, vscode etc.
				theme = NewDarkVSTheme()
				specificLetter = true
			case 'r': // rb, ro, rt, red etc.
				theme = NewRedBlackTheme()
				specificLetter = true
			case 's': // s, sw, synthwave etc.
				theme = NewSynthwaveTheme()
				specificLetter = true
			}
		}
	}

	// Initialize the VT100 terminal
	tty, err := vt100.NewTTY()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: "+err.Error())
		os.Exit(1)
	}
	defer tty.Close()

	// Run the main editor loop
	userMessage, stopParent, err := Loop(tty, fnord, lineNumber, colNumber, *forceFlag, theme, syntaxHighlight)

	// SIGQUIT the parent PID. Useful if being opened repeatedly by a find command.
	if stopParent {
		defer func() {
			syscall.Kill(os.Getppid(), syscall.SIGQUIT)
		}()
	}

	// Remove the terminal title, if the current terminal emulator supports it
	// and if NO_COLOR is not set.
	NoTitle()

	// Clear the current color attribute
	fmt.Print(vt100.Stop())

	// Respond to the error returned from the main loop, if any
	if err != nil {
		if userMessage != "" {
			quitMessage(tty, userMessage)
		} else {
			quitError(tty, err)
		}
	}

	traceComplete() // if building with -tags trace
}
