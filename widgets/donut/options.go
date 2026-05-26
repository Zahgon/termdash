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

package donut

// options.go contains configurable options for Donut.

import (
	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/cell"
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
	donutHolePercent int
	hideTextProgress bool

	textCellOpts []cell.Option
	cellOpts     []cell.Option

	labelCellOpts []cell.Option
	labelAlign    align.Horizontal
	label         string

	// The angle in degrees that represents 0 and 100% of the progress.
	startAngle int
	// The direction in which the donut completes as progress increases.
	// Positive for counter-clockwise, negative for clockwise.
	direction int
}

// validate validates the provided options.
func (o *options) validate() error { _ = "STUB: not implemented"; return nil }

// newOptions returns options with the default values set.
func newOptions() *options { _ = "STUB: not implemented"; return nil }

// DefaultHolePercent is the default value for the HolePercent
// option.
const DefaultHolePercent = 35

// HolePercent sets the size of the "hole" inside the donut as a
// percentage of the donut's radius.
// Setting this to zero disables the hole so that the donut will become just a
// circle. Valid range is 0 <= p <= 100.
func HolePercent(p int) Option { _ = "STUB: not implemented"; return *new(Option) }

// ShowTextProgress configures the Gauge so that it also displays a text
// enumerating the progress. This is the default behavior.
// If the progress is set by a call to Percent(), the displayed text will show
// the percentage, e.g. "50%". If the progress is set by a call to Absolute(),
// the displayed text will those the absolute numbers, e.g. "5/10".
//
// The progress is only displayed if there is enough space for it in the middle
// of the drawn donut.
//
// Providing this option also sets HolePercent to its default value.
func ShowTextProgress() Option { _ = "STUB: not implemented"; return *new(Option) }

// HideTextProgress disables the display of a text enumerating the progress.
func HideTextProgress() Option { _ = "STUB: not implemented"; return *new(Option) }

// TextCellOpts sets cell options on cells that contain the displayed text
// progress.
func TextCellOpts(cOpts ...cell.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

// CellOpts sets cell options on cells that contain the donut.
func CellOpts(cOpts ...cell.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

// DefaultStartAngle is the default value for the StartAngle option.
const DefaultStartAngle = 90

// StartAngle sets the starting angle in degrees, i.e. the point that will
// represent both 0% and 100% of progress.
// Valid values are in range 0 <= angle < 360.
// Angles start at the X axis and grow counter-clockwise.
func StartAngle(angle int) Option { _ = "STUB: not implemented"; return *new(Option) }

// Clockwise sets the donut widget for a progression in the clockwise
// direction. This is the default option.
func Clockwise() Option { _ = "STUB: not implemented"; return *new(Option) }

// CounterClockwise sets the donut widget for a progression in the counter-clockwise
// direction.
func CounterClockwise() Option { _ = "STUB: not implemented"; return *new(Option) }

// Label sets a text label to be displayed under the donut.
func Label(text string, cOpts ...cell.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// DefaultLabelAlign is the default value for the LabelAlign option.
const DefaultLabelAlign = align.HorizontalCenter

// LabelAlign sets the alignment of the label under the donut.
func LabelAlign(la align.Horizontal) Option { _ = "STUB: not implemented"; return *new(Option) }
