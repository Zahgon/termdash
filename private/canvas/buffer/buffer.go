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

// Package buffer implements a 2-D buffer of cells.
package buffer

import (
	"image"

	"github.com/mum4k/termdash/cell"
)

// NewCells breaks the provided text into cells and applies the options.
func NewCells(text string, opts ...cell.Option) []*Cell { _ = "STUB: not implemented"; return nil }

// Cell represents a single cell on the terminal.
type Cell struct {
	// Rune is the rune stored in the cell.
	Rune rune

	// Opts are the cell options.
	Opts *cell.Options
}

// String implements fmt.Stringer.
func (c *Cell) String() string { _ = "STUB: not implemented"; return "" }

// NewCell returns a new cell.
func NewCell(r rune, opts ...cell.Option) *Cell { _ = "STUB: not implemented"; return nil }

// Copy returns a copy the cell.
func (c *Cell) Copy() *Cell { _ = "STUB: not implemented"; return nil }

// Apply applies the provided options to the cell.
func (c *Cell) Apply(opts ...cell.Option) { _ = "STUB: not implemented"; return }

// Buffer is a 2-D buffer of cells.
// The axes increase right and down.
// Uninitialized buffer is invalid, use New to create an instance.
// Don't set cells directly, use the SetCell method instead which safely
// handles limits and wide unicode characters.
type Buffer [][]*Cell

// New returns a new Buffer of the provided size.
func New(size image.Point) (Buffer, error) { _ = "STUB: not implemented"; return *new(Buffer), nil }

// SetCell sets the rune of the specified cell in the buffer. Returns the
// number of cells the rune occupies, wide runes can occupy multiple cells when
// printed on the terminal. See http://www.unicode.org/reports/tr11/.
// Use the options to specify which attributes to modify, if an attribute
// option isn't specified, the attribute retains its previous value.
func (b Buffer) SetCell(p image.Point, r rune, opts ...cell.Option) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Even if the rune is invisible, like the zero-value rune, it still
// occupies at least the target cell.

// IsPartial returns true if the cell at the specified point holds a part of a
// full width rune from a previous cell. See
// http://www.unicode.org/reports/tr11/.
func (b Buffer) IsPartial(p image.Point) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// RemWidth returns the remaining width (horizontal row of cells) available
// from and inclusive of the specified point.
func (b Buffer) RemWidth(p image.Point) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Size returns the size of the buffer.
func (b Buffer) Size() image.Point { _ = "STUB: not implemented"; return *new(image.Point) }
