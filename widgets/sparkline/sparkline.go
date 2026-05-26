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

// Package sparkline is a widget that draws a graph showing a series of values as vertical bars.
package sparkline

import (
	"image"
	"sync"

	"github.com/mum4k/termdash/private/canvas"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgetapi"
)

// SparkLine draws a graph showing a series of values as vertical bars.
//
// Bars can have sub-cell height. The graphs scale adjusts dynamically based on
// the largest visible value.
//
// Implements widgetapi.Widget. This object is thread-safe.
type SparkLine struct {
	// data are the data points the SparkLine displays.
	data []int

	// lastWidth is the width of the canvas as of the last time when Draw was called.
	lastWidth int

	// mu protects the SparkLine.
	mu sync.Mutex

	// opts are the provided options.
	opts *options
}

// New returns a new SparkLine.
func New(opts ...Option) (*SparkLine, error) { _ = "STUB: not implemented"; return nil, nil }

// Draw draws the SparkLine widget onto the canvas.
// Implements widgetapi.Widget.Draw.
func (sl *SparkLine) Draw(cvs *canvas.Canvas, meta *widgetapi.Meta) error {
	_ = "STUB: not implemented"
	return nil
}

// Last spark represents full cell.

// Label is placed immediately above the SparkLine.

// ValueCapacity returns the number of values that can fit into the canvas.
// This is essentially the number of available cells on the canvas as observed
// on the last call to draw. Returns zero if draw wasn't called.
//
// Note that this capacity changes each time the terminal resizes, so there is
// no guarantee this remains the same next time Draw is called.
// Should be used as a hint only.
func (sl *SparkLine) ValueCapacity() int { _ = "STUB: not implemented"; return 0 }

// Add adds data points to the SparkLine.
// Each data point is represented by one bar on the SparkLine. Zero value data
// points are valid and are represented by an empty space on the SparkLine
// (i.e. a missing bar).
//
// At least one data point must be provided. All data points must be positive
// integers.
//
// The last added data point will be the one displayed all the way on the right
// of the SparkLine. If there are more data points than we can fit bars to the
// width of the SparkLine, only the last n data points that fit will be
// visible.
//
// Provided options override values set when New() was called.
func (sl *SparkLine) Add(data []int, opts ...Option) error { _ = "STUB: not implemented"; return nil }

// Clear removes all the data points in the SparkLine, effectively returning to
// an empty graph.
func (sl *SparkLine) Clear() { _ = "STUB: not implemented"; return }

// Keyboard input isn't supported on the SparkLine widget.
func (*SparkLine) Keyboard(k *terminalapi.Keyboard, meta *widgetapi.EventMeta) error {
	_ = "STUB: not implemented"
	return nil
}

// Mouse input isn't supported on the SparkLine widget.
func (*SparkLine) Mouse(m *terminalapi.Mouse, meta *widgetapi.EventMeta) error {
	_ = "STUB: not implemented"
	return nil
}

// area returns the area of the canvas available to the SparkLine.
func (sl *SparkLine) area(cvs *canvas.Canvas) image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

// Height is determined based on options (fixed height / label).

// Reserve one line for the label.

// minSize returns the minimum canvas size for the SparkLine based on the options.
func (sl *SparkLine) minSize() image.Point {
	_ = "STUB: not implemented"
	// At least one data point.
	return *new(image.Point)
}

// At least one line of characters.

// One line for the text label.

// Options implements widgetapi.Widget.Options.
func (sl *SparkLine) Options() widgetapi.Options {
	_ = "STUB: not implemented"
	return *new(widgetapi.Options)
}

// Fix the height to the one specified.
