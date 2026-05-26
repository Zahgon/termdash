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

// Package trig implements various trigonometrical calculations.
package trig

import (
	"image"
)

// CirclePointAtAngle given an angle in degrees and a circle midpoint and
// radius, calculates coordinates of a point on the circle at that angle.
// Angles are zero at the X axis and grow counter-clockwise.
func CirclePointAtAngle(degrees int, mid image.Point, radius int) image.Point {
	_ = "STUB: not implemented"
	return *new(image.Point)
}

// Y coordinates grow down on the canvas.

// CircleAngleAtPoint given a point on a circle and its midpoint,
// calculates the angle in degrees.
// Angles are zero at the X axis and grow counter-clockwise.
func CircleAngleAtPoint(point, mid image.Point) int { _ = "STUB: not implemented"; return 0 }

// PointIsIn asserts whether the provided point is inside of a shape outlined
// with the provided points.
// Does not verify that the shape is closed or complete, it merely counts the
// number of intersections with the shape on one row.
func PointIsIn(p image.Point, points []image.Point) bool { _ = "STUB: not implemented"; return false }

// Not inside if it is on the shape.

// maps y->x

const (
	// MinAngle is the smallest valid angle in degrees.
	MinAngle = 0
	// MaxAngle is the largest valid angle in degrees.
	MaxAngle = 360
)

// angleRange represents a range of angles in degrees.
// The range includes all angles such that start <= angle <= end.
type angleRange struct {
	// start is the start if the range.
	// This is always less or equal to the end.
	start int

	// end is the end of the range.
	end int
}

// contains asserts whether the specified angle is in the range.
func (ar *angleRange) contains(angle int) bool { _ = "STUB: not implemented"; return false }

// normalizeRange normalizes the start and end angles in degrees into ranges of
// angles. Useful for cases where the 0/360 point falls within the range.
// E.g:
//
//	0,25   => angleRange{0, 26}
//	0,360  => angleRange{0, 361}
//	359,20 => angleRange{359, 361}, angleRange{0, 21}
func normalizeRange(start, end int) ([]*angleRange, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The range is crossing the 0/360 degree point.
// Break it into multiple ranges.

// RangeSize returns the size of the degree range.
// E.g:
//
//	0,25  => 25
//	359,1 => 2
func RangeSize(start, end int) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// RangeMid returns an angle that lies in the middle between start and end.
// E.g:
//
//	0,10   => 5
//	350,10 => 0
func RangeMid(start, end int) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// FilterByAngle filters the provided points, returning only those that fall
// within the starting and the ending angle on a circle with the provided mid
// point.
func FilterByAngle(points []image.Point, mid image.Point, start, end int) ([]image.Point, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Edge case, this might mean 0 or 360.
// Decide based on where we are starting.
