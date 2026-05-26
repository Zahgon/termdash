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

// Package attrrange simplifies tracking of attributes that apply to a range of
// items.
// Refer to the examples in the test file for details on usage.
package attrrange

// AttrRange is a range of items that share the same attributes.
type AttrRange struct {
	// Low is the first position where these attributes apply.
	Low int

	// High is the end of the range. The attributes apply to all items in range
	// Low <= b < high.
	High int

	// AttrIdx is the index of the attributes that apply to this range.
	AttrIdx int
}

// newAttrRange returns a new AttrRange instance.
func newAttrRange(low, high, attrIdx int) *AttrRange { _ = "STUB: not implemented"; return nil }

// Tracker tracks attributes that apply to a range of items.
// This object is not thread safe.
type Tracker struct {
	// ranges maps low indices of ranges to the attribute ranges.
	ranges map[int]*AttrRange
}

// NewTracker returns a new tracker of ranges that share the same attributes.
func NewTracker() *Tracker { _ = "STUB: not implemented"; return nil }

// Add adds a new range of items that share attributes with the specified
// index.
// The low position of the range must not overlap with low position of any
// existing range.
func (t *Tracker) Add(low, high, attrIdx int) error { _ = "STUB: not implemented"; return nil }

// ForPosition returns attribute index that apply to the specified position.
// Returns ErrNotFound when the requested position wasn't found in any of the
// known ranges.
func (t *Tracker) ForPosition(pos int) (*AttrRange, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
