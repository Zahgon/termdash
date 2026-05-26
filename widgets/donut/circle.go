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

package donut

// circle.go assists in calculation of points and angles on a circle.

import (
	"image"
)

// startEndAngles given progress indicators and the desired start angle and
// direction, returns the starting and the ending angle of the partial circle
// that represents this progress.
func startEndAngles(current, total, startAngle, direction int) (start, end int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// midAndRadius given an area of a braille canvas, determines the mid point in
// pixels and radius to draw the largest circle that fits.
// The circle's mid point is always positioned on the {0,1} pixel in the chosen
// cell so that any text inside of it can be visually centered.
func midAndRadius(ar image.Rectangle) (image.Point, int) {
	_ = "STUB: not implemented"
	return *new(image.Point), 0
}

// Calculate radius based on the smaller axis.

// availableCells given a radius returns the number of cells that are available
// within the circle and the coordinates of the first cell.
// These coordinates are for a normal (non-braille) canvas.
// That is the cells that do not contain any of the circle points. This is
// important since normal characters and braille characters cannot share the
// same cell.
func availableCells(mid image.Point, radius int) (int, image.Point) {
	_ = "STUB: not implemented"
	return 0, *new(image.Point)
}

// Pixels available for the text only.
// Subtract one for the circle itself.
