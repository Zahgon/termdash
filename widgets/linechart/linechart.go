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

// Package linechart contains a widget that displays line charts.
package linechart

import (
	"image"
	"sync"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/private/canvas"
	"github.com/mum4k/termdash/private/canvas/braille"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgetapi"
	"github.com/mum4k/termdash/widgets/linechart/internal/axes"
	"github.com/mum4k/termdash/widgets/linechart/internal/zoom"
)

// seriesValues represent values stored in the series.
type seriesValues struct {
	// values are the values in the series.
	values []float64
	// min is the smallest value, zero if values is empty.
	min float64
	// max is the largest value, zero if values is empty.
	max float64

	seriesCellOpts []cell.Option
	// The custom labels provided on a call to Series and a bool indicating if
	// the labels were provided. This allows resetting them to nil.
	xLabelsSet bool
	xLabels    map[int]string
}

// newSeriesValues returns a new seriesValues instance.
func newSeriesValues(values []float64) *seriesValues {
	_ = "STUB: not implemented"
	// Copy to avoid external modifications. See #174.
	return nil
}

// LineChart draws line charts.
//
// Each line chart has an identifying label and a set of values that are
// plotted.
//
// The size of the two axes is determined from the values.
// The X axis will have a number of evenly distributed data points equal to the
// largest count of values among all the labeled line charts.
// The Y axis will be sized so that it can conveniently accommodate the largest
// value among all the labeled line charts. This determines the used scale.
//
// LineChart supports mouse based zoom, zooming is achieved by either
// highlighting an area on the graph (left mouse clicking and dragging) or by
// using the mouse scroll button.
//
// Implements widgetapi.Widget. This object is thread-safe.
type LineChart struct {
	// mu protects the LineChart widget.
	mu sync.RWMutex

	// series are the series that will be plotted.
	// Keyed by the name of the series and updated by calling Series.
	series map[string]*seriesValues

	// yMin are the min and max values for the Y axis.
	yMin, yMax float64

	// capacity is the last observed value capacity in pixels when Draw was
	// called.
	capacity int

	// opts are the provided options.
	opts *options

	// xLabels that were provided on a call to Series.
	xLabels map[int]string

	// zoom tracks the zooming of the X axis.
	zoom *zoom.Tracker
}

// New returns a new line chart widget.
func New(opts ...Option) (*LineChart, error) { _ = "STUB: not implemented"; return nil, nil }

// SeriesOption is used to provide options to Series.
type SeriesOption interface {
	// set sets the provided option.
	set(*seriesValues)
}

// seriesOption implements SeriesOption.
type seriesOption func(*seriesValues)

// set implements SeriesOption.set.
func (so seriesOption) set(sv *seriesValues) {
	_ = "STUB: not implemented"

	// SeriesCellOpts sets the cell options for this series.
	// Note that the braille canvas has resolution of 2x4 pixels per cell, but each
	// cell can only have one set of cell options set. Meaning that where series
	// share a cell, the last drawn series sets the cell options. Series are drawn
	// in alphabetical order based on their name.
	return
}

func SeriesCellOpts(co ...cell.Option) SeriesOption {
	_ = "STUB: not implemented"
	return *new(SeriesOption)
}

// SeriesXLabels is used to provide custom labels for the X axis.
// The argument maps the positions in the provided series to the desired label.
// The labels are only used if they fit under the axis.
// Custom labels are property of the line chart, since there is only one X axis,
// providing multiple custom labels overwrites the previous value.
func SeriesXLabels(labels map[int]string) SeriesOption {
	_ = "STUB: not implemented"
	return *new(SeriesOption)
}

// Copy to avoid external modifications. See #174.

// yMinMax determines the min and max values for the Y axis.
func (lc *LineChart) yMinMax() (float64, float64) { _ = "STUB: not implemented"; return 0, 0 }

// ValueCapacity returns the number of values that could be fit onto the X axis
// without a need to rescale the X axis. This is essentially the number of
// available pixels on the braille canvas based on the width of the LineChart
// as observed on the last call to draw. Returns zero if draw wasn't called.
//
// Note that this capacity changes each time the terminal resizes, so there is
// no guarantee this remains the same next time Draw is called.
// Should be used as a hint only.
func (lc *LineChart) ValueCapacity() int { _ = "STUB: not implemented"; return 0 }

// Series sets the values that should be displayed as the line chart with the
// provided label.
// The values that should not be displayed on the line chart should be represented
// as math.NaN values on the values slice.
// Subsequent calls with the same label replace any previously provided values.
func (lc *LineChart) Series(label string, values []float64, opts ...SeriesOption) error {
	_ = "STUB: not implemented"
	return nil
}

// xDetails returns the details for the X axis given the specified minimum and
// maximum value to display.
func (lc *LineChart) xDetails(cvs *canvas.Canvas, reqYWidth, min, max int) (*axes.XDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// xDetailsForCap adjusts the X details according to the capacity of the
// braille canvas (how many values can it fit).
// If the capacity cannot accommodate all the values, the starting value of the
// X axis is adjusted so that it displays the last n values that fit.
// Returns unadjusted xd if all the values fit.
func (lc *LineChart) xDetailsForCap(cvs *canvas.Canvas, bc *braille.Canvas, xd *axes.XDetails, yd *axes.YDetails) (*axes.XDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// axesDetails determines the details about the X and Y axes.
func (lc *LineChart) axesDetails(cvs *canvas.Canvas) (*axes.XDetails, *axes.YDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Draw draws the values as line charts.
// Implements widgetapi.Widget.Draw.
func (lc *LineChart) Draw(cvs *canvas.Canvas, meta *widgetapi.Meta) error {
	_ = "STUB: not implemented"
	return nil
}

// drawAxes draws the X,Y axes and their labels.
func (lc *LineChart) drawAxes(cvs *canvas.Canvas, xd *axes.XDetails, yd *axes.YDetails) error {
	_ = "STUB: not implemented"
	return nil
}

// graphAr returns the area available for the graph itself sized so that it
// fits between the axes and the canvas borders.
func (lc *LineChart) graphAr(cvs *canvas.Canvas, xd *axes.XDetails, yd *axes.YDetails) image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

// drawSeries draws the graph representing the stored series.
// Returns XDetails that might be adjusted to not start at zero value if some
// of the series didn't fit the graphs and XAxisUnscaled was provided.
// If the series has NaN values they will be ignored and not draw on the graph.
func (lc *LineChart) drawSeries(cvs *canvas.Canvas, xd *axes.XDetails, yd *axes.YDetails) (*axes.XDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Skip over series that don't have at least two points since we can't
// draw a line for just one point.
// Skip over series that fall under the minimum value on the X axis.

// Skip the values that are missing.

// Don't draw lines for values that aren't supposed to be visible.
// These are either values outside of the current zoom or
// values at the beginning of a series that falls before athe
// start of an unscaled X axis when the XAxisUnscaled option is
// provided.

// highlightRange highlights the range of X columns on the braille canvas.
func (lc *LineChart) highlightRange(bc *braille.Canvas, hRange *zoom.Range) error {
	_ = "STUB: not implemented"
	return nil
}

// Keyboard implements widgetapi.Widget.Keyboard.
func (lc *LineChart) Keyboard(k *terminalapi.Keyboard, meta *widgetapi.EventMeta) error {
	_ = "STUB: not implemented"
	return nil
}

// Mouse implements widgetapi.Widget.Mouse.
func (lc *LineChart) Mouse(m *terminalapi.Mouse, meta *widgetapi.EventMeta) error {
	_ = "STUB: not implemented"
	return nil
}

// minSize determines the minimum required size to draw the line chart.
func (lc *LineChart) minSize() image.Point {
	_ = "STUB: not implemented"
	// At the very least we need:
	// - n cells width for the Y axis and its labels as reported by it.
	// - at least 1 cell width for the graph.
	return *new(image.Point)
}

// And for the height:
// - n cells width for the X axis and its labels as reported by it.
// - at least 2 cell height for the graph.

// Options implements widgetapi.Widget.Options.
func (lc *LineChart) Options() widgetapi.Options {
	_ = "STUB: not implemented"
	return *new(widgetapi.Options)
}

// maxXValue returns the maximum value on the X axis among all the series.
// lc.mu must be held when calling this method.
func (lc *LineChart) maxXValue() int { _ = "STUB: not implemented"; return 0 }

// minMax is a wrapper around numbers.MinMax that controls
// the output if the values are NaN and sets defaults if it's
// the case.
func minMax(values []float64) (x, y float64) { _ = "STUB: not implemented"; return 0, 0 }
