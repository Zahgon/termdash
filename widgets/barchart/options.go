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

package barchart

// options.go contains configurable options for BarChart.

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/private/draw"
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
	barChar     rune
	barWidth    int
	barGap      int
	showValues  bool
	barColors   []cell.Color
	labelColors []cell.Color
	valueColors []cell.Color
	labels      []string
}

// validate validates the provided options.
func (o *options) validate() error { _ = "STUB: not implemented"; return nil }

// newOptions returns options with the default values set.
func newOptions() *options { _ = "STUB: not implemented"; return nil }

// DefaultChar is the default value for the Char option.
const DefaultChar = draw.DefaultRectChar

// Char sets the rune that is used when drawing the rectangle representing the
// bars.
func Char(ch rune) Option { _ = "STUB: not implemented"; return *new(Option) }

// BarWidth sets the width of the bars. If not set, or set to zero, the bars
// use all the space available to the widget. Must be a positive or zero
// integer.
func BarWidth(width int) Option { _ = "STUB: not implemented"; return *new(Option) }

// DefaultBarGap is the default value for the BarGap option.
const DefaultBarGap = 1

// BarGap sets the width of the space between the bars.
// Must be a positive or zero integer.
// Defaults to DefaultBarGap.
func BarGap(width int) Option { _ = "STUB: not implemented"; return *new(Option) }

// ShowValues tells the bar chart to display the actual values inside each of the bars.
func ShowValues() Option { _ = "STUB: not implemented"; return *new(Option) }

// DefaultBarColor is the default color of a bar, unless specified otherwise
// via the BarColors option.
const DefaultBarColor = cell.ColorRed

// BarColors sets the colors of each of the bars.
// Bars are created on a call to Values(), each value ends up in its own Bar.
// The first supplied color applies to the bar displaying the first value.
// Any bars that don't have a color specified use the DefaultBarColor.
func BarColors(colors []cell.Color) Option { _ = "STUB: not implemented"; return *new(Option) }

// DefaultLabelColor is the default color of a bar label, unless specified
// otherwise via the LabelColors option.
const DefaultLabelColor = cell.ColorGreen

// LabelColors sets the colors of each of the labels under the bars.
// Bars are created on a call to Values(), each value ends up in its own Bar.
// The first supplied color applies to the label of the bar displaying the
// first value. Any labels that don't have a color specified use the
// DefaultLabelColor.
func LabelColors(colors []cell.Color) Option { _ = "STUB: not implemented"; return *new(Option) }

// Labels sets the labels displayed under each bar,
// Bars are created on a call to Values(), each value ends up in its own Bar.
// The first supplied label applies to the bar displaying the first value.
// If not specified, the corresponding bar (or all the bars) don't have a
// label.
func Labels(labels []string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Copy to avoid external modifications. See #174.

// DefaultValueColor is the default color of a bar value, unless specified
// otherwise via the ValueColors option.
const DefaultValueColor = cell.ColorYellow

// ValueColors sets the colors of each of the values in the bars. Bars are
// created on a call to Values(), each value ends up in its own Bar. The first
// supplied color applies to the bar displaying the first value. Any values
// that don't have a color specified use the DefaultValueColor.
func ValueColors(colors []cell.Color) Option { _ = "STUB: not implemented"; return *new(Option) }
