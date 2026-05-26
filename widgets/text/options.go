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
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/mouse"
	"github.com/mum4k/termdash/private/wrap"
)

// options.go contains configurable options for Text.

// Option is used to provide options to New().
type Option interface {
	// set sets the provided option.
	set(*options)
}

// options stores the provided options.
type options struct {
	scrollUp         rune
	scrollDown       rune
	wrapMode         wrap.Mode
	rollContent      bool
	maxTextCells     int
	disableScrolling bool
	mouseUpButton    mouse.Button
	mouseDownButton  mouse.Button
	keyUp            keyboard.Key
	keyDown          keyboard.Key
	keyPgUp          keyboard.Key
	keyPgDown        keyboard.Key
}

// newOptions returns a new options instance.
func newOptions(opts ...Option) *options { _ = "STUB: not implemented"; return nil }

// validate validates the provided options.
func (o *options) validate() error { _ = "STUB: not implemented"; return nil }

// option implements Option.
type option func(*options)

// set implements Option.set.
func (o option) set(opts *options) {
	_ = "STUB: not implemented"

	// ScrollRunes configures the text widgets scroll runes, shown at the top and
	// bottom of a scrollable text widget. If not provided, the default scroll
	// runes will be used.
	return
}

func ScrollRunes(up, down rune) Option { _ = "STUB: not implemented"; return *new(Option) }

// The default scroll runes for content scrolling
const (
	DefaultScrollUpRune   = '⇧'
	DefaultScrollDownRune = '⇩'
)

// WrapAtWords configures the text widget so that it automatically wraps lines
// that are longer than the width of the widget at word boundaries. If not
// provided, long lines are trimmed instead.
func WrapAtWords() Option { _ = "STUB: not implemented"; return *new(Option) }

// WrapAtRunes configures the text widget so that it automatically wraps lines
// that are longer than the width of the widget at rune boundaries. If not
// provided, long lines are trimmed instead.
func WrapAtRunes() Option { _ = "STUB: not implemented"; return *new(Option) }

// RollContent configures the text widget so that it rolls the text content up
// if more text than the size of the container is added. If not provided, the
// content is trimmed instead.
func RollContent() Option { _ = "STUB: not implemented"; return *new(Option) }

// DisableScrolling disables the scrolling of the content using keyboard and
// mouse.
func DisableScrolling() Option { _ = "STUB: not implemented"; return *new(Option) }

// The default mouse buttons for content scrolling.
const (
	DefaultScrollMouseButtonUp   = mouse.ButtonWheelUp
	DefaultScrollMouseButtonDown = mouse.ButtonWheelDown
)

// ScrollMouseButtons configures the mouse buttons that scroll the content.
// The provided buttons must be unique, e.g. the same button cannot be both up
// and down.
func ScrollMouseButtons(up, down mouse.Button) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// The default keys for content scrolling.
const (
	DefaultScrollKeyUp       = keyboard.KeyArrowUp
	DefaultScrollKeyDown     = keyboard.KeyArrowDown
	DefaultScrollKeyPageUp   = keyboard.KeyPgUp
	DefaultScrollKeyPageDown = keyboard.KeyPgDn
)

// ScrollKeys configures the keyboard keys that scroll the content.
// The provided keys must be unique, e.g. the same key cannot be both up and
// down.
func ScrollKeys(up, down, pageUp, pageDown keyboard.Key) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// The default value for the MaxTextCells option.
// Use zero as no limit, for logs you may wish to try 10,000 or higher.
const (
	DefaultMaxTextCells = 0
)

// MaxTextCells limits the text content to this number of terminal cells.
// This is useful when sending large amounts of text to the Text widget, e.g.
// when tailing logs as it will limit the memory usage.
// When the newly added content goes over this number of cells, the Text widget
// behaves as a circular buffer and drops earlier content to accommodate the
// new one.
// Note the count is in cells, not runes, some wide runes can take multiple
// terminal cells.
func MaxTextCells(max int) Option { _ = "STUB: not implemented"; return *new(Option) }
