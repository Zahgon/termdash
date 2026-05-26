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

// hv_line.go contains code that draws horizontal and vertical lines.

import (
	"image"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/private/canvas"
)

// HVLineOption is used to provide options to HVLine().
type HVLineOption interface {
	// set sets the provided option.
	set(*hVLineOptions)
}

// hVLineOptions stores the provided options.
type hVLineOptions struct {
	cellOpts  []cell.Option
	lineStyle linestyle.LineStyle
}

// newHVLineOptions returns a new hVLineOptions instance.
func newHVLineOptions() *hVLineOptions { _ = "STUB: not implemented"; return nil }

// hVLineOption implements HVLineOption.
type hVLineOption func(*hVLineOptions)

// set implements HVLineOption.set.
func (o hVLineOption) set(opts *hVLineOptions) {
	_ = "STUB: not implemented"

	// DefaultLineStyle is the default value for the HVLineStyle option.
	return
}

const DefaultLineStyle = linestyle.Light

// HVLineStyle sets the style of the line.
// Defaults to DefaultLineStyle.
func HVLineStyle(ls linestyle.LineStyle) HVLineOption {
	_ = "STUB: not implemented"
	return *new(HVLineOption)
}

// HVLineCellOpts sets options on the cells that contain the line.
func HVLineCellOpts(cOpts ...cell.Option) HVLineOption {
	_ = "STUB: not implemented"
	return *new(HVLineOption)
}

// HVLine represents one horizontal or vertical line.
type HVLine struct {
	// Start is the cell where the line starts.
	Start image.Point
	// End is the cell where the line ends.
	End image.Point
}

// HVLines draws horizontal or vertical lines. Handles drawing of the correct
// characters for locations where any two lines cross (e.g. a corner, a T shape
// or a cross). Each line must be at least two cells long. Both start and end
// must be on the same horizontal (same X coordinate) or same vertical (same Y
// coordinate) line.
func HVLines(c *canvas.Canvas, lines []HVLine, opts ...HVLineOption) error {
	_ = "STUB: not implemented"
	return nil
}

// hVLine represents a line that will be drawn on the canvas.
type hVLine struct {
	// start is the starting point of the line.
	start image.Point

	// end is the ending point of the line.
	end image.Point

	// mainPart is either parts[vLine] or parts[hLine] depending on whether
	// this is horizontal or vertical line.
	mainPart rune

	// opts are the options provided in a call to HVLine().
	opts *hVLineOptions
}

// newHVLine creates a new hVLine instance.
// Swaps start and end if necessary, so that horizontal drawing is always left
// to right and vertical is always top down.
func newHVLine(c *canvas.Canvas, start, end image.Point, opts *hVLineOptions) (*hVLine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// horizontal determines if this is a horizontal line.
func (hvl *hVLine) horizontal() bool { _ = "STUB: not implemented"; return false }

// vertical determines if this is a vertical line.
func (hvl *hVLine) vertical() bool { _ = "STUB: not implemented"; return false }
