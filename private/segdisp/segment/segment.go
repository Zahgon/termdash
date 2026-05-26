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

// Package segment provides functions that draw a single segment.
package segment

import (
	"image"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/private/canvas/braille"
)

// Type identifies the type of the segment that is drawn.
type Type int

// String implements fmt.Stringer()
func (st Type) String() string { _ = "STUB: not implemented"; return "" }

// segmentTypeNames maps Type values to human readable names.
var segmentTypeNames = map[Type]string{
	Horizontal: "Horizontal",
	Vertical:   "Vertical",
}

const (
	segmentTypeUnknown Type = iota

	// Horizontal is a horizontal segment.
	Horizontal
	// Vertical is a vertical segment.
	Vertical

	segmentTypeMax // Used for validation.
)

// Option is used to provide options.
type Option interface {
	// set sets the provided option.
	set(*options)
}

// options stores the provided options.
type options struct {
	cellOpts      []cell.Option
	skipSlopesLTE int
	reverseSlopes bool
}

// option implements Option.
type option func(*options)

// set implements Option.set.
func (o option) set(opts *options) {
	_ = "STUB: not implemented"

	// CellOpts sets options on the cells that contain the segment.
	// Cell options on a braille canvas can only be set on the entire cell, not per
	// pixel.
	return
}

func CellOpts(cOpts ...cell.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

// SkipSlopesLTE if provided instructs HV to not create slopes at the ends of a
// segment if the height of the horizontal or the width of the vertical segment
// is less or equal to the provided value.
func SkipSlopesLTE(v int) Option { _ = "STUB: not implemented"; return *new(Option) }

// ReverseSlopes if provided reverses the order in which slopes are drawn.
// This only has a visible effect when the horizontal segment has height of two
// or the vertical segment has width of two.
// Without this option segments with height / width of two look like this:
// x  -   |
// x --- ||
// x      |
// x
// x With this option:
// x
// x --- |
// x  -  ||
// x     |
func ReverseSlopes() Option { _ = "STUB: not implemented"; return *new(Option) }

// validArea validates the provided area.
func validArea(ar image.Rectangle) error { _ = "STUB: not implemented"; return nil }

// HV draws a horizontal or a vertical display segment, filling the provided area.
// The segment will have slopes on both of its ends.
func HV(bc *braille.Canvas, ar image.Rectangle, st Type, opts ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

// nextHVLineFn is a function that determines the start and end points of a line
// number num in a horizontal or a vertical segment.
type nextHVLineFn func(num int, ar image.Rectangle, opt *options) (image.Point, image.Point)

// nextHorizLine determines the start and end point of individual lines in a
// horizontal segment.
func nextHorizLine(num int, ar image.Rectangle, opt *options) (image.Point, image.Point) {
	_ = "STUB: not implemented"
	// Start and end points of the full row without adjustments for slopes.
	return *new(image.Point), *new(image.Point)
}

// No slopes under these dimensions as we don't have the resolution.

// Don't adjust rows that fall exactly in the middle of the segment height.
// E.g when height divides oddly, we want the middle row to take the full
// width:
//     --
//    ----
//     --
//
// And when the height divides oddly, we want the two middle rows to take
// the full width:
//     --
//    ----
//    ----
//     --
// We only do this for segments that are at least three rows tall.
// For smaller segments we still want this behavior:
//     --
//    ----

// On evenly divided height, we need one less adjustment on every
// row above the half, since two rows are taking the full width
// as shown above.

// nextVertLine determines the start and end point of individual lines in a
// vertical segment.
func nextVertLine(num int, ar image.Rectangle, opt *options) (image.Point, image.Point) {
	_ = "STUB: not implemented"
	// Start and end points of the full column without adjustments for slopes.
	return *new(image.Point), *new(image.Point)
}

// No slopes under these dimensions as we don't have the resolution.

// Don't adjust lines that fall exactly in the middle of the segment height.
// E.g when width divides oddly, we want the middle line to take the full
// height:
//    |
//   |||
//   |||
//    |
//
// And when the width divides oddly, we want the two middle columns to take
// the full height:
//    ||
//   ||||
//   ||||
//    ||
//
// We only do this for segments that are at least three columns wide.
// For smaller segments we still want this behavior:
//     |
//    ||
//    ||
//     |

// On evenly divided width, we need one less adjustment on every
// column above the half, since two lines are taking the full
// height as shown above.

// adjustHoriz given start and end points that identify a horizontal line,
// returns points that are adjusted towards each other on the line by the
// specified amount.
// I.e. the start is moved to the right and the end is moved to the left.
// The points won't be allowed to cross each other.
// The segWidth is the full width of the segment we are drawing.
func adjustHoriz(start, end image.Point, segWidth int, adjust int) (image.Point, image.Point) {
	_ = "STUB: not implemented"
	return *new(image.Point), *new(image.Point)
}

// The width of the segment divides evenly, place start and end next to each other.
// E.g: 0 1 2  3  4 5
//      - - ns ne - -

// The width of the segment divides oddly, place both start and end on the mid point.
// E.g: 0 1  2   3 4
//      - - nsne - -

// adjustVert given start and end points that identify a vertical line,
// returns points that are adjusted towards each other on the line by the
// specified amount.
// I.e. the start is moved down and the end is moved up.
// The points won't be allowed to cross each other.
// The segHeight is the full height of the segment we are drawing.
func adjustVert(start, end image.Point, segHeight int, adjust int) (image.Point, image.Point) {
	_ = "STUB: not implemented"
	return *new(image.Point), *new(image.Point)
}

// swapCoord returns a point with its X and Y coordinates swapped.
func swapCoord(p image.Point) image.Point { _ = "STUB: not implemented"; return *new(image.Point) }

// DiagonalType determines the type of diagonal segment.
type DiagonalType int

// String implements fmt.Stringer()
func (dt DiagonalType) String() string { _ = "STUB: not implemented"; return "" }

// diagonalTypeNames maps DiagonalType values to human readable names.
var diagonalTypeNames = map[DiagonalType]string{
	LeftToRight: "LeftToRight",
	RightToLeft: "RightToLeft",
}

const (
	diagonalTypeUnknown DiagonalType = iota
	// LeftToRight is a diagonal segment from top left to bottom right.
	LeftToRight
	// RightToLeft is a diagonal segment from top right to bottom left.
	RightToLeft

	diagonalTypeMax // Used for validation.
)

// nextDiagLineFn is a function that determines the start and end points of a line
// number num in a diagonal segment.
// Points start and end define the first diagonal exactly in the middle.
// Points prevStart and prevEnd define line num-1.
type nextDiagLineFn func(num int, start, end, prevStart, prevEnd image.Point) (image.Point, image.Point)

// DiagonalOption is used to provide options.
type DiagonalOption interface {
	// set sets the provided option.
	set(*diagonalOptions)
}

// diagonalOptions stores the provided diagonal options.
type diagonalOptions struct {
	cellOpts []cell.Option
}

// diagonalOption implements DiagonalOption.
type diagonalOption func(*diagonalOptions)

// set implements DiagonalOption.set.
func (o diagonalOption) set(opts *diagonalOptions) {
	_ = "STUB: not implemented"

	// DiagonalCellOpts sets options on the cells that contain the diagonal
	// segment.
	// Cell options on a braille canvas can only be set on the entire cell, not per
	// pixel.
	return
}

func DiagonalCellOpts(cOpts ...cell.Option) DiagonalOption {
	_ = "STUB: not implemented"
	return *new(DiagonalOption)
}

// Diagonal draws a diagonal segment of the specified width filling the area.
func Diagonal(bc *braille.Canvas, ar image.Rectangle, width int, dt DiagonalType, opts ...DiagonalOption) error {
	_ = "STUB: not implemented"
	return nil
}

// nextLRLine is a function that determines the start and end points of the
// next line of a left-to-right diagonal segment.
func nextLRLine(num int, start, end, prevStart, prevEnd image.Point) (image.Point, image.Point) {
	_ = "STUB: not implemented"
	return *new(image.Point), *new(image.Point)
}

// Every odd line is placed above the mid diagonal.

// Every even line is placed under the mid diagonal.

// nextRLLine is a function that determines the start and end points of the
// next line of a right-to-left diagonal segment.
func nextRLLine(num int, start, end, prevStart, prevEnd image.Point) (image.Point, image.Point) {
	_ = "STUB: not implemented"
	return *new(image.Point), *new(image.Point)
}

// Every odd line is placed above the mid diagonal.

// Every even line is placed under the mid diagonal.
