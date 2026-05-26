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

// Package barchart implements a widget that draws multiple bars displaying
// values and their relative ratios.
package barchart

import (
	"image"
	"sync"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/private/canvas"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgetapi"
)

// BarChart displays multiple bars showing relative ratios of values.
//
// Each bar can have a text label under it explaining the meaning of the value
// and can display the value itself inside the bar.
//
// Implements widgetapi.Widget. This object is thread-safe.
type BarChart struct {
	// values are the values provided on a call to Values(). These are the
	// individual bars that will be drawn.
	values []int
	// max is the maximum value of a bar. A bar having this value takes all the
	// vertical space.
	max int

	// lastWidth is the width of the canvas as of the last time when Draw was called.
	lastWidth int

	// mu protects the BarChart.
	mu sync.Mutex

	// opts are the provided options.
	opts *options
}

// New returns a new BarChart.
func New(opts ...Option) (*BarChart, error) { _ = "STUB: not implemented"; return nil, nil }

// Draw draws the BarChart widget onto the canvas.
// Implements widgetapi.Widget.Draw.
func (bc *BarChart) Draw(cvs *canvas.Canvas, meta *widgetapi.Meta) error {
	_ = "STUB: not implemented"
	return nil
}

// Value might be so small so that the rectangle is zero.

// textLoc represents the location of the drawn text.
type textLoc int

const (
	insideBar textLoc = iota
	underBar
)

// drawText draws the provided text inside or under the i-th bar.
func (bc *BarChart) drawText(cvs *canvas.Canvas, i int, text string, color cell.Color, loc textLoc) error {
	_ = "STUB: not implemented"
	// Rectangle representing area in which the text will be aligned.
	return nil
}

// Align the text within the bar itself.

// Align the text within the entire column where the bar is, this
// includes the space for any label under the bar.

// barWidth determines the width of a single bar based on options and the canvas.
func (bc *BarChart) barWidth(cvs *canvas.Canvas) int { _ = "STUB: not implemented"; return 0 }

// No width when we have no values.

// Prefer width set via the options.

// barHeight determines the height of the i-th bar based on the value it is displaying.
func (bc *BarChart) barHeight(cvs *canvas.Canvas, i, value int) int {
	_ = "STUB: not implemented"
	return 0
}

// One line for the bar labels.

// barRect returns a rectangle that represents the i-th bar on the canvas that
// displays the specified value.
func (bc *BarChart) barRect(cvs *canvas.Canvas, i, value int) (image.Rectangle, error) {
	_ = "STUB: not implemented"
	return *new(image.Rectangle), nil
}

// One line for the bar labels.

// barColor safely determines the color for the i-th bar.
// Colors are optional and don't have to be specified for all the bars.
func (bc *BarChart) barColor(i int) cell.Color { _ = "STUB: not implemented"; return *new(cell.Color) }

// valColor safely determines the color for the i-th value.
// Colors are optional and don't have to be specified for all the values.
func (bc *BarChart) valColor(i int) cell.Color { _ = "STUB: not implemented"; return *new(cell.Color) }

// label safely determines the label and its color for the i-th bar.
// Labels are optional and don't have to be specified for all the bars.
func (bc *BarChart) label(i int) (string, cell.Color) {
	_ = "STUB: not implemented"
	return "", *new(cell.Color)
}

// ValueCapacity returns the number of values that can fit into the canvas.
// This is essentially the number of available cells on the canvas as observed
// on the last call to draw. Returns zero if draw wasn't called.
//
// Note that this capacity changes each time the terminal resizes, so there is
// no guarantee this remains the same next time Draw is called.
// Should be used as a hint only.
func (bc *BarChart) ValueCapacity() int { _ = "STUB: not implemented"; return 0 }

// Values sets the values to be displayed by the BarChart.
// Each value ends up in its own bar. The values must not be negative and must
// be less or equal the maximum value. A bar displaying the maximum value is a
// full bar, taking all available vertical space.
// Provided options override values set when New() was called.
func (bc *BarChart) Values(values []int, max int, opts ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

// Copy to avoid external modifications. See #174.

// Keyboard input isn't supported on the BarChart widget.
func (*BarChart) Keyboard(k *terminalapi.Keyboard, meta *widgetapi.EventMeta) error {
	_ = "STUB: not implemented"
	return nil
}

// Mouse input isn't supported on the BarChart widget.
func (*BarChart) Mouse(m *terminalapi.Mouse, meta *widgetapi.EventMeta) error {
	_ = "STUB: not implemented"
	return nil
}

// Options implements widgetapi.Widget.Options.
func (bc *BarChart) Options() widgetapi.Options {
	_ = "STUB: not implemented"
	return *new(widgetapi.Options)
}

// Request at least one cell of width from the infra, but not more even if
// we have more values. Otherwise Draw would never get called and we would
// never update bc.lastWidth and the result of ValueCapacity().
// Draw will stil refuse to draw if the canvas is too small, but the user
// will have an option to send less values.

// minBarWidth determines the minimum possible width of a bar based on the
// options.
func (bc *BarChart) minBarWidth() int { _ = "STUB: not implemented"; return 0 }

// At least one char for the bar itself.

// minSize determines the minimum required size of the canvas.
func (bc *BarChart) minSize() image.Point { _ = "STUB: not implemented"; return *new(image.Point) }

// At least one character vertically to display the bar.

// One line for the labels.

// validateValues validates the provided values and maximum.
func validateValues(values []int, max int) error { _ = "STUB: not implemented"; return nil }

// valueCapacity calculates the value capacity given the width of bars, gaps
// and canvas.
func valueCapacity(barWidth, gapWidth, cvsWidth float64) int { _ = "STUB: not implemented"; return 0 }

// values * barWidth + (values - 1) * gapWidth = cvsWidth
// values * barWidth + values * gapWidth - gapWidth = cvsWidth
// values * (barWidth + gapWidth) = cvsWidth + gapWidth
