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

// Package fakewidget implements a fake widget that is useful for testing the
// termdash infrastructure.
package fakewidget

import (
	"image"
	"sync"

	"github.com/mum4k/termdash/private/canvas"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgetapi"
)

// outputLines are the number of lines written by this plugin.
const outputLines = 4

const (
	sizeLine = iota
	keyboardLine
	mouseLine
	focusLine
)

// MinimumSize is the minimum size required to draw this widget.
var MinimumSize = image.Point{24, 5}

// Event is an event that should be delivered to the fake widget.
type Event struct {
	// Ev is the event to deliver.
	Ev terminalapi.Event
	// Meta is metadata about the event.
	Meta *widgetapi.EventMeta
}

// Mirror is a fake widget. The fake widget draws a border around its assigned
// canvas and writes the size of its assigned canvas on the first line of the
// canvas.
//
// It writes the last received keyboard event onto the second line. It
// writes the last received mouse event onto the third line. If the widget was
// focused at the time of the event, the event will be prepended with a "F:".
//
// If a non-empty string is provided via the Text() method, that text will be
// written right after the canvas size on the first line. If the widget's
// container is focused it writes "focus" onto the fourth line.
//
// The widget requests the same options that are provided to the constructor.
// If the options or canvas size don't allow for the lines mentioned above, the
// widget skips the ones it has no space for.
//
// This is thread-safe and must not be copied.
// Implements widgetapi.Widget.
type Mirror struct {
	// lines are the lines that will be drawn on the canvas.
	lines []string

	// text is the text provided by the last call to Text().
	text string

	// mu protects lines.
	mu sync.RWMutex

	// opts options for this widget.
	opts widgetapi.Options
}

// New returns a new fake widget.
// The widget will return the provided options on a call to Options().
func New(opts widgetapi.Options) *Mirror { _ = "STUB: not implemented"; return nil }

// Draw draws up to there lines on the canvas, assuming there is space for
// them. Returns an error if the canvas is so small that it cannot even draw a
// 2x2 border on it, or of any of the text lines end up being longer than the
// width of the canvas.
// Draw implements widgetapi.Widget.Draw.
func (mi *Mirror) Draw(cvs *canvas.Canvas, meta *widgetapi.Meta) error {
	_ = "STUB: not implemented"
	return nil
}

// Text stores a text that should be displayed right after the canvas size on
// the first line of the output.
func (mi *Mirror) Text(txt string) {
	_ = "STUB: not implemented"

	// Keyboard draws the received key on the canvas.
	// Sending the keyboard.KeyEsc causes this widget to forget the last keyboard
	// event and return an error instead.
	// Keyboard implements widgetapi.Widget.Keyboard.
	return
}

func (mi *Mirror) Keyboard(k *terminalapi.Keyboard, meta *widgetapi.EventMeta) error {
	_ = "STUB: not implemented"
	return nil
}

// Mouse draws the canvas coordinates of the mouse event and the name of the
// received mouse button on the canvas.
// Sending the mouse.ButtonRight causes this widget to forget the last mouse
// event and return an error instead.
// Mouse implements widgetapi.Widget.Mouse.
func (mi *Mirror) Mouse(m *terminalapi.Mouse, meta *widgetapi.EventMeta) error {
	_ = "STUB: not implemented"
	return nil
}

// Options implements widgetapi.Widget.Options.
func (mi *Mirror) Options() widgetapi.Options {
	_ = "STUB: not implemented"

	// Draw draws the content that would be expected after placing the Mirror
	// widget onto the provided canvas and forwarding the given events.
	return *new(widgetapi.Options)
}

func Draw(t terminalapi.Terminal, cvs *canvas.Canvas, meta *widgetapi.Meta, opts widgetapi.Options, events ...*Event) error {
	_ = "STUB: not implemented"
	return nil
}

// MustDraw is like Draw, but panics on all errors.
func MustDraw(t terminalapi.Terminal, cvs *canvas.Canvas, meta *widgetapi.Meta, opts widgetapi.Options, events ...*Event) {
	_ = "STUB: not implemented"
	return
}

// DrawWithMirror is like Draw, but uses the provided Mirror instead of creating one.
func DrawWithMirror(mirror *Mirror, t terminalapi.Terminal, cvs *canvas.Canvas, meta *widgetapi.Meta, events ...*Event) error {
	_ = "STUB: not implemented"
	return nil
}

// MustDrawWithMirror is like DrawWithMirror, but panics on all errors.
func MustDrawWithMirror(mirror *Mirror, t terminalapi.Terminal, cvs *canvas.Canvas, meta *widgetapi.Meta, events ...*Event) {
	_ = "STUB: not implemented"
	return
}
