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

package draw

// border.go contains code that draws borders.

import (
	"image"

	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/private/canvas"
)

// BorderOption is used to provide options to Border().
type BorderOption interface {
	// set sets the provided option.
	set(*borderOptions)
}

// borderOptions stores the provided options.
type borderOptions struct {
	cellOpts      []cell.Option
	lineStyle     linestyle.LineStyle
	title         string
	titleOM       OverrunMode
	titleCellOpts []cell.Option
	titleHAlign   align.Horizontal
}

// borderOption implements BorderOption.
type borderOption func(bOpts *borderOptions)

// set implements BorderOption.set.
func (bo borderOption) set(bOpts *borderOptions) {
	_ = "STUB: not implemented"

	// DefaultBorderLineStyle is the default value for the BorderLineStyle option.
	return
}

const DefaultBorderLineStyle = linestyle.Light

// BorderLineStyle sets the style of the line used to draw the border.
func BorderLineStyle(ls linestyle.LineStyle) BorderOption {
	_ = "STUB: not implemented"
	return *new(BorderOption)
}

// BorderCellOpts sets options on the cells that create the border.
func BorderCellOpts(opts ...cell.Option) BorderOption {
	_ = "STUB: not implemented"
	return *new(BorderOption)
}

// BorderTitle sets a title for the border.
func BorderTitle(title string, overrun OverrunMode, opts ...cell.Option) BorderOption {
	_ = "STUB: not implemented"
	return *new(BorderOption)
}

// BorderTitleAlign configures the horizontal alignment for the title.
func BorderTitleAlign(h align.Horizontal) BorderOption {
	_ = "STUB: not implemented"
	return *new(BorderOption)
}

// borderChar returns the correct border character from the parts for the use
// at the specified point of the border. Returns -1 if no character should be at
// this point.
func borderChar(p image.Point, border image.Rectangle, parts map[linePart]rune) rune {
	_ = "STUB: not implemented"
	return 0
}

// drawTitle draws a text title at the top of the border.
func drawTitle(c *canvas.Canvas, border image.Rectangle, opt *borderOptions) error {
	_ = "STUB: not implemented"
	// Don't attempt to draw the title if there isn't space for at least one rune.
	// The title must not overwrite any of the corner runes on the border so we
	// need the following minimum width.
	return nil
}

// One space for the top left corner char.

// One space for the top right corner char.

// Border draws a border on the canvas.
func Border(c *canvas.Canvas, border image.Rectangle, opts ...BorderOption) error {
	_ = "STUB: not implemented"
	return nil
}
