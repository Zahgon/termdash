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

package terminalapi

import (
	"image"

	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/mouse"
)

// event.go defines events that can be received through the terminal API.

// Event represents an input event.
type Event interface {
	isEvent()
}

// Keyboard is the event used when a key is pressed.
// Implements terminalapi.Event.
type Keyboard struct {
	// Key is the pressed key.
	Key keyboard.Key
}

func (*Keyboard) isEvent() {
	_ = "STUB: not implemented"

	// String implements fmt.Stringer.
	return
}

func (k Keyboard) String() string { _ = "STUB: not implemented"; return "" }

// Resize is the event used when the terminal was resized.
// Implements terminalapi.Event.
type Resize struct {
	// Size is the new size of the terminal.
	Size image.Point
}

func (*Resize) isEvent() {
	_ = "STUB: not implemented"

	// String implements fmt.Stringer.
	return
}

func (r Resize) String() string { _ = "STUB: not implemented"; return "" }

// Mouse is the event used when the mouse is moved or a mouse button is
// pressed.
// Implements terminalapi.Event.
type Mouse struct {
	// Position of the mouse on the terminal.
	Position image.Point
	// Button identifies the pressed button if any.
	Button mouse.Button
}

func (*Mouse) isEvent() {
	_ = "STUB: not implemented"

	// String implements fmt.Stringer.
	return
}

func (m Mouse) String() string { _ = "STUB: not implemented"; return "" }

// Error is an event indicating an error while processing input.
type Error string

// NewError returns a new Error event.
func NewError(e string) *Error { _ = "STUB: not implemented"; return nil }

// NewErrorf returns a new Error event, arguments are similar to fmt.Sprintf.
func NewErrorf(format string, args ...interface{}) *Error { _ = "STUB: not implemented"; return nil }

func (*Error) isEvent() {
	_ = "STUB: not implemented"

	// Error returns the error that occurred.
	return
}

func (e *Error) Error() error { _ = "STUB: not implemented"; return nil }

// String implements fmt.Stringer.
func (e Error) String() string { _ = "STUB: not implemented"; return "" }
