package main

import (
	"errors"

	"github.com/xyproto/vt100"
)

// Position represents a position on the screen, including how far down the view has scrolled
type Position struct {
	sx          int // the position of the cursor in the current scrollview
	sy          int // the position of the cursor in the current scrollview
	offsetX     int // how far one has scrolled along the X axis
	offsetY     int // how far one has scrolled along the Y axis
	scrollSpeed int // how many lines to scroll, when scrolling up and down
	savedX      int // for smart down cursor movement
}

// NewPosition returns a new Position struct
func NewPosition(scrollSpeed int) *Position {
	return &Position{0, 0, 0, 0, scrollSpeed, 0}
}

// Copy will create a new Position struct that is a copy of this one
func (p *Position) Copy() *Position {
	return &Position{p.sx, p.sy, p.offsetX, p.offsetY, p.scrollSpeed, p.savedX}
}

// ScreenX returns the screen X position in the current view
func (p *Position) ScreenX() int {
	return p.sx
}

// ScreenY returns the screen Y position in the current view
func (p *Position) ScreenY() int {
	return p.sy
}

// OffsetX returns the X scroll offset for the current view
func (p *Position) OffsetX() int {
	return p.offsetX
}

// OffsetY returns the Y scroll offset for the current view
func (p *Position) OffsetY() int {
	return p.offsetY
}

// SetX will set the screen X position
func (p *Position) SetX(c *vt100.Canvas, x int) {
	p.sx = x
	w := 80 // default width
	if c != nil {
		w = int(c.W())
	}
	if x < w {
		p.offsetX = 0
	} else {
		p.offsetX = (x - w) + 1
		p.sx -= p.offsetX
	}
}

// SetY will set the screen Y position
func (p *Position) SetY(y int) {
	p.sy = y
}

// DecY will decrease Y by 1
func (p *Position) DecY() {
	p.sy--
	if p.sy < 0 {
		p.sy = 0
	}
}

// IncY will increase Y by 1
func (p *Position) IncY(c *vt100.Canvas) {
	h := 25 // default height
	if c != nil {
		h = int(c.H())
	}

	p.sy++
	if p.sy > (h - 1) {
		p.sy = (h - 1)
	}
}

// SetOffsetX will set the screen X scrolling offset
func (p *Position) SetOffsetX(offsetX int) {
	p.offsetX = offsetX
}

// SetOffsetY will set the screen Y scrolling offset
func (p *Position) SetOffsetY(offsetY int) {
	p.offsetY = offsetY
}

// Up will move the cursor up
func (p *Position) Up() error {
	if p.sy <= 0 {
		return errors.New("already at the top of the canvas")
	}
	p.sy--
	return nil
}

// Down will move the cursor down
func (p *Position) Down(c *vt100.Canvas) error {
	h := 25 // default height
	if c != nil {
		h = int(c.H())
	}
	if p.sy >= h-1 {
		return errors.New("already at the bottom of the canvas")
	}
	p.sy++
	return nil
}

// AtStartOfScreenLine returns true if the position is at the very start of the line, regardless of whitespace and scrolling
func (p *Position) AtStartOfScreenLine() bool {
	return p.sx == 0
}

// AtStartOfTheLine returns true if the position is at the very start of the line, and the line is not scrolled
func (p *Position) AtStartOfTheLine() bool {
	return p.sx == 0 && p.offsetX == 0
}

// LineIndex returns the current line index this position is at
func (p *Position) LineIndex() LineIndex {
	return LineIndex(p.offsetY + p.sy)
}

// LineNumber returns the current line number this Position is at
func (p *Position) LineNumber() LineNumber {
	return LineIndex(p.offsetY + p.sy).LineNumber()
}

// ColNumber returns the current column number this Position is at
func (p *Position) ColNumber() ColNumber {
	return ColIndex(p.offsetX + p.sx).ColNumber()
}
