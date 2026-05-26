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
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/terminal/terminalapi"
)

// tcell representation of the space key
var tcellSpaceKey = tcell.Key(' ')

// tcellToTd maps tcell key values to the termdash format.
var tcellToTd = map[tcell.Key]keyboard.Key{
	tcellSpaceKey:           keyboard.KeySpace,
	tcell.KeyF1:             keyboard.KeyF1,
	tcell.KeyF2:             keyboard.KeyF2,
	tcell.KeyF3:             keyboard.KeyF3,
	tcell.KeyF4:             keyboard.KeyF4,
	tcell.KeyF5:             keyboard.KeyF5,
	tcell.KeyF6:             keyboard.KeyF6,
	tcell.KeyF7:             keyboard.KeyF7,
	tcell.KeyF8:             keyboard.KeyF8,
	tcell.KeyF9:             keyboard.KeyF9,
	tcell.KeyF10:            keyboard.KeyF10,
	tcell.KeyF11:            keyboard.KeyF11,
	tcell.KeyF12:            keyboard.KeyF12,
	tcell.KeyInsert:         keyboard.KeyInsert,
	tcell.KeyDelete:         keyboard.KeyDelete,
	tcell.KeyHome:           keyboard.KeyHome,
	tcell.KeyEnd:            keyboard.KeyEnd,
	tcell.KeyPgUp:           keyboard.KeyPgUp,
	tcell.KeyPgDn:           keyboard.KeyPgDn,
	tcell.KeyUp:             keyboard.KeyArrowUp,
	tcell.KeyDown:           keyboard.KeyArrowDown,
	tcell.KeyLeft:           keyboard.KeyArrowLeft,
	tcell.KeyRight:          keyboard.KeyArrowRight,
	tcell.KeyEnter:          keyboard.KeyEnter,
	tcell.KeyCtrlA:          keyboard.KeyCtrlA,
	tcell.KeyCtrlB:          keyboard.KeyCtrlB,
	tcell.KeyCtrlC:          keyboard.KeyCtrlC,
	tcell.KeyCtrlD:          keyboard.KeyCtrlD,
	tcell.KeyCtrlE:          keyboard.KeyCtrlE,
	tcell.KeyCtrlF:          keyboard.KeyCtrlF,
	tcell.KeyCtrlG:          keyboard.KeyCtrlG,
	tcell.KeyCtrlJ:          keyboard.KeyCtrlJ,
	tcell.KeyCtrlK:          keyboard.KeyCtrlK,
	tcell.KeyCtrlL:          keyboard.KeyCtrlL,
	tcell.KeyCtrlN:          keyboard.KeyCtrlN,
	tcell.KeyCtrlO:          keyboard.KeyCtrlO,
	tcell.KeyCtrlP:          keyboard.KeyCtrlP,
	tcell.KeyCtrlQ:          keyboard.KeyCtrlQ,
	tcell.KeyCtrlR:          keyboard.KeyCtrlR,
	tcell.KeyCtrlS:          keyboard.KeyCtrlS,
	tcell.KeyCtrlT:          keyboard.KeyCtrlT,
	tcell.KeyCtrlU:          keyboard.KeyCtrlU,
	tcell.KeyCtrlV:          keyboard.KeyCtrlV,
	tcell.KeyCtrlW:          keyboard.KeyCtrlW,
	tcell.KeyCtrlX:          keyboard.KeyCtrlX,
	tcell.KeyCtrlY:          keyboard.KeyCtrlY,
	tcell.KeyCtrlZ:          keyboard.KeyCtrlZ,
	tcell.KeyBackspace:      keyboard.KeyBackspace,
	tcell.KeyTab:            keyboard.KeyTab,
	tcell.KeyBacktab:        keyboard.KeyBacktab,
	tcell.KeyEscape:         keyboard.KeyEsc,
	tcell.KeyCtrlBackslash:  keyboard.KeyCtrlBackslash,
	tcell.KeyCtrlRightSq:    keyboard.KeyCtrlRsqBracket,
	tcell.KeyCtrlUnderscore: keyboard.KeyCtrlUnderscore,
	tcell.KeyBackspace2:     keyboard.KeyBackspace2,
	tcell.KeyCtrlSpace:      keyboard.KeyCtrlSpace,
}

// convKey converts a tcell keyboard event to the termdash format.
func convKey(event *tcell.EventKey) terminalapi.Event {
	_ = "STUB: not implemented"
	return *new(terminalapi.Event)
}

// convMouse converts a tcell mouse event to the termdash format.
// Since tcell supports many combinations of mouse events, such as multiple mouse buttons pressed at the same time,
// this function returns nil if the event is unsupported by termdash.
func convMouse(event *tcell.EventMouse) terminalapi.Event {
	_ = "STUB: not implemented"
	return *new(terminalapi.Event)
}

// tcell uses signed int16 for button masks, and negative values are invalid

// Get wheel events

// Return wheel event if found

// Unknown event to termdash

// convResize converts a tcell resize event to the termdash format.
func convResize(event *tcell.EventResize) terminalapi.Event {
	_ = "STUB: not implemented"
	return *new(terminalapi.Event)
}

// toTermdashEvents converts a tcell event to the termdash event format.
// This function returns nil if the event is unsupported by termdash.
func toTermdashEvents(event tcell.Event) []terminalapi.Event { _ = "STUB: not implemented"; return nil }
