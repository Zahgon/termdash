// Copyright 2019 Google Inc.
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

// braille_circle.go contains code that draws circles on a braille canvas.

import (
	"image"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/private/canvas/braille"
)

// BrailleCircleOption is used to provide options to BrailleCircle.
type BrailleCircleOption interface {
	// set sets the provided option.
	set(*brailleCircleOptions)
}

// brailleCircleOptions stores the provided options.
type brailleCircleOptions struct {
	cellOpts    []cell.Option
	filled      bool
	pixelChange braillePixelChange

	arcOnly     bool
	startDegree int
	endDegree   int
}

// newBrailleCircleOptions returns a new brailleCircleOptions instance.
func newBrailleCircleOptions() *brailleCircleOptions { _ = "STUB: not implemented"; return nil }

// validate validates the provided options.
func (opts *brailleCircleOptions) validate() error { _ = "STUB: not implemented"; return nil }

// brailleCircleOption implements BrailleCircleOption.
type brailleCircleOption func(*brailleCircleOptions)

// set implements BrailleCircleOption.set.
func (o brailleCircleOption) set(opts *brailleCircleOptions) {
	_ = "STUB: not implemented"

	// BrailleCircleCellOpts sets options on the cells that contain the circle.
	// Cell options on a braille canvas can only be set on the entire cell, not per
	// pixel.
	return
}

func BrailleCircleCellOpts(cOpts ...cell.Option) BrailleCircleOption {
	_ = "STUB: not implemented"
	return *new(BrailleCircleOption)
}

// BrailleCircleFilled indicates that the drawn circle should be filled.
func BrailleCircleFilled() BrailleCircleOption {
	_ = "STUB: not implemented"
	return *new(BrailleCircleOption)
}

// BrailleCircleArcOnly indicates that only a portion of the circle should be drawn.
// The arc will be between the two provided angles in degrees.
// Each angle must be in range 0 <= angle <= 360. Start and end must not be equal.
// The zero angle is on the X axis, angles grow counter-clockwise.
func BrailleCircleArcOnly(startDegree, endDegree int) BrailleCircleOption {
	_ = "STUB: not implemented"
	return *new(BrailleCircleOption)
}

// BrailleCircleClearPixels changes the behavior of BrailleCircle, so that it
// clears the pixels belonging to the circle instead of setting them.
// Useful in order to "erase" a circle from the canvas as opposed to drawing one.
func BrailleCircleClearPixels() BrailleCircleOption {
	_ = "STUB: not implemented"
	return *new(BrailleCircleOption)
}

// BrailleCircle draws an approximated circle with the specified mid point and radius.
// The mid point must be a valid pixel within the canvas.
// All the points that form the circle must fit into the canvas.
// The smallest valid radius is two.
func BrailleCircle(bc *braille.Canvas, mid image.Point, radius int, opts ...BrailleCircleOption) error {
	_ = "STUB: not implemented"
	return nil
}

// drawPoints draws the points onto the canvas.
func drawPoints(bc *braille.Canvas, points []image.Point, opt *brailleCircleOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// fillCircle fills a circle that consists of the provided point and has the
// mid point and radius.
func fillCircle(bc *braille.Canvas, points []image.Point, mid image.Point, radius int, opt *brailleCircleOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// Determine a fill point that should be inside of the circle sector.

// Ensure the fill point falls inside the circle.
// If drawing a partial circle, it must also fall within points belonging
// to the opening.
// This might not be true if drawing a partial circle and the arc is very
// small.

// openingPoints returns points on the lines from the mid point to the circle
// opening when drawing an incomplete circle.
func openingPoints(mid image.Point, radius int, opt *brailleCircleOptions) []image.Point {
	_ = "STUB: not implemented"
	return nil
}

// circlePoints returns a list of points that represent a circle with
// the specified mid point and radius.
func circlePoints(mid image.Point, radius int) []image.Point { _ = "STUB: not implemented"; return nil }

// Bresenham algorithm.
// https://en.wikipedia.org/wiki/Midpoint_circle_algorithm

// Cheap multiplication by two.
