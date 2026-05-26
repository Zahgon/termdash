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

// Package wrap implements line wrapping at character or word boundaries.
package wrap

import (
	"github.com/mum4k/termdash/private/canvas/buffer"
)

// Mode sets the wrapping mode.
type Mode int

// String implements fmt.Stringer()
func (m Mode) String() string { _ = "STUB: not implemented"; return "" }

// modeNames maps Mode values to human readable names.
var modeNames = map[Mode]string{
	Never:   "WrapModeNever",
	AtRunes: "WrapModeAtRunes",
	AtWords: "WrapModeAtWords",
}

const (
	// Never is the default wrapping mode, which disables line wrapping.
	Never Mode = iota

	// AtRunes is a wrapping mode where if the width of the text crosses the
	// width of the canvas, wrapping is performed at rune boundaries.
	AtRunes

	// AtWords is a wrapping mode where if the width of the text crosses the
	// width of the canvas, wrapping is performed at word boundaries. The
	// wrapping still switches back to the AtRunes mode for any words that are
	// longer than the width.
	AtWords
)

// ValidText validates the provided text for wrapping.
// The text must not be empty, contain any control or
// space characters other than '\n' and ' '.
func ValidText(text string) error { _ = "STUB: not implemented"; return nil }

// Allowed space and control runes.

// ValidCells validates the provided cells for wrapping.
// The text in the cells must follow the same rules as described for ValidText.
func ValidCells(cells []*buffer.Cell) error { _ = "STUB: not implemented"; return nil }

// Cells returns the cells wrapped into individual lines according to the
// specified width and wrapping mode.
//
// This function consumes any cells that contain newline characters and uses
// them to start new lines.
//
// If the mode is AtWords, this function also drops cells with leading space
// character before a word at which the wrap occurs.
func Cells(cells []*buffer.Cell, width int, m Mode) ([][]*buffer.Cell, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// cellScannerState is a state in the FSM that scans the input text and identifies
// newlines.
type cellScannerState func(*cellScanner) cellScannerState

// cellScanner tracks the progress of scanning the input cells when finding
// lines.
type cellScanner struct {
	// cells are the cells being scanned.
	cells []*buffer.Cell

	// nextIdx is the index of the cell that will be returned by next.
	nextIdx int

	// wordStartIdx stores the starting index of the current word.
	// A starting position of a word includes any leading space characters.
	// E.g.: hello   world
	//            ^
	//            lastWordIdx
	wordStartIdx int
	// wordEndIdx stores the ending index of the current word.
	// The word consists of all indexes that are
	// wordStartIdx <= idx < wordEndIdx.
	// A word also includes any punctuation after it.
	wordEndIdx int

	// width is the width of the canvas the text will be drawn on.
	width int

	// posX tracks the horizontal position of the current cell on the canvas.
	posX int

	// mode is the wrapping mode.
	mode Mode

	// atRunesInWord overrides the mode back to AtRunes.
	atRunesInWord bool

	// lines are the identified lines.
	lines [][]*buffer.Cell

	// line is the current line.
	line []*buffer.Cell
}

// newCellScanner returns a scanner of the provided cells.
func newCellScanner(cells []*buffer.Cell, width int, m Mode) *cellScanner {
	_ = "STUB: not implemented"
	return nil
}

// next returns the next cell and advances the scanner.
// Returns nil when there are no more cells to scan.
func (cs *cellScanner) next() *buffer.Cell { _ = "STUB: not implemented"; return nil }

// peek returns the next cell without advancing the scanner's position.
// Returns nil when there are no more cells to peek at.
func (cs *cellScanner) peek() *buffer.Cell { _ = "STUB: not implemented"; return nil }

// peekPrev returns the previous cell without changing the scanner's position.
// Returns nil if the scanner is at the first cell.
func (cs *cellScanner) peekPrev() *buffer.Cell { _ = "STUB: not implemented"; return nil }

// wordCells returns all the cells that belong to the current word.
func (cs *cellScanner) wordCells() []*buffer.Cell { _ = "STUB: not implemented"; return nil }

// wordWidth returns the width of the current word in cells when printed on the
// terminal.
func (cs *cellScanner) wordWidth() int { _ = "STUB: not implemented"; return 0 }

// isWordStart determines if the scanner is at the beginning of a word.
func (cs *cellScanner) isWordStart() bool { _ = "STUB: not implemented"; return false }

// scanCellRunes scans the cells a rune at a time.
func scanCellRunes(cs *cellScanner) cellScannerState {
	_ = "STUB: not implemented"
	return *new(cellScannerState)
}

// runeToCurrentLine scans a single cell rune onto the current line.
func runeToCurrentLine(cs *cellScanner) cellScannerState {
	_ = "STUB: not implemented"
	return *

	// Move horizontally within the line for each scanned cell.
	new(cellScannerState)
}

// Copy the cell into the current line.

// newLineForLineBreak processes a newline character cell.
func newLineForLineBreak(cs *cellScanner) cellScannerState {
	_ = "STUB: not implemented"
	return *new(cellScannerState)
}

// newLineForAtRunes processes a line wrap at rune boundaries due to canvas width.
func newLineForAtRunes(cs *cellScanner) cellScannerState {
	_ = "STUB: not implemented"
	// The character on which we wrapped will be printed and is the start of
	// new line.
	return *new(cellScannerState)
}

// scanEOF terminates the scanning.
func scanEOF(cs *cellScanner) cellScannerState {
	_ = "STUB: not implemented"
	// Need to add the current line if it isn't empty, or if the previous rune
	// was a newline.
	// Newlines aren't copied onto the lines so just checking for emptiness
	// isn't enough. We still want to include trailing empty newlines if
	// they are in the input text.
	return *new(cellScannerState)
}

// markWordStart stores the starting position of the current word.
func markWordStart(cs *cellScanner) cellScannerState {
	_ = "STUB: not implemented"
	return *new(cellScannerState)
}

// scanWord scans the entire word until it finds its end.
func scanWord(cs *cellScanner) cellScannerState {
	_ = "STUB: not implemented"
	return *new(cellScannerState)
}

// wordToCurrentLine decides how to place the word into the output.
func wordToCurrentLine(cs *cellScanner) cellScannerState {
	_ = "STUB: not implemented"
	return *new(cellScannerState)
}

// Place the word onto the current line.

// wrapWord wraps the word onto the next line or lines.
func wrapWord(cs *cellScanner) cellScannerState {
	_ = "STUB: not implemented"
	// Edge-case - the word starts the line and immediately doesn't fit.
	return *new(cellScannerState)
}

// Skip the leading space when word wrapping.

// Replace the last placed rune with a dash indicating we wrapped the
// word. Only do this for half-width runes.

// Reset the scanner's position back to start scanning at the first
// rune of this word that wasn't placed.

// Edge-case width is one, no space to put the dash rune.

// isWordCell determines if the cell contains a rune that belongs to a word.
func isWordCell(c *buffer.Cell) bool { _ = "STUB: not implemented"; return false }

// runeWrapNeeded returns true if wrapping is needed for the rune at the horizontal
// position on the canvas that has the specified width.
func runeWrapNeeded(r rune, posX, width int) bool { _ = "STUB: not implemented"; return false }
