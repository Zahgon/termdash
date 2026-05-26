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

// Package area provides functions working with image areas.
package area

import (
	"image"
)

// Size returns the size of the provided area.
func Size(area image.Rectangle) image.Point { _ = "STUB: not implemented"; return *new(image.Point) }

// FromSize returns the corresponding area for the provided size.
func FromSize(size image.Point) (image.Rectangle, error) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle), nil
}

// hSplit returns two new areas created by splitting the provided area at the
// specified percentage of its height, applying the percentage to the top or
// bottom area, depending on the reversed flag. The percentage must be in the
// range 0 <= heightPerc <= 100.
// Can return zero size areas.
func hSplit(area image.Rectangle, heightPerc int, reversed bool) (top image.Rectangle, bottom image.Rectangle, err error) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle), *new(image.Rectangle), nil
}

// HSplit returns two new areas created by splitting the provided area at the
// specified percentage of its height, applying the percentage to the top area.
// The percentage must be in the range 0 <= heightPerc <= 100.
// Can return zero size areas.
func HSplit(area image.Rectangle, heightPerc int) (top image.Rectangle, bottom image.Rectangle, err error) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle), *new(image.Rectangle), nil
}

// HSplitReversed returns two new areas created by splitting the provided area
// at the specified percentage of its height, applying the percentage to the
// bottom area. The percentage must be in the range 0 <= heightPerc <= 100.
// Can return zero size areas.
func HSplitReversed(area image.Rectangle, heightPerc int) (top image.Rectangle, bottom image.Rectangle, err error) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle), *new(image.Rectangle), nil
}

// vSplit returns two new areas created by splitting the provided area at the
// specified percentage of its width, applying the percentage to the left or
// right area, depending on the reversed flag. The percentage must be in the
// range 0 <= widthPerc <= 100.
// Can return zero size areas.
func vSplit(area image.Rectangle, widthPerc int, reversed bool) (left image.Rectangle, right image.Rectangle, err error) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle), *new(image.Rectangle), nil
}

// VSplit returns two new areas created by splitting the provided area at the
// specified percentage of its width, applying the percentage to the left area.
// The percentage must be in the range 0 <= widthPerc <= 100.
// Can return zero size areas.
func VSplit(area image.Rectangle, widthPerc int) (left image.Rectangle, right image.Rectangle, err error) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle), *new(image.Rectangle), nil
}

// VSplitReversed returns two new areas created by splitting the provided area
// at the specified percentage of its width, applying the percentage to the
// right area. The percentage must be in the range 0 <= widthPerc <= 100.
// Can return zero size areas.
func VSplitReversed(area image.Rectangle, widthPerc int) (left image.Rectangle, right image.Rectangle, err error) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle), *new(image.Rectangle), nil
}

// vSplitCells returns two new areas created by splitting the provided area
// after the specified amount of cells of its width, applied to the left or
// right area, depending on the reversed flag. The number of cells must be a
// zero or a positive integer. Providing a zero returns left=image.ZR,
// right=area. Providing a number equal or larger to area's width returns
// left=area, right=image.ZR.
func vSplitCells(area image.Rectangle, cells int, reversed bool) (left image.Rectangle, right image.Rectangle, err error) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle), *new(image.Rectangle), nil
}

// VSplitCells returns two new areas created by splitting the provided area
// after the specified amount of cells of its width, as applied to the left
// area. The number of cells must be a zero or a positive integer. Providing a
// zero returns left=image.ZR, right=area. Providing a number equal or larger to
// area's width returns left=area, right=image.ZR.
func VSplitCells(area image.Rectangle, cells int) (left image.Rectangle, right image.Rectangle, err error) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle), *new(image.Rectangle), nil
}

// VSplitCellsReversed returns two new areas created by splitting the provided
// area after the specified amount of cells of its width, as applied to the
// right area. The number of cells must be a zero or a positive integer.
// Providing a zero returns left=image.ZR, right=area. Providing a number equal
// or larger to area's width returns left=area, right=image.ZR.
func VSplitCellsReversed(area image.Rectangle, cells int) (left image.Rectangle, right image.Rectangle, err error) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle), *new(image.Rectangle), nil
}

// hSplitCells returns two new areas created by splitting the provided area
// after the specified amount of cells of its height, applied to the top or
// bottom area, depending on the reversed flag. The number of cells must be a
// zero or a positive integer. Providing a zero returns top=image.ZR,
// bottom=area. Providing a number equal or larger to area's height returns
// top=area, bottom=image.ZR.
func hSplitCells(area image.Rectangle, cells int, reversed bool) (top image.Rectangle, bottom image.Rectangle, err error) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle), *new(image.Rectangle), nil
}

// HSplitCells returns two new areas created by splitting the provided area
// after the specified amount of cells of its height, as applied to the top
// area. The number of cells must be a zero or a positive integer. Providing a
// zero returns top=image.ZR, bottom=area. Providing a number equal or larger to
// area's height returns top=area, bottom=image.ZR.
func HSplitCells(area image.Rectangle, cells int) (top image.Rectangle, bottom image.Rectangle, err error) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle), *new(image.Rectangle), nil
}

// HSplitCellsReversed returns two new areas created by splitting the provided
// area after the specified amount of cells of its height, as applied to the
// bottom area. The number of cells must be a zero or a positive integer.
// Providing a zero returns top=area, bottom=image.ZR. Providing a number equal
// or larger to area's height returns top=image.ZR, bottom=area.
func HSplitCellsReversed(area image.Rectangle, cells int) (top image.Rectangle, bottom image.Rectangle, err error) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle), *new(image.Rectangle), nil
}

// ExcludeBorder returns a new area created by subtracting a border around the
// provided area. Return the zero area if there isn't enough space to exclude
// the border.
func ExcludeBorder(area image.Rectangle) image.Rectangle {
	_ = "STUB: not implemented"
	// If the area dimensions are smaller than this, subtracting a point for the
	// border on each of its sides results in a zero area.
	return *new(image.Rectangle)
}

// WithRatio returns the largest area that has the requested ratio but is
// either equal or smaller than the provided area. Returns zero area if the
// area or the ratio are zero, or if there is no such area.
func WithRatio(area image.Rectangle, ratio image.Point) image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

// Shrink returns a new area whose size is reduced by the specified amount of
// cells. Can return a zero area if there is no space left in the area.
// The values must be zero or positive integers.
func Shrink(area image.Rectangle, topCells, rightCells, bottomCells, leftCells int) (image.Rectangle, error) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle), nil
}

// ShrinkPercent returns a new area whose size is reduced by percentage of its
// width or height. Can return a zero area if there is no space left in the area.
// The topPerc and bottomPerc indicate the percentage of area's height.
// The rightPerc and leftPerc indicate the percentage of area's width.
// The percentages must be in range 0 <= v <= 100.
func ShrinkPercent(area image.Rectangle, topPerc, rightPerc, bottomPerc, leftPerc int) (image.Rectangle, error) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle), nil
}

// MoveUp returns a new area that is moved up by the specified amount of cells.
// Returns an error if the move would result in negative Y coordinates.
// The values must be zero or positive integers.
func MoveUp(area image.Rectangle, cells int) (image.Rectangle, error) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle), nil
}

// MoveDown returns a new area that is moved down by the specified amount of
// cells.
// The values must be zero or positive integers.
func MoveDown(area image.Rectangle, cells int) (image.Rectangle, error) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle), nil
}
