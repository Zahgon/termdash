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

package text

import (
	"image"

	"github.com/mum4k/termdash/private/canvas"
)

// line_trim.go contains code that trims lines that are too long.

type trimResult struct {
	// trimmed is set to true if the current and the following runes on this
	// line are trimmed.
	trimmed bool

	// curPoint is the updated current point the drawing should continue on.
	curPoint image.Point
}

// drawTrimChar draws the horizontal ellipsis '…' character as the last
// character in the canvas on the specified line.
func drawTrimChar(cvs *canvas.Canvas, line int) error { _ = "STUB: not implemented"; return nil }

// If the penultimate cell contains a full-width rune, we need to clear it
// first. Otherwise the trim char would cover just half of it.

// lineTrim determines if the current line needs to be trimmed. The cvs is the
// canvas assigned to the widget, the curPoint is the current point the widget
// is going to place the curRune at. If line trimming is needed, this function
// replaces the last character with the horizontal ellipsis '…' character.
func lineTrim(cvs *canvas.Canvas, curPoint image.Point, curRune rune, opts *options) (*trimResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Don't trim if the widget is configured to wrap lines.

// Newline characters are never trimmed, they start the next line.
