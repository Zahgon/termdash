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

package text

// scroll.go contains code that tracks the current scrolling position.

// scrollTracker tracks the current scrolling position for the Text widget.
//
// The text widget displays the contained text buffer as lines of text that fit
// the widget's canvas. The main goal of this object is to inform the text
// widget which should be the first drawn line from the buffer. This depends on
// two things, the scrolling position based on user inputs and whether the text
// widget is configured to roll the content up as new text is added by the
// client.
//
// The rolling Vs. scrolling state is tracked in an FSM implemented in this
// file.
//
// The client can scroll the content by either a keyboard or a mouse event. The
// widget receives these events concurrently with requests to redraw the
// content, so this objects keeps a track of all the scrolling events that
// happened since the last redraw and consumes them when calculating which is
// the first drawn line on the next redraw event.
//
// This is not thread safe.
type scrollTracker struct {
	// scroll stores user requests to scroll up (negative) or down (positive).
	// E.g. -1 means up by one line and 2 means down by two lines.
	scroll int

	// scrollPage stores user requests to scroll up (negative) or down
	// (positive) by a page of content. E.g. -1 means up by one page and 2
	// means down by two pages.
	scrollPage int

	// first tracks the first line that will be printed.
	first int

	// state is the state of the scrolling FSM.
	state rollState
}

// newScrollTracker returns a new scroll tracker.
func newScrollTracker(opts *options) *scrollTracker { _ = "STUB: not implemented"; return nil }

// upOneLine processes a user request to scroll up by one line.
func (st *scrollTracker) upOneLine() {
	_ = "STUB: not implemented"

	// downOneLine processes a user request to scroll down by one line.
	return
}

func (st *scrollTracker) downOneLine() {
	_ = "STUB: not implemented"

	// upOnePage processes a user request to scroll up by one page.
	return
}

func (st *scrollTracker) upOnePage() {
	_ = "STUB: not implemented"

	// downOnePage processes a user request to scroll down by one page.
	return
}

func (st *scrollTracker) downOnePage() {
	_ = "STUB: not implemented"

	// doScroll processes any outstanding scroll requests and calculates the
	// resulting first line.
	return
}

func (st *scrollTracker) doScroll(lines, height int) int { _ = "STUB: not implemented"; return 0 }

// firstLine returns the number of the first line that should be drawn on a
// canvas of the specified height if there is the provided number of lines of
// text.
func (st *scrollTracker) firstLine(lines, height int) int {
	_ = "STUB: not implemented"
	// Execute the scrolling FSM.
	return 0
}

// rollState is a state in the scrolling FSM.
type rollState func(st *scrollTracker, lines, height int) rollState

// rollingDisabled is a state where content rolling was disabled by the
// configuration of the Text widget.
func rollingDisabled(st *scrollTracker, lines, height int) rollState {
	_ = "STUB: not implemented"
	return *new(rollState)
}

// rollToEnd is a state in which the last line of the content is always
// visible. When new content arrives, it is rolled upwards.
func rollToEnd(st *scrollTracker, lines, height int) rollState {
	_ = "STUB: not implemented"
	// If the user didn't scroll, just roll the content so that the last line
	// is visible.
	return *new(rollState)
}

// rollingPaused is a state in which the user scrolled up and made the last
// line scroll out of the view, so the content rolling is paused.
func rollingPaused(st *scrollTracker, lines, height int) rollState {
	_ = "STUB: not implemented"
	return *new(rollState)
}

// lastLineVisible returns true if the last text line from within the buffer of
// the text widget is visible on the canvas when drawing of the text starts
// from the specified start line, there is the provided total amount of lines
// and the canvas has the height.
func lastLineVisible(start, lines, height int) bool { _ = "STUB: not implemented"; return false }

// normalizeScroll returns normalized position of the first line that should be
// drawn when drawing the specified number of lines on a canvas with the
// provided height.
func normalizeScroll(first, lines, height int) int { _ = "STUB: not implemented"; return 0 }

// Scrolling not necessary if the content fits.
