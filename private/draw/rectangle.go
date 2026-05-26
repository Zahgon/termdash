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

// rectangle.go draws a rectangle.

import (
	"image"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/private/canvas"
)

// RectangleOption is used to provide options to the Rectangle function.
type RectangleOption interface {
	// set sets the provided option.
	set(*rectOptions)
}

// rectOptions stores the provided options.
type rectOptions struct {
	cellOpts []cell.Option
	char     rune
}

// rectOption implements RectangleOption.
type rectOption func(rOpts *rectOptions)

// set implements RectangleOption.set.
func (ro rectOption) set(rOpts *rectOptions) {
	_ = "STUB: not implemented"

	// RectCellOpts sets options on the cells that create the rectangle.
	return
}

func RectCellOpts(opts ...cell.Option) RectangleOption {
	_ = "STUB: not implemented"
	return *new(RectangleOption)
}

// DefaultRectChar is the default value for the RectChar option.
const DefaultRectChar = ' '

// RectChar sets the character used in each of the cells of the rectangle.
func RectChar(c rune) RectangleOption { _ = "STUB: not implemented"; return *new(RectangleOption) }

// Rectangle draws a filled rectangle on the canvas.
func Rectangle(c *canvas.Canvas, r image.Rectangle, opts ...RectangleOption) error {
	_ = "STUB: not implemented"
	return nil
}
