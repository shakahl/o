# o [![Build Status](https://travis-ci.com/xyproto/o.svg?branch=master)](https://travis-ci.com/xyproto/o) [![Go Report Card](https://goreportcard.com/badge/github.com/xyproto/o)](https://goreportcard.com/report/github.com/xyproto/o) [![License](https://img.shields.io/badge/license-BSD-green.svg?style=flat)](https://raw.githubusercontent.com/xyproto/o/master/LICENSE)

`o` is a small and fast, but limited text editor.

* Compiles with either `go` or `gccgo`.
* Tested with `st`, `urxvt` and `xfce4-terminal`.
* Tested on Arch Linux and FreeBSD.

For a more feature complete editor that is also written in Go, check out [micro](https://github.com/zyedidia/micro).

<!--## Screenshot

![screenshot](img/screenshot.png)
-->

## Quick start

You can install `o` with Go 1.10 or later (development version):

    go get -u github.com/xyproto/o

## Features and limitations

* Has syntax highlighting for Go code.
* Never asks before saving or quitting. Be careful!
* `Home` and `End` are not detected by the key handler. `ctrl-a` and `ctrl-e` works, though.
* Can format Go or C++ code, just press `ctrl-w`. This uses either `goimports` (`go get golang.org/x/tools/cmd/goimports`) or `clang-format`.
* Will strip trailing whitespace whenever it can.
* Must be given a filename at start.
* Requires that `/dev/tty` is available.
* Copy, cut and paste is only for one line at a time, and only within the editor.
* Some letters can not be typed in. Like `æ`.
* May take a line number as the second argument, with an optional `+` prefix.
* The text will be red if a loaded file is read-only.
* The terminal needs to be resized to show the second half of lines that are longer than the terminal width.
* If the filename is `COMMIT_EDITMSG`, the look and feel will be adjusted for git commit messages.

## Known bugs

* Find and GoTo line may jump to the wrong locations.

## Spinner

When loading large files, an animated spinner will appear. The loading operation can be stopped at any time by pressing `esc`, `q` or `ctrl-q`.

![progress](img/progress.gif)

## Hotkeys

* `ctrl-q` - quit
* `ctrl-s` - save
* `ctrl-w` - format the current file with `go fmt`
* `ctrl-a` - go to start of line, then start of text on the same line
* `ctrl-e` - go to end of line
* `ctrl-p` - scroll up 10 lines
* `ctrl-n` - scroll down 10 lines
* `ctrl-k` - delete characters to the end of the line, then delete the line
* `ctrl-g` - toggle filename/line/column/unicode/word count status display
* `ctrl-d` - delete a single character
* `ctrl-t` - toggle syntax highlighting
* `ctrl-r` - toggle text or draw mode (for ASCII graphics)
* `ctrl-x` - cut the current line
* `ctrl-c` - copy the current line
* `ctrl-v` - paste the current line
* `ctrl-b` - bookmark the current position
* `ctrl-j` - jump to the bookmark
* `ctrl-h` - show a minimal help text
* `ctrl-u` - undo
* `ctrl-l` - jump to a specific line
* `ctrl-f` - find a string. Press ctrl-f and return to repeat the search.
* `esc` - redraw the screen and clear the last search.

## Size

The `o` executable is only **464k** when built with GCC 9.1 (for 64-bit Linux). This isn't as small as [e3](https://sites.google.com/site/e3editor/), an editor written in assembly (which is **234k**), but it's resonably lean. :sunglasses:

    go build -gccgoflags '-Os -s'

For comparison, it's **2.9M** when building with Go 1.13 and no particular build flags are given.

## Jumping to a specific line when opening a file

These four ways of opening `file.txt` at line `7` are supported:

* `o file.txt 7`
* `o file.txt +7`
* `o file.txt:7`
* `o file.txt+7`

This also means that filenames containing `+` or `:` are not supported.

## General info

* Version: 2.5.2
* License: 3-clause BSD
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
