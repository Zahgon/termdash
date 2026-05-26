// Copyright 2018 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package draw

// text.go contains code that prints UTF-8 encoded strings on the canvas.

import (
	"image"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/private/canvas"
)

// OverrunMode represents
type OverrunMode int

// String implements fmt.Stringer()
func (om OverrunMode) String() string { _ = "STUB: not implemented"; return "" }

// overrunModeNames maps OverrunMode values to human readable names.
var overrunModeNames = map[OverrunMode]string{
	OverrunModeStrict:   "OverrunModeStrict",
	OverrunModeTrim:     "OverrunModeTrim",
	OverrunModeThreeDot: "OverrunModeThreeDot",
}

const (
	// OverrunModeStrict verifies that the drawn value fits the canvas and
	// returns an error if it doesn't.
	OverrunModeStrict OverrunMode = iota

	// OverrunModeTrim trims the part of the text that doesn't fit.
	OverrunModeTrim

	// OverrunModeThreeDot trims the text and places the horizontal ellipsis
	// '…' character at the end.
	OverrunModeThreeDot
)

// TextOption is used to provide options to Text().
type TextOption interface {
	// set sets the provided option.
	set(*textOptions)
}

// textOptions stores the provided options.
type textOptions struct {
	cellOpts    []cell.Option
	maxX        int
	overrunMode OverrunMode
}

// textOption implements TextOption.
type textOption func(*textOptions)

// set implements TextOption.set.
func (to textOption) set(tOpts *textOptions) {
	_ = "STUB: not implemented"

	// TextCellOpts sets options on the cells that contain the text.
	return
}

func TextCellOpts(opts ...cell.Option) TextOption {
	_ = "STUB: not implemented"
	return *new(TextOption)
}

// TextMaxX sets a limit on the X coordinate (column) of the drawn text.
// The X coordinate of all cells used by the text must be within
// start.X <= X < TextMaxX.
// If not provided, the width of the canvas is used as TextMaxX.
func TextMaxX(x int) TextOption { _ = "STUB: not implemented"; return *new(TextOption) }

// TextOverrunMode indicates what to do with text that overruns the TextMaxX()
// or the width of the canvas if TextMaxX() isn't specified.
// Defaults to OverrunModeStrict.
func TextOverrunMode(om OverrunMode) TextOption { _ = "STUB: not implemented"; return *new(TextOption) }

// TrimText trims the provided text so that it fits the specified amount of cells.
func TrimText(text string, maxCells int, om OverrunMode) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Nothing to do if the text fits.

// Only write the rune if it still fits, i.e. don't cut
// full-width runes in half.

// Text prints the provided text on the canvas starting at the provided point.
func Text(c *canvas.Canvas, text string, start image.Point, opts ...TextOption) error {
	_ = "STUB: not implemented"
	return nil
}

// ResizeNeeded draws an unicode character indicating that the canvas size is
// too small to draw meaningful content.
func ResizeNeeded(cvs *canvas.Canvas) error { _ = "STUB: not implemented"; return nil }
