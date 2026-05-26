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

// braille_line.go contains code that draws lines on a braille canvas.

import (
	"image"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/private/canvas/braille"
)

// braillePixelChange represents an action on a pixel on the braille canvas.
type braillePixelChange int

// String implements fmt.Stringer()
func (bpc braillePixelChange) String() string { _ = "STUB: not implemented"; return "" }

// braillePixelChangeNames maps braillePixelChange values to human readable names.
var braillePixelChangeNames = map[braillePixelChange]string{
	braillePixelChangeSet:   "braillePixelChangeSet",
	braillePixelChangeClear: "braillePixelChangeClear",
}

const (
	braillePixelChangeUnknown braillePixelChange = iota

	braillePixelChangeSet
	braillePixelChangeClear
)

// BrailleLineOption is used to provide options to BrailleLine().
type BrailleLineOption interface {
	// set sets the provided option.
	set(*brailleLineOptions)
}

// brailleLineOptions stores the provided options.
type brailleLineOptions struct {
	cellOpts    []cell.Option
	pixelChange braillePixelChange
}

// newBrailleLineOptions returns a new brailleLineOptions instance.
func newBrailleLineOptions() *brailleLineOptions { _ = "STUB: not implemented"; return nil }

// brailleLineOption implements BrailleLineOption.
type brailleLineOption func(*brailleLineOptions)

// set implements BrailleLineOption.set.
func (o brailleLineOption) set(opts *brailleLineOptions) {
	_ = "STUB: not implemented"

	// BrailleLineCellOpts sets options on the cells that contain the line.
	// Cell options on a braille canvas can only be set on the entire cell, not per
	// pixel.
	return
}

func BrailleLineCellOpts(cOpts ...cell.Option) BrailleLineOption {
	_ = "STUB: not implemented"
	return *new(BrailleLineOption)
}

// BrailleLineClearPixels changes the behavior of BrailleLine, so that it
// clears the pixels belonging to the line instead of setting them.
// Useful in order to "erase" a line from the canvas as opposed to drawing one.
func BrailleLineClearPixels() BrailleLineOption {
	_ = "STUB: not implemented"
	return *new(BrailleLineOption)
}

// BrailleLine draws an approximated line segment on the braille canvas between
// the two provided points.
// Both start and end must be valid points within the canvas. Start and end can
// be the same point in which case only one pixel will be set on the braille
// canvas.
// The start or end coordinates must not be negative.
func BrailleLine(bc *braille.Canvas, start, end image.Point, opts ...BrailleLineOption) error {
	_ = "STUB: not implemented"
	return nil
}

// brailleLinePoints returns the points to set when drawing the line.
func brailleLinePoints(start, end image.Point) []image.Point {
	_ = "STUB: not implemented"
	// Implements Bresenham's line algorithm.
	// https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
	return nil
}

// lineLow returns points that create a line whose horizontal projection
// (end.X - start.X) is longer than its vertical projection
// (end.Y - start.Y).
func lineLow(x0, y0, x1, y1 int) []image.Point { _ = "STUB: not implemented"; return nil }

// lineHigh returns points that createa line whose vertical projection
// (end.Y - start.Y) is longer than its horizontal projection
// (end.X - start.X).
func lineHigh(x0, y0, x1, y1 int) []image.Point { _ = "STUB: not implemented"; return nil }
