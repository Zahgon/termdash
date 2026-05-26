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

package button

// options.go contains configurable options for Button.

import (
	"time"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/widgetapi"
)

// Option is used to provide options.
type Option interface {
	// set sets the provided option.
	set(*options)
}

// option implements Option.
type option func(*options)

// set implements Option.set.
func (o option) set(opts *options) {
	_ = "STUB: not implemented"

	// options holds the provided options.
	return
}

type options struct {
	fillColor             cell.Color
	focusedFillColor      *cell.Color
	pressedFillColor      *cell.Color
	textColor             cell.Color
	textHorizontalPadding int
	shadowColor           cell.Color
	disableShadow         bool
	height                int
	width                 int
	focusedKeys           map[keyboard.Key]bool
	globalKeys            map[keyboard.Key]bool
	keyUpDelay            time.Duration
}

// validate validates the provided options.
func (o *options) validate() error { _ = "STUB: not implemented"; return nil }

// keyScope stores a key and its scope.
type keyScope struct {
	key   keyboard.Key
	scope widgetapi.KeyScope
}

// newOptions returns options with the default values set.
func newOptions(text string) *options { _ = "STUB: not implemented"; return nil }

// FillColor sets the fill color of the button.
func FillColor(c cell.Color) Option { _ = "STUB: not implemented"; return *new(Option) }

// FocusedFillColor sets the fill color of the button when the widget's
// container is focused.
// Defaults to FillColor.
func FocusedFillColor(c cell.Color) Option { _ = "STUB: not implemented"; return *new(Option) }

// PressedFillColor sets the fill color of the button when it is pressed.
// Defaults to FillColor.
func PressedFillColor(c cell.Color) Option { _ = "STUB: not implemented"; return *new(Option) }

// TextColor sets the color of the text label in the button.
func TextColor(c cell.Color) Option { _ = "STUB: not implemented"; return *new(Option) }

// ShadowColor sets the color of the shadow under the button.
func ShadowColor(c cell.Color) Option { _ = "STUB: not implemented"; return *new(Option) }

// DefaultHeight is the default for the Height option.
const DefaultHeight = 3

// Height sets the height of the button in cells.
// Must be a positive non-zero integer.
// Defaults to DefaultHeight.
func Height(cells int) Option { _ = "STUB: not implemented"; return *new(Option) }

// Width sets the width of the button in cells.
// Must be a positive non-zero integer.
// Defaults to the auto-width based on the length of the text label.
// Not all the width may be available to the text if TextHorizontalPadding is
// set to a non-zero integer.
func Width(cells int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WidthFor sets the width of the button as if it was displaying the provided text.
// Useful when displaying multiple buttons with the intention to set all of
// their sizes equal to the one with the longest text.
func WidthFor(text string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Key configures the keyboard key that presses the button.
// The widget responds to this key only if its container is focused.
//
// Clears all keys set by Key() or Keys() previously.
func Key(k keyboard.Key) Option { _ = "STUB: not implemented"; return *new(Option) }

// GlobalKey is like Key, but makes the widget respond to the key even if its
// container isn't focused.
//
// Clears all keys set by GlobalKey() or GlobalKeys() previously.
func GlobalKey(k keyboard.Key) Option { _ = "STUB: not implemented"; return *new(Option) }

// Keys is like Key, but allows to configure multiple keys.
//
// Clears all keys set by Key() or Keys() previously.
func Keys(keys ...keyboard.Key) Option { _ = "STUB: not implemented"; return *new(Option) }

// GlobalKeys is like GlobalKey, but allows to configure multiple keys.
//
// Clears all keys set by GlobalKey() or GlobalKeys() previously.
func GlobalKeys(keys ...keyboard.Key) Option { _ = "STUB: not implemented"; return *new(Option) }

// DefaultKeyUpDelay is the default value for the KeyUpDelay option.
const DefaultKeyUpDelay = 250 * time.Millisecond

// KeyUpDelay is the amount of time the button will remain "pressed down" after
// triggered by the configured key. Termbox doesn't emit events for key
// releases so the button simulates it by timing it.
// This only works if the manual termdash redraw or the periodic redraw
// interval are reasonably close to this delay.
// The duration cannot be negative.
// Defaults to DefaultKeyUpDelay.
func KeyUpDelay(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// DisableShadow when provided the button will not have a shadow area and will
// have no animation when pressed.
func DisableShadow() Option { _ = "STUB: not implemented"; return *new(Option) }

// DefaultTextHorizontalPadding is the default value for the HorizontalPadding option.
const DefaultTextHorizontalPadding = 1

// TextHorizontalPadding sets padding on the left and right side of the
// button's text as the amount of cells.
func TextHorizontalPadding(p int) Option { _ = "STUB: not implemented"; return *new(Option) }

// widthFor returns the required width for the specified text.
func widthFor(text string) int { _ = "STUB: not implemented"; return 0 }
