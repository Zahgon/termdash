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

package dotseg

// attributes.go calculates attributes needed when determining placement of
// segments.

import (
	"image"

	"github.com/mum4k/termdash/private/segdisp/sixteen"
)

// attributes contains attributes needed to draw the segment display.
// Refer to doc/segment_placement.svg for a visual aid and explanation of the
// usage of the square roots.
type attributes struct {
	// bcAr is the area the attributes were created for.
	bcAr image.Rectangle

	// segSize is the width of a vertical or height of a horizontal segment.
	segSize int

	// sixteen are attributes of a 16-segment display when placed on the same
	// area.
	sixteen *sixteen.Attributes
}

// newAttributes calculates attributes needed to place the segments for the
// provided pixel area.
func newAttributes(bcAr image.Rectangle) *attributes { _ = "STUB: not implemented"; return nil }

// segArea returns the area for the specified segment.
func (a *attributes) segArea(seg Segment) (image.Rectangle, error) {
	_ = "STUB: not implemented"
	// Dots have double width of normal segments to fill more space in the
	// segment display.
	return *new(image.Rectangle), nil
}

// An area representing the dot which gets aligned and moved into position
// below.

// moveBySize is the multiplier of segment size to determine by how many
// pixels to move D1 and D2 up and down from the center.

// Align at the middle of the bottom.

// Shift up to where the sixteen segment actually places its bottom
// segments.

// Shift further up by one segment size, since the dots have double width.
