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

// Package segmentdisplay is a widget that displays text by simulating a
// segment display.
package segmentdisplay

import (
	"image"
	"strings"
	"sync"

	"github.com/mum4k/termdash/private/attrrange"
	"github.com/mum4k/termdash/private/canvas"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgetapi"
)

// SegmentDisplay displays ASCII content by simulating a segment display.
//
// Automatically determines the size of individual segments with goal of
// maximizing the segment size or with fitting the entire text depending on the
// provided options.
//
// Segment displays support only a subset of ASCII characters, provided options
// determine the behavior when an unsupported character is encountered.
//
// Implements widgetapi.Widget. This object is thread-safe.
type SegmentDisplay struct {
	// buff contains the text to be displayed.
	buff strings.Builder

	// givenWOpts are write options given for the text in buff.
	givenWOpts []*writeOptions
	// wOptsTracker tracks the positions in a buff to which the givenWOpts apply.
	wOptsTracker *attrrange.Tracker

	// lastCanFit is the number of segments that could fit the area the last
	// time Draw was called.
	lastCanFit int

	// dotChars are characters that are drawn using the dot segment.
	// All other characters are draws using the 16-segment display.
	dotChars map[rune]bool

	// mu protects the widget.
	mu sync.Mutex

	// opts are the provided options.
	opts *options
}

// New returns a new SegmentDisplay.
func New(opts ...Option) (*SegmentDisplay, error) { _ = "STUB: not implemented"; return nil, nil }

// TextChunk is a part of or the full text that will be displayed.
type TextChunk struct {
	text  string
	wOpts *writeOptions
}

// NewChunk creates a new text chunk.
func NewChunk(text string, wOpts ...WriteOption) *TextChunk { _ = "STUB: not implemented"; return nil }

// Write writes text for the widget to display. Subsequent calls replace text
// written previously. All the provided text chunks are broken into characters
// and each character is displayed in one segment.
//
// The provided write options determine the behavior when text contains
// unsupported characters and set cell options for cells that contain
// individual display segments.
//
// Each of the text chunks can have its own options. At least one chunk must be
// specified.
//
// Any provided options override options given to New.
func (sd *SegmentDisplay) Write(chunks []*TextChunk, opts ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

// Capacity returns the number of characters that can fit into the canvas.
// This is essentially the number of individual segments that can fit on the
// canvas at the time the last call to draw. Returns zero if draw wasn't
// called.
//
// Note that this capacity changes each time the terminal resizes, so there is
// no guarantee this remains the same next time Draw is called.
// Should be used as a hint only.
func (sd *SegmentDisplay) Capacity() int { _ = "STUB: not implemented"; return 0 }

// Reset resets the widget back to empty content.
func (sd *SegmentDisplay) Reset() { _ = "STUB: not implemented"; return }

// reset is the implementation of Reset.
// Caller must hold sd.mu.
func (sd *SegmentDisplay) reset() { _ = "STUB: not implemented"; return }

// preprocess determines the size of individual segments maximizing their
// height or the amount of displayed characters based on the specified options.
// Returns the area required for a single segment, the text that we can fit and
// size of gaps between segments in cells.
func (sd *SegmentDisplay) preprocess(cvsAr image.Rectangle) (*segArea, error) {
	_ = "STUB: not implemented"
	return nil,
		// We're guaranteed by Write to only have ASCII characters.
		nil
}

// Draw draws the SegmentDisplay widget onto the canvas.
// Implements widgetapi.Widget.Draw.
func (sd *SegmentDisplay) Draw(cvs *canvas.Canvas, meta *widgetapi.Meta) error {
	_ = "STUB: not implemented"
	return nil
}

// Text options for the current byte.

// Get the next write options.

// drawChar draws a single character onto the provided canvas.
func (sd *SegmentDisplay) drawChar(dCvs *canvas.Canvas, c rune, wOpts *writeOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// Keyboard input isn't supported on the SegmentDisplay widget.
func (*SegmentDisplay) Keyboard(k *terminalapi.Keyboard, meta *widgetapi.EventMeta) error {
	_ = "STUB: not implemented"
	return nil
}

// Mouse input isn't supported on the SegmentDisplay widget.
func (*SegmentDisplay) Mouse(m *terminalapi.Mouse, meta *widgetapi.EventMeta) error {
	_ = "STUB: not implemented"
	return nil
}

// Options implements widgetapi.Widget.Options.
func (sd *SegmentDisplay) Options() widgetapi.Options {
	_ = "STUB: not implemented"
	return *

	// The smallest supported size of a display segment.
	new(widgetapi.Options)
}
