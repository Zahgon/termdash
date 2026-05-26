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

package segmentdisplay

import (
	"github.com/mum4k/termdash/align"
)

// options.go contains configurable options for SegmentDisplay.

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
	hAlign          align.Horizontal
	vAlign          align.Vertical
	maximizeSegSize bool
	gapPercent      int
}

// validate validates the provided options.
func (o *options) validate() error { _ = "STUB: not implemented"; return nil }

// newOptions returns options with the default values set.
func newOptions() *options { _ = "STUB: not implemented"; return nil }

// AlignHorizontal sets the horizontal alignment for the individual display
// segments. Defaults to alignment in the center.
func AlignHorizontal(h align.Horizontal) Option { _ = "STUB: not implemented"; return *new(Option) }

// AlignVertical sets the vertical alignment for the individual display
// segments. Defaults to alignment in the middle
func AlignVertical(v align.Vertical) Option { _ = "STUB: not implemented"; return *new(Option) }

// MaximizeSegmentHeight tells the widget to maximize the height of the
// individual display segments.
// When this option is set and the user has provided more text than we can fit
// on the canvas, the widget will prefer to maximize height of individual
// characters which will result in earlier trimming of the text.
func MaximizeSegmentHeight() Option { _ = "STUB: not implemented"; return *new(Option) }

// MaximizeDisplayedText tells the widget to maximize the amount of characters
// that are displayed.
// When this option is set and the user has provided more text than we can fit
// on the canvas, the widget will prefer to decrease the height of individual
// characters and fit more of them on the canvas.
// This is the default behavior.
func MaximizeDisplayedText() Option { _ = "STUB: not implemented"; return *new(Option) }

// DefaultGapPercent is the default value for the GapPercent option.
const DefaultGapPercent = 20

// GapPercent sets the size of the horizontal gap between individual segments
// (characters) expressed as a percentage of the segment height.
func GapPercent(perc int) Option { _ = "STUB: not implemented"; return *new(Option) }
