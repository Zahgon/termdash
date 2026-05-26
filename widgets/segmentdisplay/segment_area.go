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

package segmentdisplay

// segment_area.go contains code that determines how many segments we can fit
// in the canvas.

import (
	"image"
)

// segArea contains information about the area that will contain the segments.
type segArea struct {
	// segment is the area for one segment.
	segment image.Rectangle
	// canFit is the number of segments we can fit on the canvas.
	canFit int
	// gapPixels is the size of gaps between segments in pixels.
	gapPixels int
	// gaps is the number of gaps that will be drawn.
	gaps int
}

// needArea returns the complete area required for all the segments that we can
// fit and any gaps.
func (sa *segArea) needArea() image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

// newSegArea calculates the area for segments given available canvas area,
// length of the text to be displayed and the size of gap between segments
func newSegArea(cvsAr image.Rectangle, textLen, gapPercent int) (*segArea, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Don't insert gaps after the last segment in the text or the last
// segment we can fit.

// Only insert gaps if we can still fit one more segment with the gap.

// Gap is needed but doesn't fit together with the next segment.
// So insert neither.

// maximizeFit finds the largest individual segment size that enables us to fit
// the most characters onto a canvas with the provided area. Returns the area
// required for a single segment and the number of segments we can fit.
func maximizeFit(cvsAr image.Rectangle, textLen, gapPercent int) (*segArea, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
