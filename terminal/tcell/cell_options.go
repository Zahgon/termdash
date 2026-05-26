// Copyright 2020 Google Inc.
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

package tcell

import (
	tcell "github.com/gdamore/tcell/v2"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/terminal/terminalapi"
)

// cellColor converts termdash cell color to the tcell format.
func cellColor(c cell.Color) tcell.Color { _ = "STUB: not implemented"; return *new(tcell.Color) }

// Subtract one, because cell.ColorBlack has value one instead of zero.
// Zero is used for cell.ColorDefault instead.

// colorToMode adjusts the color to the color mode.
func colorToMode(c cell.Color, colorMode terminalapi.ColorMode) cell.Color {
	_ = "STUB: not implemented"
	return *new(cell.Color)
}

// Add one for cell.ColorDefault.

// Add one for cell.ColorDefault.

// Add one for cell.ColorDefault.

// Add one for cell.ColorDefault.

// cellOptsToStyle converts termdash cell color to the tcell format.
func cellOptsToStyle(opts *cell.Options, colorMode terminalapi.ColorMode) tcell.Style {
	_ = "STUB: not implemented"
	return *new(tcell.Style)
}
