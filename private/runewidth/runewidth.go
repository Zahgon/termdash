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

// Package runewidth is a wrapper over github.com/mattn/go-runewidth which
// gives different treatment to certain runes with ambiguous width.
package runewidth

// Option is used to provide options.
type Option interface {
	// set sets the provided option.
	set(*options)
}

// options stores the provided options.
type options struct {
	runeWidths map[rune]int
}

// newOptions create a new instance of options.
func newOptions() *options { _ = "STUB: not implemented"; return nil }

// option implements Option.
type option func(*options)

// set implements Option.set.
func (o option) set(opts *options) {
	_ = "STUB: not implemented"

	// CountAsWidth overrides the default behavior, counting the specified runes as
	// the specified width. Can be provided multiple times to specify an override
	// for multiple runes.
	return
}

func CountAsWidth(r rune, width int) Option { _ = "STUB: not implemented"; return *new(Option) }

// RuneWidth returns the number of cells needed to draw r.
// Background in http://www.unicode.org/reports/tr11/.
//
// Treats runes used internally by termdash as single-cell (half-width) runes
// regardless of the locale. I.e. runes that are used to draw lines, boxes,
// indicate resize or text trimming was needed and runes used by the braille
// canvas.
//
// This should be safe, since even in locales where these runes have ambiguous
// width, we still place all the character content around them so they should
// have be half-width.
func RuneWidth(r rune, opts ...Option) int { _ = "STUB: not implemented"; return 0 }

// StringWidth is like RuneWidth, but returns the number of cells occupied by
// all the runes in the string.
func StringWidth(s string, opts ...Option) int { _ = "STUB: not implemented"; return 0 }

// inTable determines if the rune falls within the table.
// Copied from github.com/mattn/go-runewidth/blob/master/runewidth.go.
func inTable(r rune, t table) bool {
	_ = "STUB: not implemented"
	// func (t table) IncludesRune(r rune) bool {
	return false
}

type interval struct {
	first rune
	last  rune
}

type table []interval

// exceptions runes defined here are always considered to be half-width even if
// they might be ambiguous in some contexts.
var exceptions = table{
	// Characters used by termdash to indicate text trim or scroll.
	{0x2026, 0x2026},
	{0x21c4, 0x21c4},
	{0x21e7, 0x21e7},
	{0x21e9, 0x21e9},

	// Box drawing, used as line-styles.
	// https://en.wikipedia.org/wiki/Box-drawing_character
	{0x2500, 0x257F},

	// Block elements used as sparks.
	// https://en.wikipedia.org/wiki/Box-drawing_character
	{0x2580, 0x258F},
}
