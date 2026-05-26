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

// Package alignfor provides functions that align elements.
package alignfor

import (
	"image"

	"github.com/mum4k/termdash/align"
)

// hAlign aligns the given area in the rectangle horizontally.
func hAlign(rect image.Rectangle, ar image.Rectangle, h align.Horizontal) (image.Rectangle, error) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle), nil
}

// Use gap from above.

// vAlign aligns the given area in the rectangle vertically.
func vAlign(rect image.Rectangle, ar image.Rectangle, v align.Vertical) (image.Rectangle, error) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle), nil
}

// Use gap from above.

// Rectangle aligns the area within the rectangle returning the
// aligned area. The area must fall within the rectangle.
func Rectangle(rect image.Rectangle, ar image.Rectangle, h align.Horizontal, v align.Vertical) (image.Rectangle, error) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle), nil
}

// Text aligns the text within the given rectangle, returns the start point for the text.
// For the purposes of the alignment this assumes that text will be trimmed if
// it overruns the rectangle.
// This only supports a single line of text, the text must not contain non-printable characters,
// allows empty text.
func Text(rect image.Rectangle, text string, h align.Horizontal, v align.Vertical) (image.Point, error) {
	_ = "STUB: not implemented"
	return *new(image.Point), nil
}

// For the purposes of aligning the text, assume that it will be
// trimmed to the available space.
