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

// Package testbraille provides helpers for tests that use the braille package.
package testbraille

import (
	"image"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/private/canvas"
	"github.com/mum4k/termdash/private/canvas/braille"
	"github.com/mum4k/termdash/private/faketerm"
)

// MustNew returns a new canvas or panics.
func MustNew(area image.Rectangle) *braille.Canvas { _ = "STUB: not implemented"; return nil }

// MustApply applies the canvas on the terminal or panics.
func MustApply(bc *braille.Canvas, t *faketerm.Terminal) { _ = "STUB: not implemented"; return }

// MustSetPixel sets the specified pixel or panics.
func MustSetPixel(bc *braille.Canvas, p image.Point, opts ...cell.Option) {
	_ = "STUB: not implemented"
	return
}

// MustClearPixel clears the specified pixel or panics.
func MustClearPixel(bc *braille.Canvas, p image.Point, opts ...cell.Option) {
	_ = "STUB: not implemented"
	return
}

// MustCopyTo copies the braille canvas onto the provided canvas or panics.
func MustCopyTo(bc *braille.Canvas, dst *canvas.Canvas) { _ = "STUB: not implemented"; return }

// MustSetCellOpts sets the cell options or panics.
func MustSetCellOpts(bc *braille.Canvas, cellPoint image.Point, opts ...cell.Option) {
	_ = "STUB: not implemented"
	return
}

// MustSetAreaCellOpts sets the cell options in the area or panics.
func MustSetAreaCellOpts(bc *braille.Canvas, cellArea image.Rectangle, opts ...cell.Option) {
	_ = "STUB: not implemented"
	return
}
