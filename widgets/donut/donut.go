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

// Package donut is a widget that displays the progress of an operation as a
// partial or full circle.
package donut

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

// Donut displays the progress of an operation by filling a partial circle and
// eventually by completing a full circle. The circle can have a "hole" in the
// middle, which is where the name comes from.
//
// Implements widgetapi.Widget. This object is thread-safe.
type Donut struct {
	// pt indicates how current and total are interpreted.
	pt progressType
	// current is the current progress that will be drawn.
	current int
	// total is the value that represents completion.
	// For progressTypePercent, this is 100, for progressTypeAbsolute this is
	// the total provided by the caller.
	total int
	// mu protects the Donut.
	mu sync.Mutex

	// opts are the provided options.
	opts *options
}

// New returns a new Donut.
func New(opts ...Option) (*Donut, error) { _ = "STUB: not implemented"; return nil, nil }

// Absolute sets the progress in absolute numbers, e.g. 7 out of 10.
// The total amount must be a non-zero positive integer. The done amount must
// be a zero or a positive integer such that done <= total.
// Provided options override values set when New() was called.
func (d *Donut) Absolute(done, total int, opts ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

// Percent sets the current progress in percentage.
// The provided value must be between 0 and 100.
// Provided options override values set when New() was called.
func (d *Donut) Percent(p int, opts ...Option) error { _ = "STUB: not implemented"; return nil }

// progressText returns the textual representation of the current progress.
func (d *Donut) progressText() string { _ = "STUB: not implemented"; return "" }

// holeRadius calculates the radius of the "hole" in the donut.
// Returns zero if no hole should be drawn.
func (d *Donut) holeRadius(donutRadius int) int { _ = "STUB: not implemented"; return 0 }

// Smallest possible circle radius.

// drawText draws the text label showing the progress.
// The text is only drawn if the radius of the donut "hole" is large enough to
// accommodate it.
// The mid point addresses coordinates in pixels on a braille canvas.
func (d *Donut) drawText(cvs *canvas.Canvas, mid image.Point, holeR int) error {
	_ = "STUB: not implemented"
	return nil
}

// drawLabel draws the text label in the area.
func (d *Donut) drawLabel(cvs *canvas.Canvas, labelAr image.Rectangle) error {
	_ = "STUB: not implemented"
	return nil
}

// Draw draws the Donut widget onto the canvas.
// Implements widgetapi.Widget.Draw.
func (d *Donut) Draw(cvs *canvas.Canvas, meta *widgetapi.Meta) error {
	_ = "STUB: not implemented"
	return nil
}

// No progress recorded, so nothing to do.

// Reserving area for the label might have resulted in donutAr being
// too small.

// Keyboard input isn't supported on the Donut widget.
func (*Donut) Keyboard(k *terminalapi.Keyboard, meta *widgetapi.EventMeta) error {
	_ = "STUB: not implemented"
	return nil
}

// Mouse input isn't supported on the Donut widget.
func (*Donut) Mouse(m *terminalapi.Mouse, meta *widgetapi.EventMeta) error {
	_ = "STUB: not implemented"
	return nil
}

// minSize is the smallest area we can draw donut on.
var minSize = image.Point{3, 3}

// Options implements widgetapi.Widget.Options.
func (d *Donut) Options() widgetapi.Options {
	_ = "STUB: not implemented"
	return *

	// We are drawing a circle, ensure equal ratio of rows and columns.
	// This is adjusted for the inequality of the braille canvas.
	new(widgetapi.Options)
}

// The smallest circle that "looks" like a circle on the canvas.

// donutAndLabel splits the canvas area into an area for the donut and an
// area under the donut for the text label.
func donutAndLabel(cvsAr image.Rectangle) (donAr, labelAr image.Rectangle, err error) {
	_ = "STUB: not implemented"
	return *

	// Two lines for the text label at the bottom.
	// One for the text itself and one for visual space between the donut and
	// the label.
	new(image.Rectangle), *new(image.Rectangle), nil
}
