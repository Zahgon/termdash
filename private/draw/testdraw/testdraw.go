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

// Package testdraw provides helpers for tests that use the draw package.
package testdraw

import (
	"image"

	"github.com/mum4k/termdash/private/canvas"
	"github.com/mum4k/termdash/private/canvas/braille"
	"github.com/mum4k/termdash/private/draw"
)

// MustBorder draws border on the canvas or panics.
func MustBorder(c *canvas.Canvas, border image.Rectangle, opts ...draw.BorderOption) {
	_ = "STUB: not implemented"
	return
}

// MustText draws the text on the canvas or panics.
func MustText(c *canvas.Canvas, text string, start image.Point, opts ...draw.TextOption) {
	_ = "STUB: not implemented"
	return
}

// MustVerticalText draws the vertical text on the canvas or panics.
func MustVerticalText(c *canvas.Canvas, text string, start image.Point, opts ...draw.VerticalTextOption) {
	_ = "STUB: not implemented"
	return
}

// MustRectangle draws the rectangle on the canvas or panics.
func MustRectangle(c *canvas.Canvas, r image.Rectangle, opts ...draw.RectangleOption) {
	_ = "STUB: not implemented"
	return
}

// MustHVLines draws the vertical / horizontal lines or panics.
func MustHVLines(c *canvas.Canvas, lines []draw.HVLine, opts ...draw.HVLineOption) {
	_ = "STUB: not implemented"
	return
}

// MustBrailleLine draws the braille line or panics.
func MustBrailleLine(bc *braille.Canvas, start, end image.Point, opts ...draw.BrailleLineOption) {
	_ = "STUB: not implemented"
	return
}

// MustBrailleCircle draws the braille circle or panics.
func MustBrailleCircle(bc *braille.Canvas, mid image.Point, radius int, opts ...draw.BrailleCircleOption) {
	_ = "STUB: not implemented"
	return
}

// MustResizeNeeded draws the character or panics.
func MustResizeNeeded(cvs *canvas.Canvas) { _ = "STUB: not implemented"; return }
