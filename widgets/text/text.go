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

// Package text contains a widget that displays textual data.
package text

import (
	"image"
	"sync"

	"github.com/mum4k/termdash/private/canvas"
	"github.com/mum4k/termdash/private/canvas/buffer"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgetapi"
)

// Text displays a block of text.
//
// Each line of the text is either trimmed or wrapped according to the provided
// options. The entire text content is either trimmed or rolled up through the
// canvas according to the provided options.
//
// By default the widget supports scrolling of content with either the keyboard
// or mouse. See the options for the default keys and mouse buttons.
//
// Implements widgetapi.Widget. This object is thread-safe.
type Text struct {
	// content is the text content that will be displayed in the widget as
	// provided by the caller (i.e. not wrapped or pre-processed).
	content []*buffer.Cell
	// wrapped is the content wrapped to the current width of the canvas.
	wrapped [][]*buffer.Cell

	// scroll tracks scrolling the position.
	scroll *scrollTracker

	// lastWidth stores the width of the last canvas the widget drew on.
	// Used to determine if the previous line wrapping was invalidated.
	lastWidth int
	// contentChanged indicates if the text content of the widget changed since
	// the last drawing. Used to determine if the previous line wrapping was
	// invalidated.
	contentChanged bool

	// mu protects the Text widget.
	mu sync.Mutex

	// opts are the provided options.
	opts *options
}

// New returns a new text widget.
func New(opts ...Option) (*Text, error) { _ = "STUB: not implemented"; return nil, nil }

// Reset resets the widget back to empty content.
func (t *Text) Reset() { _ = "STUB: not implemented"; return }

// reset implements Reset, caller must hold t.mu.
func (t *Text) reset() { _ = "STUB: not implemented"; return }

// contentCells calculates the number of cells the content takes to display on
// terminal.
func (t *Text) contentCells() int { _ = "STUB: not implemented"; return 0 }

// Write writes text for the widget to display. Multiple calls append
// additional text. The text contain cannot control characters
// (unicode.IsControl) or space character (unicode.IsSpace) other than:
//
//	' ', '\n'
//
// Any newline ('\n') characters are interpreted as newlines when displaying
// the text.
func (t *Text) Write(text string, wOpts ...WriteOption) error {
	_ = "STUB: not implemented"
	return nil
}

// If MaxTextCells has been set, limit the content if needed.

// minLinesForMarkers are the minimum amount of lines required on the canvas in
// order to draw the scroll markers ('⇧' and '⇩').
const minLinesForMarkers = 3

// drawScrollUp draws the scroll up marker on the first line if there is more
// text "above" the canvas due to the scrolling position. Returns true if the
// marker was drawn.
func (t *Text) drawScrollUp(cvs *canvas.Canvas, cur image.Point, fromLine int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// drawScrollDown draws the scroll down marker on the last line if there is
// more text "below" the canvas due to the scrolling position. Returns true if
// the marker was drawn.
func (t *Text) drawScrollDown(cvs *canvas.Canvas, cur image.Point, fromLine int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// draw draws the text context on the canvas starting at the specified line.
func (t *Text) draw(cvs *canvas.Canvas) error {
	_ = "STUB: not implemented"
	// Tracks the current drawing position on the canvas.
	return nil
}

// Scroll up marker.

// Move to the next line.
// Skip one line of text, the marker replaced it.

// Scroll down marker.

// Skip all lines falling after (under) the canvas.

// Skip over any characters trimmed on the current line.

// Move within the same line.

// Move to the next line.

// Draw draws the text onto the canvas.
// Implements widgetapi.Widget.Draw.
func (t *Text) Draw(cvs *canvas.Canvas, meta *widgetapi.Meta) error {
	_ = "STUB: not implemented"
	return nil
}

// The previous text preprocessing (line wrapping) is invalidated when
// new text is added or the width of the canvas changed.

// Nothing to draw if there's no text.

// Keyboard implements widgetapi.Widget.Keyboard.
func (t *Text) Keyboard(k *terminalapi.Keyboard, meta *widgetapi.EventMeta) error {
	_ = "STUB: not implemented"
	return nil
}

// Mouse implements widgetapi.Widget.Mouse.
func (t *Text) Mouse(m *terminalapi.Mouse, meta *widgetapi.EventMeta) error {
	_ = "STUB: not implemented"
	return nil
}

// Options of the widget
func (t *Text) Options() widgetapi.Options {
	_ = "STUB: not implemented"
	return *new(widgetapi.Options)
}

// At least one line with at least one full-width rune.

// truncateToCells truncates the beginning of text, so that it can be displayed
// in at most maxCells. Setting maxCells to zero disables truncating.
func truncateToCells(text string, maxCells int) string { _ = "STUB: not implemented"; return "" }
