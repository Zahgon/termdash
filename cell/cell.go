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

// Package cell implements cell options and attributes.
package cell

// Option is used to provide options for cells on a 2-D terminal.
type Option interface {
	// Set sets the provided option.
	Set(*Options)
}

// Options stores the provided options.
type Options struct {
	FgColor       Color
	BgColor       Color
	Bold          bool
	Italic        bool
	Underline     bool
	Strikethrough bool
	Inverse       bool
	Blink         bool
	Dim           bool
}

// Set allows existing options to be passed as an option.
func (o *Options) Set(other *Options) {
	_ = "STUB: not implemented"

	// NewOptions returns a new Options instance after applying the provided options.
	return
}

func NewOptions(opts ...Option) *Options { _ = "STUB: not implemented"; return nil }

// option implements Option.
type option func(*Options)

// Set implements Option.set.
func (co option) Set(opts *Options) {
	_ = "STUB: not implemented"

	// FgColor sets the foreground color of the cell.
	return
}

func FgColor(color Color) Option { _ = "STUB: not implemented"; return *new(Option) }

// BgColor sets the background color of the cell.
func BgColor(color Color) Option { _ = "STUB: not implemented"; return *new(Option) }

// Bold makes cell's text bold.
func Bold() Option { _ = "STUB: not implemented"; return *new(Option) }

// Italic makes cell's text italic. Only works when using the tcell backend.
func Italic() Option { _ = "STUB: not implemented"; return *new(Option) }

// Underline makes cell's text underlined.
func Underline() Option { _ = "STUB: not implemented"; return *new(Option) }

// Strikethrough strikes through the cell's text. Only works when using the tcell backend.
func Strikethrough() Option { _ = "STUB: not implemented"; return *new(Option) }

// Inverse inverts the colors of the cell's text.
func Inverse() Option { _ = "STUB: not implemented"; return *new(Option) }

// Blink makes the cell's text blink. Only works when using the tcell backend.
func Blink() Option { _ = "STUB: not implemented"; return *new(Option) }

// Dim makes the cell foreground color dim. Only works when using the tcell backend.
func Dim() Option { _ = "STUB: not implemented"; return *new(Option) }
