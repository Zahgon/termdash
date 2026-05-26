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

package axes

// label.go contains code that calculates the positions of labels on the axes.

import (
	"image"
)

// LabelOrientation represents the orientation of text labels.
type LabelOrientation int

// String implements fmt.Stringer()
func (lo LabelOrientation) String() string { _ = "STUB: not implemented"; return "" }

// labelOrientationNames maps LabelOrientation values to human readable names.
var labelOrientationNames = map[LabelOrientation]string{
	LabelOrientationHorizontal: "LabelOrientationHorizontal",
	LabelOrientationVertical:   "LabelOrientationVertical",
}

const (
	// LabelOrientationHorizontal is the default label orientation where text
	// flows horizontally.
	LabelOrientationHorizontal LabelOrientation = iota

	// LabelOrientationVertical is an orientation where text flows vertically.
	LabelOrientationVertical
)

// Label is one value label on an axis.
type Label struct {
	// Value if the value to be displayed.
	Value *Value

	// Position of the label within the canvas.
	Pos image.Point
}

// yLabels returns labels that should be placed next to the Y axis.
// The labelWidth is the width of the area from the left-most side of the
// canvas until the Y axis (not including the Y axis). This is the area where
// the labels will be placed and aligned.
// Labels are returned in an increasing value order.
// Label value is not trimmed to the provided labelWidth, the label width is
// only used to align the labels. Alignment is done with the assumption that
// longer labels will be trimmed.
func yLabels(scale *YScale, labelWidth int) ([]*Label, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If we have data, place at least two labels, first and last.

// rowLabelArea determines the area available for labels on the specified row.
// The row is the Y coordinate of the row, Y coordinates grow down.
func rowLabelArea(row int, labelWidth int) image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

// rowLabel returns label for the specified row.
func rowLabel(scale *YScale, y int, labelWidth int) (*Label, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// xSpace represents an available space among the X axis.
type xSpace struct {
	// min is the current relative coordinate.
	// These are zero based, i.e. not adjusted to axisStart.
	cur int
	// max is the maximum relative coordinate.
	// These are zero based, i.e. not adjusted to axisStart.
	// The xSpace instance contains points 0 <= x < max
	max int

	// graphZero is the (0, 0) point on the graph.
	graphZero image.Point
}

// newXSpace returns a new xSpace instance initialized for the provided width.
func newXSpace(graphZero image.Point, graphWidth int) *xSpace {
	_ = "STUB: not implemented"
	return nil
}

// Implements fmt.Stringer.
func (xs *xSpace) String() string { _ = "STUB: not implemented"; return "" }

// Remaining returns the remaining size on the X axis.
func (xs *xSpace) Remaining() int { _ = "STUB: not implemented"; return 0 }

// Relative returns the relative coordinate within the space, these are zero
// based.
func (xs *xSpace) Relative() image.Point { _ = "STUB: not implemented"; return *new(image.Point) }

// LabelPos returns the absolute coordinate on the canvas where a label should
// be placed. The is the coordinate that represents the current relative
// coordinate of the space.
func (xs *xSpace) LabelPos() image.Point { _ = "STUB: not implemented"; return *new(image.Point) }

// First down is the axis, second the label.

// Sub subtracts the specified size from the beginning of the available
// space.
func (xs *xSpace) Sub(size int) error { _ = "STUB: not implemented"; return nil }

// xLabels returns labels that should be placed under the X axis.
// The graphZero is the (0, 0) point of the graph area on the canvas.
// Labels are returned in an increasing value order.
// Returned labels shouldn't be trimmed, their count is adjusted so that they
// fit under the width of the axis.
// The customLabels map value positions in the series to the desired custom
// label. These are preferred if present.
func xLabels(scale *XScale, graphZero image.Point, customLabels map[int]string, lo LabelOrientation) ([]*Label, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// colLabel returns a label placed at the beginning of the space.
// The space is adjusted according to how much space was taken by the label.
// Returns nil, nil if the label doesn't fit in the space.
func colLabel(scale *XScale, space *xSpace, customLabels map[int]string, lo LabelOrientation) (*Label, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
