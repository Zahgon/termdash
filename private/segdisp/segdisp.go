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

// Package segdisp provides utilities used by all segment display types.
package segdisp

import (
	"image"

	"github.com/mum4k/termdash/private/canvas"
	"github.com/mum4k/termdash/private/canvas/braille"
)

// Minimum valid size of a cell canvas in order to draw a segment display.
const (
	// MinCols is the smallest valid amount of columns in a cell area.
	MinCols = 6
	// MinRowPixels is the smallest valid amount of rows in a cell area.
	MinRows = 5
)

// aspectRatio is the desired aspect ratio of a single segment display.
var aspectRatio = image.Point{3, 5}

// Required when given an area of cells, returns either an area of the same
// size or a smaller area that is required to draw one segment display (i.e.
// one character).
// Returns a smaller area when the provided area didn't have the required
// aspect ratio.
// Returns an error if the area is too small to draw a segment display, i.e.
// smaller than MinCols x MinRows.
func Required(cellArea image.Rectangle) (image.Rectangle, error) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle), nil
}

// ToBraille converts the canvas into a braille canvas and returns a pixel area
// with aspect ratio adjusted for the segment display.
func ToBraille(cvs *canvas.Canvas) (*braille.Canvas, image.Rectangle, error) {
	_ = "STUB: not implemented"
	return nil, *new(image.Rectangle), nil
}

// SegmentSize given an area for the display segment determines the size of
// individual segments, i.e. the width of a vertical or the height of a
// horizontal segment.
func SegmentSize(ar image.Rectangle) int {
	_ = "STUB: not implemented"
	// widthPerc is the relative width of a segment to the width of the canvas.
	return 0
}

// Segments with odd number of pixels in their width/height look
// better, since the spike at the top of their slopes has only one
// pixel.
