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

package gauge

// options.go contains configurable options for Gauge.

import (
	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/private/draw"
)

// Option is used to provide options.
type Option interface {
	// set sets the provided option.
	set(*options)
}

// options holds the provided options.
type options struct {
	gaugeChar        rune
	hideTextProgress bool
	height           int
	textLabel        string
	hTextAlign       align.Horizontal
	vTextAlign       align.Vertical
	color            cell.Color
	filledTextColor  cell.Color
	emptyTextColor   cell.Color
	// If set, draws a border around the gauge.
	border            linestyle.LineStyle
	borderCellOpts    []cell.Option
	borderTitle       string
	borderTitleHAlign align.Horizontal
	// If set draws a vertical line representing the threshold.
	threshold          int
	thresholdCellOpts  []cell.Option
	thresholdLineStyle linestyle.LineStyle
}

// newOptions returns options with the default values set.
func newOptions() *options { _ = "STUB: not implemented"; return nil }

// validate validates the provided options.
func (o *options) validate() error { _ = "STUB: not implemented"; return nil }

// option implements Option.
type option func(*options)

// set implements Option.set.
func (o option) set(opts *options) {
	_ = "STUB: not implemented"

	// DefaultChar is the default value for the Char option.
	return
}

const DefaultChar = draw.DefaultRectChar

// Char sets the rune that is used when drawing the rectangle representing the
// current progress.
func Char(ch rune) Option { _ = "STUB: not implemented"; return *new(Option) }

// ShowTextProgress configures the Gauge so that it also displays a text
// enumerating the progress. This is the default behavior.
// If the progress is set by a call to Percent(), the displayed text will show
// the percentage, e.g. "50%". If the progress is set by a call to Absolute(),
// the displayed text will whos the absolute numbers, e.g. "5/10".
func ShowTextProgress() Option { _ = "STUB: not implemented"; return *new(Option) }

// HideTextProgress disables the display of a text enumerating the progress.
func HideTextProgress() Option { _ = "STUB: not implemented"; return *new(Option) }

// Height sets the height of the drawn Gauge. Must be a positive number.
// Defaults to zero which means the height of the container.
func Height(height int) Option { _ = "STUB: not implemented"; return *new(Option) }

// TextLabel configures the Gauge to display the provided text.
// If the ShowTextProgress() option is also provided, this label is drawn right
// after the progress text.
func TextLabel(text string) Option { _ = "STUB: not implemented"; return *new(Option) }

// DefaultColor is the default value for the Color option.
const DefaultColor = cell.ColorGreen

// Color sets the color of the gauge.
func Color(c cell.Color) Option { _ = "STUB: not implemented"; return *new(Option) }

// DefaultFilledTextColor is the default value for the FilledTextColor option.
const DefaultFilledTextColor = cell.ColorBlack

// FilledTextColor sets color of the text progress and text label for the
// portion of the text that falls within the filled up part of the Gauge. I.e.
// text on the Gauge.
func FilledTextColor(c cell.Color) Option { _ = "STUB: not implemented"; return *new(Option) }

// DefaultEmptyTextColor is the default value for the EmptyTextColor option.
const DefaultEmptyTextColor = cell.ColorDefault

// EmptyTextColor sets color of the text progress and text label for the
// portion of the text that falls outside the filled up part of the Gauge. I.e.
// text in the empty area the Gauge didn't fill yet.
func EmptyTextColor(c cell.Color) Option { _ = "STUB: not implemented"; return *new(Option) }

// DefaultHorizontalTextAlign is the default value for the HorizontalTextAlign option.
const DefaultHorizontalTextAlign = align.HorizontalCenter

// HorizontalTextAlign sets the horizontal alignment of the text progress and
// text label.
func HorizontalTextAlign(h align.Horizontal) Option { _ = "STUB: not implemented"; return *new(Option) }

// DefaultVerticalTextAlign is the default value for the VerticalTextAlign option.
const DefaultVerticalTextAlign = align.VerticalMiddle

// VerticalTextAlign sets the vertical alignment of the text progress and
// text label.
func VerticalTextAlign(v align.Vertical) Option { _ = "STUB: not implemented"; return *new(Option) }

// Border configures the gauge to have a border of the specified style.
func Border(ls linestyle.LineStyle, cOpts ...cell.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// BorderTitle sets a text title within the border.
func BorderTitle(title string) Option { _ = "STUB: not implemented"; return *new(Option) }

// BorderTitleAlign sets the horizontal alignment for the border title.
// Defaults to alignment on the left.
func BorderTitleAlign(h align.Horizontal) Option { _ = "STUB: not implemented"; return *new(Option) }

// Threshold configures the Gauge to display a vertical threshold line at value
// t. If the progress is set by a call to Percent(), t represents a percentage,
// e.g. "40" means line is displayed at 40%. If the progress is set by a call to
// Absolute(), the threshold is considered an absolute number.
// Threshold must be positive to be displayed. If the threshold is zero or
// greater than total, it won't be displayed. Defaults to zero.
func Threshold(t int, ls linestyle.LineStyle, cOpts ...cell.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
