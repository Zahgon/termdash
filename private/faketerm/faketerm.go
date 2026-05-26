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

// Package faketerm is a fake implementation of the terminal for the use in tests.
package faketerm

import (
	"context"
	"image"
	"sync"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/private/canvas/buffer"
	"github.com/mum4k/termdash/private/event/eventqueue"
	"github.com/mum4k/termdash/terminal/terminalapi"
)

// Option is used to provide options.
type Option interface {
	// set sets the provided option.
	set(*Terminal)
}

// option implements Option.
type option func(*Terminal)

// set implements Option.set.
func (o option) set(t *Terminal) {
	_ = "STUB: not implemented"

	// WithEventQueue provides a queue of events.
	// One event will be consumed from the queue each time Event() is called. If
	// not provided, Event() returns an error on each call.
	return
}

func WithEventQueue(eq *eventqueue.Unbound) Option { _ = "STUB: not implemented"; return *new(Option) }

// Terminal is a fake terminal.
// This implementation is thread-safe.
type Terminal struct {
	// buffer holds the terminal cells.
	buffer buffer.Buffer

	// events is a queue of input events.
	events *eventqueue.Unbound

	// mu protects the buffer.
	mu sync.Mutex
}

// New returns a new fake Terminal.
func New(size image.Point, opts ...Option) (*Terminal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustNew is like New, but panics on all errors.
func MustNew(size image.Point, opts ...Option) *Terminal { _ = "STUB: not implemented"; return nil }

// Resize resizes the terminal to the provided size.
// This also clears the internal buffer.
func (t *Terminal) Resize(size image.Point) error { _ = "STUB: not implemented"; return nil }

// BackBuffer returns the back buffer of the fake terminal.
func (t *Terminal) BackBuffer() buffer.Buffer {
	_ = "STUB: not implemented"
	return *new(buffer.Buffer)
}

// String prints out the buffer into a string.
// This includes the cell runes only, cell options are ignored.
// Implements fmt.Stringer.
func (t *Terminal) String() string { _ = "STUB: not implemented"; return "" }

// Size implements terminalapi.Terminal.Size.
func (t *Terminal) Size() image.Point { _ = "STUB: not implemented"; return *new(image.Point) }

// Area returns the area of the fake terminal.
func (t *Terminal) Area() image.Rectangle { _ = "STUB: not implemented"; return *new(image.Rectangle) }

// Clear implements terminalapi.Terminal.Clear.
func (t *Terminal) Clear(opts ...cell.Option) error { _ = "STUB: not implemented"; return nil }

// Flush implements terminalapi.Terminal.Flush.
func (t *Terminal) Flush() error {
	_ = "STUB: not implemented"
	// nowhere to flush to.
	return nil
}

// SetCursor implements terminalapi.Terminal.SetCursor.
func (t *Terminal) SetCursor(p image.Point) { _ = "STUB: not implemented"; return }

// HideCursor implements terminalapi.Terminal.HideCursor.
func (t *Terminal) HideCursor() { _ = "STUB: not implemented"; return }

// SetCell implements terminalapi.Terminal.SetCell.
func (t *Terminal) SetCell(p image.Point, r rune, opts ...cell.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// Event implements terminalapi.Terminal.Event.
func (t *Terminal) Event(ctx context.Context) terminalapi.Event {
	_ = "STUB: not implemented"
	return *new(terminalapi.Event)
}

// Close closes the terminal. This is a no-op on the fake terminal.
func (t *Terminal) Close() { _ = "STUB: not implemented"; return }
