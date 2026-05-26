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

// scale.go calculates the scale of the Y axis.

// YScaleMode determines whether the Y scale is anchored to the zero value.
type YScaleMode int

// String implements fmt.Stringer()
func (ysm YScaleMode) String() string { _ = "STUB: not implemented"; return "" }

// yScaleModeNames maps YScaleMode values to human readable names.
var yScaleModeNames = map[YScaleMode]string{
	YScaleModeAnchored: "YScaleModeAnchored",
	YScaleModeAdaptive: "YScaleModeAdaptive",
}

const (
	// YScaleModeAnchored is a mode in which the Y scale always starts at value
	// zero regardless of the min and max on the series.
	YScaleModeAnchored YScaleMode = iota

	// YScaleModeAdaptive is a mode where the Y scale adapts its base value
	// according to the min and max on the series.
	// I.e. it starts at min for all-positive series and at max for
	// all-negative series.
	YScaleModeAdaptive
)

// YScale is the scale of the Y axis.
type YScale struct {
	// Min is the minimum value on the axis.
	Min *Value
	// Max is the maximum value on the axis.
	Max *Value
	// Step is the step in the value between pixels.
	Step *Value

	// GraphHeight is the height in cells of the area on the canvas that is
	// dedicated to the graph itself.
	GraphHeight int
	// brailleHeight is the height of the braille canvas based on the GraphHeight.
	brailleHeight int

	// valueFormatter is the value formatter used for the labels
	// represented by the values on the scale.
	valueFormatter func(float64) string
}

// String implements fmt.Stringer.
func (ys *YScale) String() string { _ = "STUB: not implemented"; return "" }

// NewYScale calculates the scale of the Y axis, given the boundary values and
// the height of the graph. The nonZeroDecimals dictates rounding of the
// calculated scale, see NewValue for details.
// Max must be greater or equal to min. The graphHeight must be a positive
// number.
func NewYScale(min, max float64, graphHeight, nonZeroDecimals int, mode YScaleMode, valueFormatter func(float64) string) (*YScale, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// One pixel reserved for value zero.

// Anchor the axis at the zero value.

// Even in this mode, we still anchor the axis at the zero if all the
// data points are equal, so we can still draw something.

// PixelToValue given a Y coordinate of the pixel, returns its value according
// to the scale. The coordinate must be within bounds of the graph height
// provided to NewYScale. Y coordinates grow down.
func (ys *YScale) PixelToValue(y int) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

// ValueToPixel given a value, determines the Y coordinate of the pixel that
// most closely represents the value on the line chart according to the scale.
// The value must be within the bounds provided to NewYScale. Y coordinates
// grow down.
func (ys *YScale) ValueToPixel(v float64) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// CellLabel given a Y coordinate of a cell on the canvas, determines value of
// the label that should be next to it. The Y coordinate must be within the
// graphHeight provided to NewYScale. Y coordinates grow down.
func (ys *YScale) CellLabel(y int) (*Value, error) { _ = "STUB: not implemented"; return nil, nil }

// yScaleNewValue is a helper method to get new values for the y scale.
func yScaleNewValue(value float64, nonZeroDecimals int, valueFormatter func(float64) string) *Value {
	_ = "STUB: not implemented"
	return nil
}

// XScale is the scale of the X axis.
type XScale struct {
	// Min is the minimum value on the axis.
	Min *Value
	// Max is the maximum value on the axis.
	Max *Value
	// Step is the step in the value between pixels.
	Step *Value

	// GraphWidth is the width in cells of the area on the canvas that is
	// dedicated to the graph.
	GraphWidth int
	// brailleWidth is the width of the braille canvas based on the GraphWidth.
	brailleWidth int
}

// NewXScale calculates the scale of the X axis, given the boundary values and
// the width on the canvas that is available to the X axis.
// The nonZeroDecimals dictates rounding of the calculated scale, see
// NewValue for details.
// The boundary values must be positive or zero and must be min <= max.
// The graphWidth must be a positive number.
func NewXScale(min, max int, graphWidth, nonZeroDecimals int) (*XScale, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// One pixel reserved for value zero.

// String implements fmt.Stringer.
func (xs *XScale) String() string { _ = "STUB: not implemented"; return "" }

// PixelToValue given a X coordinate of the pixel, returns its value according
// to the scale. The coordinate must be within bounds of the canvas width
// provided to NewXScale. X coordinates grow right.
func (xs *XScale) PixelToValue(x int) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

// ValueToPixel given a value, determines the X coordinate of the pixel that
// most closely represents the value on the line chart according to the scale.
// The value must be within the bounds provided to NewXScale. X coordinates
// grow right.
func (xs *XScale) ValueToPixel(v int) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// ValueToCell given a value, determines the X coordinate of the cell that
// most closely represents the value on the line chart according to the scale.
// The value must be within the bounds provided to NewXScale. X coordinates
// grow right.
func (xs *XScale) ValueToCell(v int) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// CellLabel given an X coordinate of a cell on the canvas, determines value of the
// label that should be next to it. The X coordinate must be within the
// graphWidth provided to NewXScale. X coordinates grow right.
// The returned value is rounded to the nearest int, rounding half away from zero.
func (xs *XScale) CellLabel(x int) (*Value, error) { _ = "STUB: not implemented"; return nil, nil }

// positionToY, given a position within the height, returns the Y coordinate of
// the position. Positions grow up, coordinates grow down.
//
// Positions     Y Coordinates
//
//	2  |  0
//	1  |  1
//	0  |  2
func positionToY(pos int, height int) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// yToPosition is the reverse of positionToY.
func yToPosition(y int, height int) (int, error) { _ = "STUB: not implemented"; return 0, nil }
