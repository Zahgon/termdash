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

// Package canvas defines the canvas that the widgets draw on.
package canvas

import (
	"image"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/private/canvas/buffer"
	"github.com/mum4k/termdash/terminal/terminalapi"
)

// Canvas is where a widget draws its output for display on the terminal.
type Canvas struct {
	// area is the area the buffer was created for.
	// Contains absolute coordinates on the target terminal, while the buffer
	// contains relative zero-based coordinates for this canvas.
	area image.Rectangle

	// buffer is where the drawing happens.
	buffer buffer.Buffer
}

// New returns a new Canvas with a buffer for the provided area.
func New(ar image.Rectangle) (*Canvas, error) { _ = "STUB: not implemented"; return nil, nil }

// Size returns the size of the 2-D canvas.
func (c *Canvas) Size() image.Point {
	_ = "STUB: not implemented"
	return *

	// Area returns the area of the 2-D canvas.
	new(image.Point)
}

func (c *Canvas) Area() image.Rectangle { _ = "STUB: not implemented"; return *new(image.Rectangle) }

// Clear clears all the content on the canvas.
func (c *Canvas) Clear() error { _ = "STUB: not implemented"; return nil }

// SetCell sets the rune of the specified cell on the canvas. Returns the
// number of cells the rune occupies, wide runes can occupy multiple cells when
// printed on the terminal. See http://www.unicode.org/reports/tr11/.
// Use the options to specify which attributes to modify, if an attribute
// option isn't specified, the attribute retains its previous value.
func (c *Canvas) SetCell(p image.Point, r rune, opts ...cell.Option) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Cell returns a copy of the specified cell.
func (c *Canvas) Cell(p image.Point) (*buffer.Cell, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetCellOpts sets options on the specified cell of the canvas without
// modifying the content of the cell.
// Sets the default cell options if no options are provided.
// This method is idempotent.
func (c *Canvas) SetCellOpts(p image.Point, opts ...cell.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// Set the default options.

// SetAreaCells is like SetCell, but sets the specified rune and options on all
// the cells within the provided area.
// This method is idempotent.
func (c *Canvas) SetAreaCells(cellArea image.Rectangle, r rune, opts ...cell.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// SetAreaCellOpts is like SetCellOpts, but sets the specified options on all
// the cells within the provided area.
func (c *Canvas) SetAreaCellOpts(cellArea image.Rectangle, opts ...cell.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// setCellFunc is a function that sets cell content on a terminal or a canvas.
type setCellFunc func(image.Point, rune, ...cell.Option) error

// copyTo is the internal implementation of code that copies the content of a
// canvas. If a non zero offset is provided, all the copied points are offset by
// this amount.
// The dstSetCell function is called for every point in this canvas when
// copying it to the destination.
func (c *Canvas) copyTo(offset image.Point, dstSetCell setCellFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip over partial cells, i.e. cells that follow a cell
// containing a full-width rune. A full-width rune takes only
// one cell in the buffer, but two on the terminal.
// See http://www.unicode.org/reports/tr11/.

// Apply applies the canvas to the corresponding area of the terminal.
func (c *Canvas) Apply(t terminalapi.Terminal) error {
	_ = "STUB: not implemented"
	// Note - the size of the terminal might have changed since we started
	// drawing, since terminal windows are inherently racy (the user can resize
	// them at any time).
	//
	// This is ok, since the underlying terminal layer will just ignore cells
	// that are out of bounds and termdash will redraw again once it receives
	// the resize event. Regression for #281.
	return nil
}

// The image.Point{0, 0} of this canvas isn't always exactly at
// image.Point{0, 0} on the terminal.
// Depends on area assigned by the container.

// CopyTo copies the content of this canvas onto the destination canvas.
// This canvas can have an offset when compared to the destination canvas, i.e.
// the area of this canvas doesn't have to be zero-based.
func (c *Canvas) CopyTo(dst *Canvas) error { _ = "STUB: not implemented"; return nil }

// Neither of the two canvases (source and destination) have to be zero
// based. Canvas is not zero based if it is positioned elsewhere, i.e.
// providing a smaller view of another canvas.
// E.g. a widget can assign a smaller portion of its canvas to a component
// in order to restrict drawing of this component to a smaller area. To do
// this it can create a sub-canvas. This sub-canvas can have a specific
// starting position other than image.Point{0, 0} relative to the parent
// canvas. Copying this sub-canvas back onto the parent accounts for this
// offset.
