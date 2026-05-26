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

// Package termbox implements terminal using the nsf/termbox-go library.
// Prefer to use tcell instead, nsf/termbox-go is no longer maintained.
package termbox

import (
	"context"
	"image"

	"github.com/mum4k/termdash/cell"
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

	// DefaultColorMode is the default value for the ColorMode option.
	return
}

const DefaultColorMode = terminalapi.ColorMode256

// ColorMode sets the terminal color mode.
// Defaults to DefaultColorMode.
func ColorMode(cm terminalapi.ColorMode) Option { _ = "STUB: not implemented"; return *new(Option) }

// Terminal provides input and output to a real terminal. Wraps the
// nsf/termbox-go terminal implementation. This object is not thread-safe.
//
// Prefer to use tcell instead, nsf/termbox-go is no longer maintained.
//
// Implements terminalapi.Terminal.
type Terminal struct {
	// events is a queue of input events.
	events *eventqueue.Unbound

	// done gets closed when Close() is called.
	done chan struct{}

	// Options.
	colorMode terminalapi.ColorMode
}

// newTerminal creates the terminal and applies the options.
func newTerminal(opts ...Option) *Terminal { _ = "STUB: not implemented"; return nil }

// New returns a new termbox based Terminal.
// Call Close() when the terminal isn't required anymore.
func New(opts ...Option) (*Terminal, error) { _ = "STUB: not implemented"; return nil, nil }

// Stops when Close() is called.

// Size implements terminalapi.Terminal.Size.
func (t *Terminal) Size() image.Point { _ = "STUB: not implemented"; return *new(image.Point) }

// Clear implements terminalapi.Terminal.Clear.
func (t *Terminal) Clear(opts ...cell.Option) error { _ = "STUB: not implemented"; return nil }

// Flush implements terminalapi.Terminal.Flush.
func (t *Terminal) Flush() error {
	_ = "STUB: not implemented"

	// SetCursor implements terminalapi.Terminal.SetCursor.
	return nil
}

func (t *Terminal) SetCursor(p image.Point) { _ = "STUB: not implemented"; return }

// HideCursor implements terminalapi.Terminal.HideCursor.
func (t *Terminal) HideCursor() {
	_ = "STUB: not implemented"

	// SetCell implements terminalapi.Terminal.SetCell.
	return
}

func (t *Terminal) SetCell(p image.Point, r rune, opts ...cell.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// pollEvents polls and enqueues the input events.
func (t *Terminal) pollEvents() { _ = "STUB: not implemented"; return }

// Event implements terminalapi.Terminal.Event.
func (t *Terminal) Event(ctx context.Context) terminalapi.Event {
	_ = "STUB: not implemented"
	return *new(terminalapi.Event)
}

// Close closes the terminal, should be called when the terminal isn't required
// anymore to return the screen to a sane state.
// Implements terminalapi.Terminal.Close.
func (t *Terminal) Close() { _ = "STUB: not implemented"; return }
