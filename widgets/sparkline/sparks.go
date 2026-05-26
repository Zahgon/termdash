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

package sparkline

// sparks.go contains code that determines which characters should be used to
// represent a value on the SparkLine.

import (
	"fmt"

	"github.com/mum4k/termdash/private/runewidth"
)

// sparks are the characters used to draw the SparkLine.
var sparks = []rune{'▁', '▂', '▃', '▄', '▅', '▆', '▇', '█'}

// visibleMax determines the maximum visible data point given the canvas width.
// Returns a slice that contains only visible data points and the maximum value
// among them.
func visibleMax(data []int, width int) ([]int, int) { _ = "STUB: not implemented"; return nil, 0 }

// blocks represents the building blocks that display one value on a SparkLine.
// I.e. one vertical bar.
type blocks struct {
	// full is the number of fully populated blocks.
	full int

	// partSpark is the spark character from sparks that should be used in the
	// topmost block. Equals to zero if no partial block should be displayed.
	partSpark rune
}

// toBlocks determines the number of full and partial vertical blocks required
// to represent the provided value given the specified max visible value and
// number of vertical cells available to the SparkLine.
func toBlocks(value, max, vertCells int) blocks { _ = "STUB: not implemented"; return *new(blocks) }

// How many of the smallest spark elements fit into a cell.

// Scale is how much of the max does one smallest spark element represent,
// given the vertical cells that will be used to represent the value.

// How many smallest spark elements are needed to represent the value.

// init ensures that all spark characters are half-width runes.
// The SparkLine widget assumes that each value can be represented in a column
// that has a width of one cell.
func init() {
	for i, s := range sparks {
		if got := runewidth.RuneWidth(s); got > 1 {
			panic(fmt.Sprintf("all sparks must be half-width runes (width of one), spark[%d] has width %d", i, got))
		}
	}
}
