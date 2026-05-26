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

package faketerm

// diff.go provides functions that highlight differences between fake terminals.

import (
	"image"

	"github.com/mum4k/termdash/cell"
)

// optDiff is used to display differences in cell options.
type optDiff struct {
	// point indicates the cell with the differing options.
	point image.Point

	got  *cell.Options
	want *cell.Options
}

// Diff compares the two terminals, returning an empty string if there is not
// difference. If a difference is found, returns a human readable description
// of the differences.
func Diff(want, got *Terminal) string { _ = "STUB: not implemented"; return "" }
