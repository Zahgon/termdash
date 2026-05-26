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

package axes

// value.go contains code dealing with values on the line chart.

// ValueOption is used to provide options to the NewValue function.
type ValueOption interface {
	// set sets the provided option.
	set(*valueOptions)
}

type valueOptions struct {
	formatter func(v float64) string
}

// valueOption implements ValueOption.
type valueOption func(opts *valueOptions)

// set implements ValueOption.set.
func (vo valueOption) set(opts *valueOptions) {
	_ = "STUB: not implemented"

	// ValueFormatter sets a custom formatter for the value.
	return
}

func ValueFormatter(formatter func(float64) string) ValueOption {
	_ = "STUB: not implemented"
	return *new(ValueOption)
}

// Value represents one value.
type Value struct {
	// Value is the original unmodified value.
	Value float64
	// Rounded is the value rounded up to the nonZeroPlaces number of non-zero
	// decimal places.
	Rounded float64
	// ZeroDecimals indicates how many decimal places in Rounded have a value
	// of zero.
	ZeroDecimals int
	// NonZeroDecimals indicates the rounding precision used, it is provided on
	// a call to newValue.
	NonZeroDecimals int

	// formatter will format value to a string representation of the value,
	// if Formatter is not present it will fallback to default format.
	formatter func(float64) string
	// text value if this value was constructed using NewTextValue.
	text string
}

// String implements fmt.Stringer.
func (v *Value) String() string { _ = "STUB: not implemented"; return "" }

// NewValue returns a new instance representing the provided value, rounding
// the value up to the specified number of non-zero decimal places.
func NewValue(v float64, nonZeroDecimals int, opts ...ValueOption) *Value {
	_ = "STUB: not implemented"
	return nil
}

// NewTextValue constructs a value out of the provided text.
func NewTextValue(text string) *Value { _ = "STUB: not implemented"; return nil }

// Text returns textual representation of the value.
func (v *Value) Text() string { _ = "STUB: not implemented"; return "" }

func defaultFormatter(value float64, nonZeroDecimals, zeroDecimals int) string {
	_ = "STUB: not implemented"
	return ""
}
