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

// Package gauge implements a widget that displays the progress of an operation.
package gauge

import (
	"image"
	"sync"

	"github.com/mum4k/termdash/private/canvas"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgetapi"
)

// progressType indicates how was the current progress provided by the caller.
type progressType int

// String implements fmt.Stringer()
func (pt progressType) String() string { _ = "STUB: not implemented"; return "" }

// progressTypeNames maps progressType values to human readable names.
var progressTypeNames = map[progressType]string{
	progressTypePercent:  "progressTypePercent",
	progressTypeAbsolute: "progressTypeAbsolute",
}

const (
	progressTypePercent = iota
	progressTypeAbsolute
)

// Gauge displays the progress of an operation.
//
// Draws a rectangle, a progress bar with optional display of percentage and /
// or text label.
//
// Implements widgetapi.Widget. This object is thread-safe.
type Gauge struct {
	// pt indicates how current and total are interpreted.
	pt progressType
	// current is the current progress that will be drawn.
	current int
	// total is the value that represents completion.
	// For progressTypePercent, this is 100, for progressTypeAbsolute this is
	// the total provided by the caller.
	total int
	// mu protects the Gauge.
	mu sync.Mutex

	// opts are the provided options.
	opts *options
}

// New returns a new Gauge.
func New(opts ...Option) (*Gauge, error) { _ = "STUB: not implemented"; return nil, nil }

// Absolute sets the progress in absolute numbers, i.e. 7 out of 10.
// The total amount must be a non-zero positive integer. The done amount must
// be a zero or a positive integer such that done <= total.
// Provided options override values set when New() was called.
func (g *Gauge) Absolute(done, total int, opts ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

// Percent sets the current progress in percentage.
// The provided value must be between 0 and 100.
// Provided options override values set when New() was called.
func (g *Gauge) Percent(p int, opts ...Option) error { _ = "STUB: not implemented"; return nil }

// width determines the X coordinate that represents point w in rectangle ar.
// This is used to calculate the width of the gauge drawn on the provided area
// in order to represent the current progress or to figure out the coordinate
// for the threshold line.
func (g *Gauge) width(ar image.Rectangle, w int) int { _ = "STUB: not implemented"; return 0 }

// hasBorder determines of the gauge has a border.
func (g *Gauge) hasBorder() bool { _ = "STUB: not implemented"; return false }

// usable determines the usable area for the gauge itself.
func (g *Gauge) usable(cvs *canvas.Canvas) image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

// thresholdVisible determines if the threshold line should be drawn.
func (g *Gauge) thresholdVisible() bool { _ = "STUB: not implemented"; return false }

// progressText returns the textual representation of the current progress.
func (g *Gauge) progressText() string { _ = "STUB: not implemented"; return "" }

// gaugeText returns full text to be displayed within the gauge, i.e. the
// progress text and the optional label.
func (g *Gauge) gaugeText() string { _ = "STUB: not implemented"; return "" }

// drawText draws the text enumerating the progress and the text label.
func (g *Gauge) drawText(cvs *canvas.Canvas, progress image.Rectangle) error {
	_ = "STUB: not implemented"
	return nil
}

// If the current rune is full-width and only one of its cells falls
// within the filled area of the gauge, extend the gauge by one cell to
// fully cover the full-width rune.

// drawThreshold draws the threshold line.
func (g *Gauge) drawThreshold(cvs *canvas.Canvas) error { _ = "STUB: not implemented"; return nil }

// Draw draws the Gauge widget onto the canvas.
// Implements widgetapi.Widget.Draw.
func (g *Gauge) Draw(cvs *canvas.Canvas, meta *widgetapi.Meta) error {
	_ = "STUB: not implemented"
	return nil
}

// Keyboard input isn't supported on the Gauge widget.
func (g *Gauge) Keyboard(k *terminalapi.Keyboard, meta *widgetapi.EventMeta) error {
	_ = "STUB: not implemented"
	return nil
}

// Mouse input isn't supported on the Gauge widget.
func (g *Gauge) Mouse(m *terminalapi.Mouse, meta *widgetapi.EventMeta) error {
	_ = "STUB: not implemented"
	return nil
}

// maxSize determines the maximum size of the canvas.
func (g *Gauge) maxSize() image.Point { _ = "STUB: not implemented"; return *new(image.Point) }

// Add the required space for the border.

// minSize determines the minimum required size of the canvas.
func (g *Gauge) minSize() image.Point {
	_ = "STUB: not implemented"
	// Shorter gauge than this cannot display anything.
	return *new(image.Point)
}

// At least one line for the gauge itself.

// Add the required space for the border.

// Options implements widgetapi.Widget.Options.
func (g *Gauge) Options() widgetapi.Options {
	_ = "STUB: not implemented"
	return *new(widgetapi.Options)
}
