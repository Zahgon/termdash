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

// Package numbers implements various numerical functions.
package numbers

import (
	"image"
)

// RoundToNonZeroPlaces rounds the float up, so that it has at least the provided
// number of non-zero decimal places.
// Returns the rounded float and the number of leading decimal places that
// are zero. Returns the original float when places is zero. Negative places
// are treated as positive, so that -2 == 2.
func RoundToNonZeroPlaces(f float64, places int) (float64, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// multToNonZero returns multiplier for the float, so that the first decimal
// place is non-zero. The float must not be zero.
func multToNonZero(f float64) int { _ = "STUB: not implemented"; return 0 }

// placesToMult translates the number of decimal places to a multiple of 10.
func placesToMult(places int) int { _ = "STUB: not implemented"; return 0 }

// multToPlaces translates the multiple of 10 to a number of decimal places.
func multToPlaces(mult int) int { _ = "STUB: not implemented"; return 0 }

// zeroBeforeDecimal modifies the float so that it only has zero value before
// the decimal point.
func zeroBeforeDecimal(f float64) float64 { _ = "STUB: not implemented"; return 0 }

// MinMax returns the smallest and the largest value among the provided values.
// Returns (0, 0) if there are no values.
// Ignores NaN values. Allowing NaN values could lead to a corner case where all
// values can be NaN, in this case the function will return NaN as min and max.
func MinMax(values []float64) (min, max float64) { _ = "STUB: not implemented"; return 0, 0 }

// MinMaxInts returns the smallest and the largest int value among the provided
// values. Returns (0, 0) if there are no values.
func MinMaxInts(values []int) (min, max int) { _ = "STUB: not implemented"; return 0, 0 }

// DegreesToRadians converts degrees to the equivalent in radians.
func DegreesToRadians(degrees int) float64 { _ = "STUB: not implemented"; return 0 }

// RadiansToDegrees converts radians to the equivalent in degrees.
func RadiansToDegrees(radians float64) int { _ = "STUB: not implemented"; return 0 }

// Abs returns the absolute value of x.
func Abs(x int) int { _ = "STUB: not implemented"; return 0 }

// findGCF finds the greatest common factor of two integers.
func findGCF(a, b int) int { _ = "STUB: not implemented"; return 0 }

// https://en.wikipedia.org/wiki/Euclidean_algorithm

// SimplifyRatio simplifies the given ratio.
func SimplifyRatio(ratio image.Point) image.Point {
	_ = "STUB: not implemented"
	return *new(image.Point)
}

// SplitByRatio splits the provided number by the specified ratio.
func SplitByRatio(n int, ratio image.Point) image.Point {
	_ = "STUB: not implemented"
	return *new(image.Point)
}
