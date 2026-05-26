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

// Package button implements an interactive widget that can be pressed to
// activate.
package button

import (
	"image"
	"strings"
	"sync"
	"time"

	"github.com/mum4k/termdash/private/attrrange"
	"github.com/mum4k/termdash/private/button"
	"github.com/mum4k/termdash/private/canvas"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgetapi"
)

// CallbackFn is the function called when the button is pressed.
// The callback function must be light-weight, ideally just storing a value and
// returning, since more button presses might occur.
//
// The callback function must be thread-safe as the mouse or keyboard events
// that press the button are processed in a separate goroutine.
//
// If the function returns an error, the widget will forward it back to the
// termdash infrastructure which causes a panic, unless the user provided a
// termdash.ErrorHandler.
type CallbackFn func() error

// TextChunk is a part of or the full text displayed in the button.
type TextChunk struct {
	text  string
	tOpts *textOptions
}

// NewChunk creates a new text chunk. Each chunk of text can have its own cell options.
func NewChunk(text string, tOpts ...TextOption) *TextChunk { _ = "STUB: not implemented"; return nil }

// Button can be pressed using a mouse click or a configured keyboard key.
//
// Upon each press, the button invokes a callback provided by the user.
//
// Implements widgetapi.Widget. This object is thread-safe.
type Button struct {
	// text in the text label displayed in the button.
	text strings.Builder

	// givenTOpts are text options given for the button's of text.
	givenTOpts []*textOptions
	// tOptsTracker tracks the positions in a text to which the givenTOpts apply.
	tOptsTracker *attrrange.Tracker

	// mouseFSM tracks left mouse clicks.
	mouseFSM *button.FSM
	// state is the current state of the button.
	state button.State

	// keyTriggerTime is the last time the button was pressed using a keyboard
	// key. It is nil if the button was triggered by a mouse event.
	// Used to draw button presses on keyboard events, since termbox doesn't
	// provide us with release events for keys.
	keyTriggerTime *time.Time

	// callback gets called on each button press.
	callback CallbackFn

	// mu protects the widget.
	mu sync.Mutex

	// opts are the provided options.
	opts *options
}

// New returns a new Button that will display the provided text.
// Each press of the button will invoke the callback function.
// The callback function can be nil in which case pressing the button is a
// no-op.
func New(text string, cFn CallbackFn, opts ...Option) (*Button, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewFromChunks is like New, but allows specifying write options for
// individual chunks of text displayed in the button.
func NewFromChunks(chunks []*TextChunk, cFn CallbackFn, opts ...Option) (*Button, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetCallback replaces the callback function of the button with the one provided.
func (b *Button) SetCallback(cFn CallbackFn) { _ = "STUB: not implemented"; return }

// Vars to be replaced from tests.
var (
	// Runes to use in cells that contain the button.
	// Changed from tests to provide readable test failures.
	buttonRune = ' '
	// Runes to use in cells that contain the shadow.
	// Changed from tests to provide readable test failures.
	shadowRune = ' '

	// timeSince is a function that calculates duration since some time.
	timeSince = time.Since
)

// Draw draws the Button widget onto the canvas.
// Implements widgetapi.Widget.Draw.
func (b *Button) Draw(cvs *canvas.Canvas, meta *widgetapi.Meta) error {
	_ = "STUB: not implemented"
	return nil
}

// drawText draws the text inside the button.
func (b *Button) drawText(cvs *canvas.Canvas, meta *widgetapi.Meta, buttonAr image.Rectangle) error {
	_ = "STUB: not implemented"
	return nil
}

// Text options for the current byte.

// Get the next write options.

// activated asserts whether the keyboard event activated the button.
func (b *Button) keyActivated(k *terminalapi.Keyboard, meta *widgetapi.EventMeta) bool {
	_ = "STUB: not implemented"
	return false
}

// Keyboard processes keyboard events, acts as a button press on the configured
// Key.
//
// Implements widgetapi.Widget.Keyboard.
func (b *Button) Keyboard(k *terminalapi.Keyboard, meta *widgetapi.EventMeta) error {
	_ = "STUB: not implemented"
	return nil
}

// Mutex must be released when calling the callback.
// Users might call container methods from the callback like the
// Container.Update, see #205.

// mouseActivated asserts whether the mouse event activated the button.
func (b *Button) mouseActivated(m *terminalapi.Mouse) bool { _ = "STUB: not implemented"; return false }

// Mouse processes mouse events, acts as a button press if both the press and
// the release happen inside the button.
//
// Implements widgetapi.Widget.Mouse.
func (b *Button) Mouse(m *terminalapi.Mouse, meta *widgetapi.EventMeta) error {
	_ = "STUB: not implemented"
	return nil
}

// Mutex must be released when calling the callback.
// Users might call container methods from the callback like the
// Container.Update, see #205.

// shadowWidth returns the width of the shadow under the button or zero if the
// button shouldn't have any shadow.
func (b *Button) shadowWidth() int { _ = "STUB: not implemented"; return 0 }

// Options implements widgetapi.Widget.Options.
func (b *Button) Options() widgetapi.Options {
	_ = "STUB: not implemented"
	// No need to lock, as the height and width get fixed when New is called.
	return *new(widgetapi.Options)
}
