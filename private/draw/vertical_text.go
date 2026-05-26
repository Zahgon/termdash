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

package draw

// vertical_text.go contains code that prints UTF-8 encoded strings on the
// canvas in vertical columns instead of lines.

import (
	"image"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/private/canvas"
)

// VerticalTextOption is used to provide options to Text().
type VerticalTextOption interface {
	// set sets the provided option.
	set(*verticalTextOptions)
}

// verticalTextOptions stores the provided options.
type verticalTextOptions struct {
	cellOpts    []cell.Option
	maxY        int
	overrunMode OverrunMode
}

// verticalTextOption implements VerticalTextOption.
type verticalTextOption func(*verticalTextOptions)

// set implements VerticalTextOption.set.
func (vto verticalTextOption) set(vtOpts *verticalTextOptions) {
	_ = "STUB: not implemented"

	// VerticalTextCellOpts sets options on the cells that contain the text.
	return
}

func VerticalTextCellOpts(opts ...cell.Option) VerticalTextOption {
	_ = "STUB: not implemented"
	return *new(VerticalTextOption)
}

// VerticalTextMaxY sets a limit on the Y coordinate (row) of the drawn text.
// The Y coordinate of all cells used by the vertical text must be within
// start.Y <= Y < VerticalTextMaxY.
// If not provided, the height of the canvas is used as VerticalTextMaxY.
func VerticalTextMaxY(y int) VerticalTextOption {
	_ = "STUB: not implemented"
	return *new(VerticalTextOption)
}

// VerticalTextOverrunMode indicates what to do with text that overruns the
// VerticalTextMaxY() or the width of the canvas if VerticalTextMaxY() isn't
// specified.
// Defaults to OverrunModeStrict.
func VerticalTextOverrunMode(om OverrunMode) VerticalTextOption {
	_ = "STUB: not implemented"
	return *new(VerticalTextOption)
}

// VerticalText prints the provided text on the canvas starting at the provided point.
// The text is printed in a vertical orientation, i.e:
//
//	H
//	e
//	l
//	l
//	o
func VerticalText(c *canvas.Canvas, text string, start image.Point, opts ...VerticalTextOption) error {
	_ = "STUB: not implemented"
	return nil
}
