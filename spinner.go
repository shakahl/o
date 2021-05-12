package main

import (
	"fmt"
	"time"

	"github.com/xyproto/textoutput"
	"github.com/xyproto/vt100"
)

var pacmanNoColor = []string{
	"| C · · |",
	"|  C· · |",
	"|   C · |",
	"|    C· |",
	"|     C |",
	"|      C|",
	"| · · Ɔ |",
	"| · ·Ɔ  |",
	"| · Ɔ   |",
	"| ·Ɔ    |",
	"| Ɔ     |",
	"|Ɔ· · · |",
}

var pacmanColor = []string{
	"<red>| <yellow>C<blue> · ·</blue> <red>|<off>",
	"<red>| <blue> <yellow>C<blue>· · <red>|<off>",
	"<red>| <blue>  <yellow>C<blue> · <red>|<off>",
	"<red>| <blue>   <yellow>C<blue>· <red>|<off>",
	"<red>| <blue>    <yellow>C <red>|<off>",
	"<red>| <blue>     <yellow>C<red>|<off>",
	"<red>| <blue>· · <yellow>Ɔ <red>|<off>",
	"<red>| <blue>· ·<yellow>Ɔ<blue>  <red>|<off>",
	"<red>| <blue>· <yellow>Ɔ <blue>  <red>|<off>",
	"<red>| <blue>·<yellow>Ɔ<blue>    <red>|<off>",
	"<red>| <yellow>Ɔ <blue>    <red>|<off>",
	"<red>|<yellow>Ɔ<blue>· · · <red>|<off>",
}

// Spinner waits a bit, then displays a spinner together with the given message string (msg).
// If the spinner is aborted, the qmsg string is displayed.
// Returns a quit channel (chan bool).
// The spinner is shown asynchronously.
// "true" must be sent to the quit channel once whatever operating that the spinner is spinning for is completed.
func Spinner(c *vt100.Canvas, tty *vt100.TTY, umsg, qmsg string, noColor bool) chan bool {
	quitChan := make(chan bool)
	go func() {
		// Wait 4 * 4 milliseconds, while listening to the quit channel.
		// This is to delay showing the progress bar until some time has passed.
		for i := 0; i < 4; i++ {
			// Check if we should quit or wait
			select {
			case <-quitChan:
				return
			default:
				// Wait a tiny bit
				time.Sleep(4 * time.Millisecond)
			}
		}

		// If c or tty are nil, use the silent spinner
		if (c == nil) || (tty == nil) {
			// Wait for a true on the quit channel, then return
			<-quitChan
			return
		}

		var (
			// Find a good start location
			x = uint(int(c.Width()) / 7)
			y = uint(int(c.Height()) / 7)

			// Get the terminal codes for coloring the given user message the same as italics in Markdown
			msg = italicsColor.Get(umsg)
		)

		// Move the cursor there and write a message
		vt100.SetXY(x, y)
		fmt.Print(msg)

		// Store the position after the message
		x += uint(len(msg)) + 1

		// Prepare to output colored text
		var (
			o                = textoutput.NewTextOutput(true, true)
			counter          uint
			spinnerAnimation []string
		)

		// Hide the cursor
		vt100.ShowCursor(false)
		defer vt100.ShowCursor(true)

		if noColor {
			spinnerAnimation = pacmanNoColor
		} else {
			spinnerAnimation = pacmanColor
		}

		// Start the spinner
		for {
			select {
			case <-quitChan:
				return
			default:
				vt100.SetXY(x, y)
				// Iterate over the 12 different ASCII images as the counter increases
				o.Print(spinnerAnimation[counter%12])
				counter++
				// Wait for a key press (also sleeps just a bit)
				switch tty.Key() {
				case 27, 113, 17, 3: // esc, q, ctrl-q or ctrl-c
					quitMessage(tty, qmsg)
				}
			}

		}
	}()
	return quitChan
}
