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

package sixteen

// attributes.go calculates attributes needed when determining placement of
// segments.

import (
	"image"

	"github.com/mum4k/termdash/private/segdisp/segment"
)

// hvSegType maps horizontal and vertical segments to their type.
var hvSegType = map[Segment]segment.Type{
	A1: segment.Horizontal,
	A2: segment.Horizontal,
	B:  segment.Vertical,
	C:  segment.Vertical,
	D1: segment.Horizontal,
	D2: segment.Horizontal,
	E:  segment.Vertical,
	F:  segment.Vertical,
	G1: segment.Horizontal,
	G2: segment.Horizontal,
	J:  segment.Vertical,
	M:  segment.Vertical,
}

// diaSegType maps diagonal segments to their type.
var diaSegType = map[Segment]segment.DiagonalType{
	H: segment.LeftToRight,
	K: segment.RightToLeft,
	N: segment.RightToLeft,
	L: segment.LeftToRight,
}

// Attributes contains attributes needed to draw the segment display.
// Refer to doc/segment_placement.svg for a visual aid and explanation of the
// usage of the square roots.
type Attributes struct {
	// segSize is the width of a vertical or height of a horizontal segment.
	segSize int

	// diaGap is the shortest distance between slopes on two neighboring
	// perpendicular segments.
	diaGap float64

	// segPeakDist is the distance between the peak of the slope on a segment
	// and the point where the slope ends.
	segPeakDist float64

	// diaLeg is the leg of a square whose hypotenuse is the diaGap.
	diaLeg float64

	// peakToPeak is a horizontal or vertical distance between peaks of two
	// segments.
	peakToPeak int

	// shortLen is length of the shorter segment, e.g. D1.
	shortLen int

	// longLen is length of the longer segment, e.g. F.
	longLen int

	// horizLeftX is the X coordinate where the area of the segment horizontally
	// on the left starts, i.e. X coordinate of F and E.
	horizLeftX int
	// horizMidX is the X coordinate where the area of the segment horizontally in
	// the middle starts, i.e. X coordinate of J and M.
	horizMidX int
	// horizRightX is the X coordinate where the area of the segment horizontally
	// on the right starts, i.e. X coordinate of B and C.
	horizRightX int

	// vertCenY is the Y coordinate where the area of the segment vertically
	// in the center starts, i.e. Y coordinate of G1 and G2.
	vertCenY int
	// VertBotY is the Y coordinate where the area of the segment vertically
	// at the bottom starts, i.e. Y coordinate of D1 and D2.
	VertBotY int
}

// NewAttributes calculates attributes needed to place the segments for the
// provided pixel area.
func NewAttributes(bcAr image.Rectangle) *Attributes { _ = "STUB: not implemented"; return nil }

// diaPerc is the size of the diaGap in percentage of the segment's size.

// Ensure there is at least one pixel diagonally between segments so they
// don't visually blend.

// Display that has segment size of two looks more balanced with peak
// distance of two.

// Prefer odd distances to create centered look.

// Refer to doc/segment_placement.svg.
// Diagram labeled "A mid point".

// hvSegArea returns the area for the specified horizontal or vertical segment.
func (a *Attributes) hvSegArea(s Segment) image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

// hvArFromStart given start coordinates of a segment, its length and its type,
// determines its area.
func (a *Attributes) hvArFromStart(start image.Point, s Segment, length int) image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

// diaSegArea returns the area for the specified diagonal segment.
func (a *Attributes) diaSegArea(s Segment) image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

// diaBetween given four segments (two horizontal and two vertical) returns the
// area between them for a diagonal segment.
func (a *Attributes) diaBetween(top, left, right, bottom Segment) image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

// hvToDiaGapPerc is the size of gap between horizontal or vertical segment
// and the diagonal segment between them in percentage of the diaGap.
