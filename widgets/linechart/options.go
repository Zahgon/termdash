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

package linechart

import (
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/linechart/internal/axes"
)

// options.go contains configurable options for LineChart.

// Option is used to provide options to New().
type Option interface {
	// set sets the provided option.
	set(*options)
}

// options stores the provided options.
type options struct {
	axesCellOpts        []cell.Option
	xLabelCellOpts      []cell.Option
	xLabelOrientation   axes.LabelOrientation
	yLabelCellOpts      []cell.Option
	xAxisUnscaled       bool
	yAxisMode           axes.YScaleMode
	yAxisCustomScale    *customScale
	yAxisValueFormatter ValueFormatter
	zoomHightlightColor cell.Color
	zoomStepPercent     int
}

// validate validates the provided options.
func (o *options) validate() error { _ = "STUB: not implemented"; return nil }

// newOptions returns a new options instance.
func newOptions(opts ...Option) *options { _ = "STUB: not implemented"; return nil }

// option implements Option.
type option func(*options)

// set implements Option.set.
func (o option) set(opts *options) {
	_ = "STUB: not implemented"

	// AxesCellOpts set the cell options for the X and Y axes.
	return
}

func AxesCellOpts(co ...cell.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

// XLabelCellOpts set the cell options for the labels on the X axis.
func XLabelCellOpts(co ...cell.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

// XLabelsVertical makes the labels under the X axis flow vertically.
// Defaults to labels that flow horizontally.
func XLabelsVertical() Option { _ = "STUB: not implemented"; return *new(Option) }

// XLabelsHorizontal makes the labels under the X axis flow horizontally.
// This is the default option.
func XLabelsHorizontal() Option { _ = "STUB: not implemented"; return *new(Option) }

// YLabelCellOpts set the cell options for the labels on the Y axis.
func YLabelCellOpts(co ...cell.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

// YAxisAdaptive makes the Y axis adapt its base value depending on the
// provided series.
// Without this option, the Y axis always starts at the zero value regardless of
// values available in the series.
// When this option is specified and the series don't contain value zero, the Y
// axis will be adapted to the minimum value for all-positive series or the
// maximum value for all-negative series. The Y axis still starts at the zero
// value if the series contain both positive and negative values.
func YAxisAdaptive() Option { _ = "STUB: not implemented"; return *new(Option) }

// customScale is the custom scale provided via the YAxisCustomScale option.
type customScale struct {
	min, max float64
}

// YAxisCustomScale when provided, the scale of the Y axis will be based on the
// specified minimum and maximum value instead of determining those from the
// LineChart series. Useful to visually stabilize the Y axis for LineChart
// applications that continuously feed values.
// The default behavior is to continuously determine the minimum and maximum
// value from the series before drawing the LineChart.
// Even when this option is provided, the LineChart would still rescale the Y
// axis if a value is encountered that is outside of the range specified here,
// i.e. smaller than the minimum or larger than the maximum.
// Both the minimum and the maximum must be valid numbers and the minimum must
// be smaller than the maximum.
//
// Providing this option also sets YAxisAdaptive.
func YAxisCustomScale(min, max float64) Option { _ = "STUB: not implemented"; return *new(Option) }

// XAxisUnscaled when provided, stops the LineChart from rescaling the X axis
// when it can't fit all the values in the series, instead the LineCharts only
// displays the last n values that fit into its width. This is useful to create
// an impression of values rolling through the linechart right to left. Note
// that this results in hiding of values from the beginning of the series
// that didn't fit completely and might hide some shorter series as this
// effectively makes the X axis start at a non-zero value.
//
// The default behavior is to rescale the X axis to display all the values.
// This option takes no effect if all the values on the series fit into the
// LineChart area.
func XAxisUnscaled() Option { _ = "STUB: not implemented"; return *new(Option) }

// ZoomHightlightColor sets the background color of the area that is selected
// with mouse in order to zoom the linechart.
// Defaults to color number 235.
func ZoomHightlightColor(c cell.Color) Option { _ = "STUB: not implemented"; return *new(Option) }

// ZoomStepPercent sets the zooming step on each mouse scroll event as the
// percentage of the size of the X axis.
// The value must be in range 0 < value <= 100.
// Defaults to zoom.DefaultScrollStep.
func ZoomStepPercent(perc int) Option { _ = "STUB: not implemented"; return *new(Option) }

// YAxisFormattedValues sets a value formatter for the Y axis values.
// If a formatter is set, it will format the values with the desired
// ValueFormatter and will use the retuning string from the formatter
// instead of the numeric value to represent this value on the Y axis.
func YAxisFormattedValues(vfmt ValueFormatter) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// ValueFormatter will be used to format values onto string based
// representation.
// The received float64 value could be a math.NaN value.
type ValueFormatter func(value float64) string
