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

package termbox

// cell_options.go converts termdash cell options to the termbox format.

import (
	"github.com/mum4k/termdash/cell"
	tbx "github.com/nsf/termbox-go"
)

// cellColor converts termdash cell color to the termbox format.
func cellColor(c cell.Color) tbx.Attribute {
	_ = "STUB: not implemented"
	// Special cases for backward compatibility after we have aligned the
	// definition of the first 16 colors with Xterm and tcell.
	// This ensures that users that run with termbox-go don't experience any
	// change in colors.
	return *new(tbx.Attribute)
}

// cellOptsToFg converts the cell options to the termbox foreground attribute.
func cellOptsToFg(opts *cell.Options) (tbx.Attribute, error) {
	_ = "STUB: not implemented"
	return *new(tbx.Attribute), nil
}

// Termbox doesn't have an italics attribute

// Termbox doesn't have a strikethrough attribute

// Termbox doesn't have a blink attribute

// cellOptsToBg converts the cell options to the termbox background attribute.
func cellOptsToBg(opts *cell.Options) tbx.Attribute {
	_ = "STUB: not implemented"
	return *new(tbx.Attribute)
}
