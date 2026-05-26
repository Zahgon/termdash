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

package linechart

// value_formatter.go provides common implementations of ValueFormatter that can be
// used with the YAxisFormattedValues() LineChart option.

import (
	"time"
)

// durationSingleUnitPrettyFormat returns the pretty format in one single
// unit for a time.Duration, the different returned unit formats
// are: nanoseconds, microseconds, milliseconds, seconds, minutes
// hours, days.
func durationSingleUnitPrettyFormat(d time.Duration, decimals int) string {
	_ = "STUB: not implemented"
	// Check if the duration is less than 0.
	return ""
}

// Nanoseconds.

// Microseconds.

// Milliseconds.

// Seconds.

// Minutes.

// Hours.

// Days.

func suffixDecimalFormat(decimals int, suffix string) string { _ = "STUB: not implemented"; return "" }

// Safe `%` character for fmt.

// ValueFormatterSingleUnitDuration is a factory to create a custom duration
// in a single unit representation formatter based on a unit and the decimals
// to truncate.
// If the received decimal value is negative it will fallback to a 0 decimal
// value.
// The result value formatter handles NaN values, if the value formatter
// receives a NaN float64 it will return an empty string.
func ValueFormatterSingleUnitDuration(unit time.Duration, decimals int) ValueFormatter {
	_ = "STUB: not implemented"
	return *new(ValueFormatter)
}

// ValueFormatterSingleUnitSeconds is a formatter that will receive
// seconds unit in the float64 argument and will return a pretty
// format in one single unit without decimals, it doesn't round,
// it truncates.
// Received seconds that are NaN will be ignored and return an
// empty string.
func ValueFormatterSingleUnitSeconds(seconds float64) string { _ = "STUB: not implemented"; return "" }

// ValueFormatterRound is a formatter that will receive a float64
// value and will round to the nearest value without decimals.
func ValueFormatterRound(value float64) string { _ = "STUB: not implemented"; return "" }

// ValueFormatterRoundWithSuffix is a factory that returns a formatter
// that will receive a float64 value and will round to the nearest value
// without decimals adding a suffix to the final value string representation.
func ValueFormatterRoundWithSuffix(suffix string) ValueFormatter {
	_ = "STUB: not implemented"
	return *new(ValueFormatter)
}

// ValueFormatterSuffix is a factory that returns a formatter
// that will receive a float64 value and return a string representation with
// the desired number of decimal truncated and a suffix.
func ValueFormatterSuffix(decimals int, suffix string) ValueFormatter {
	_ = "STUB: not implemented"
	return *new(ValueFormatter)
}

// valueFormatterSuffixWithTransformer is a factory that returns a formatter
// that will apply a transform function to the received value before
// returning the decimal with suffix representation.
func valueFormatterSuffixWithTransformer(decimals int, suffix string, transformFunc func(float64) float64) ValueFormatter {
	_ = "STUB: not implemented"
	return *new(ValueFormatter)
}
