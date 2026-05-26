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

// Package textinput implements a widget that accepts text input.
package textinput

import (
	"image"
	"sync"

	"github.com/mum4k/termdash/private/canvas"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgetapi"
)

// TextInput accepts text input from the user.
//
// Displays an input field and an optional text label. The input field allows
// the user to edit and submit text.
//
// The text can be submitted by pressing enter or read at any time by calling
// Read. The text input field can be navigated using arrows, the Home and End
// button and using mouse.
//
// Implements widgetapi.Widget. This object is thread-safe.
type TextInput struct {
	// mu protects the widget.
	mu sync.Mutex

	// editor tracks the edits and the state of the text input field.
	editor *fieldEditor

	// forField is the area that was occupied by the text input field last
	// time Draw() was called.
	forField image.Rectangle

	// opts are the provided options.
	opts *options
}

// New returns a new TextInput.
func New(opts ...Option) (*TextInput, error) { _ = "STUB: not implemented"; return nil, nil }

// Vars to be replaced from tests.
var (
	// textFieldRune is the rune used in cells reserved for the text input
	// field if no text is present.
	// Changed from tests to provide readable test failures.
	textFieldRune rune

	// cursorRune is rune that represents the cursor position.
	cursorRune rune
)

// Read reads the content of the text input field.
func (ti *TextInput) Read() string { _ = "STUB: not implemented"; return "" }

// ReadAndClear reads the content of the text input field and clears it.
func (ti *TextInput) ReadAndClear() string { _ = "STUB: not implemented"; return "" }

// drawLabel draws the text label in the area.
func (ti *TextInput) drawLabel(cvs *canvas.Canvas, labelAr image.Rectangle) error {
	_ = "STUB: not implemented"
	return nil
}

// drawField draws the text input field.
func (ti *TextInput) drawField(cvs *canvas.Canvas, text string) error {
	_ = "STUB: not implemented"
	return nil
}

// drawCursor draws the cursor within the text input field.
func (ti *TextInput) drawCursor(cvs *canvas.Canvas, curPos int) error {
	_ = "STUB: not implemented"
	return nil
}

// Draw draws the TextInput widget onto the canvas.
// Implements widgetapi.Widget.Draw.
func (ti *TextInput) Draw(cvs *canvas.Canvas, meta *widgetapi.Meta) error {
	_ = "STUB: not implemented"
	return nil
}

// keyboard processes keyboard events.
// Returns a bool indicating if the content was submitted and the text in the
// field at submission time.
// Implements widgetapi.Widget.Keyboard.
func (ti *TextInput) keyboard(k *terminalapi.Keyboard) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

// Ignore unsupported runes.

// Ignore filtered runes.

// Keyboard processes keyboard events.
// Implements widgetapi.Widget.Keyboard.
func (ti *TextInput) Keyboard(k *terminalapi.Keyboard, meta *widgetapi.EventMeta) error {
	_ = "STUB: not implemented"
	return nil
}

// Mutex must be released when calling the callback.
// Users might call container methods from the callback like the
// Container.Update, see #205.

// Mouse processes mouse events.
// Implements widgetapi.Widget.Mouse.
func (ti *TextInput) Mouse(m *terminalapi.Mouse, meta *widgetapi.EventMeta) error {
	_ = "STUB: not implemented"
	return nil
}

// minFieldHeight is the minimum height in cells needed for the text input field.
const minFieldHeight = 1

// Options implements widgetapi.Widget.Options.
func (ti *TextInput) Options() widgetapi.Options {
	_ = "STUB: not implemented"
	return *new(widgetapi.Options)
}

// split splits the available area into label and text input areas according to
// configuration. The returned labelAr might be image.ZR if no label was
// configured.
func split(cvsAr image.Rectangle, label string, widthPerc *int) (labelAr, textAr image.Rectangle, err error) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle), *new(image.Rectangle), nil
}

// Neither a label nor width percentage specified.

// hideText returns the text with all runes replaced with hr.
func hideText(text string, hr rune) string { _ = "STUB: not implemented"; return "" }
