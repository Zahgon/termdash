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

// Binary termdashdemo demonstrates the functionality of termdash and its various widgets.
// Exits when 'q' is pressed.
package main

import (
	"context"
	"flag"
	"log"
	"sync"
	"time"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/termbox"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/barchart"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/donut"
	"github.com/mum4k/termdash/widgets/gauge"
	"github.com/mum4k/termdash/widgets/linechart"
	"github.com/mum4k/termdash/widgets/segmentdisplay"
	"github.com/mum4k/termdash/widgets/sparkline"
	"github.com/mum4k/termdash/widgets/text"
	"github.com/mum4k/termdash/widgets/textinput"
)

// redrawInterval is how often termdash redraws the screen.
const redrawInterval = 250 * time.Millisecond

// widgets holds the widgets used by this demo.
type widgets struct {
	segDist  *segmentdisplay.SegmentDisplay
	input    *textinput.TextInput
	rollT    *text.Text
	spGreen  *sparkline.SparkLine
	spRed    *sparkline.SparkLine
	gauge    *gauge.Gauge
	heartLC  *linechart.LineChart
	barChart *barchart.BarChart
	donut    *donut.Donut
	leftB    *button.Button
	rightB   *button.Button
	sineLC   *linechart.LineChart

	buttons *layoutButtons
}

// newWidgets creates all widgets used by this demo.
func newWidgets(ctx context.Context, t terminalapi.Terminal, c *container.Container) (*widgets, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// layoutType represents the possible layouts the buttons switch between.
type layoutType int

const (
	// layoutAll displays all the widgets.
	layoutAll layoutType = iota
	// layoutText focuses onto the text widget.
	layoutText
	// layoutSparkLines focuses onto the sparklines.
	layoutSparkLines
	// layoutLineChart focuses onto the linechart.
	layoutLineChart
)

// gridLayout prepares container options that represent the desired screen layout.
// This function demonstrates the use of the grid builder.
// gridLayout() and contLayout() demonstrate the two available layout APIs and
// both produce equivalent layouts for layoutType layoutAll.
func gridLayout(w *widgets, lt layoutType) ([]container.Option, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// contLayout prepares container options that represent the desired screen layout.
// This function demonstrates the direct use of the container API.
// gridLayout() and contLayout() demonstrate the two available layout APIs and
// both produce equivalent layouts for layoutType layoutAll.
// contLayout only produces layoutAll.
func contLayout(w *widgets) ([]container.Option, error) { _ = "STUB: not implemented"; return nil, nil }

// rootID is the ID assigned to the root container.
const rootID = "root"

// Terminal implementations
const (
	termboxTerminal = "termbox"
	tcellTerminal   = "tcell"
)

func main() {
	terminalPtr := flag.String("terminal",
		"tcell",
		"The terminal implementation to use. Available implementations are 'termbox' and 'tcell' (default = tcell).")
	flag.Parse()

	var t terminalapi.Terminal
	var err error
	switch terminal := *terminalPtr; terminal {
	case termboxTerminal:
		t, err = termbox.New(termbox.ColorMode(terminalapi.ColorMode256))
	case tcellTerminal:
		t, err = tcell.New(tcell.ColorMode(terminalapi.ColorMode256))
	default:
		log.Fatalf("Unknown terminal implementation '%s' specified. Please choose between 'termbox' and 'tcell'.", terminal)
		return
	}

	if err != nil {
		panic(err)
	}
	defer t.Close()

	c, err := container.New(t, container.ID(rootID))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	w, err := newWidgets(ctx, t, c)
	if err != nil {
		panic(err)
	}
	lb, err := newLayoutButtons(c, w)
	if err != nil {
		panic(err)
	}
	w.buttons = lb

	gridOpts, err := gridLayout(w, layoutAll) // equivalent to contLayout(w)
	if err != nil {
		panic(err)
	}

	if err := c.Update(rootID, gridOpts...); err != nil {
		panic(err)
	}

	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == keyboard.KeyEsc || k.Key == keyboard.KeyCtrlC {
			cancel()
		}
	}
	if err := termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(quitter), termdash.RedrawInterval(redrawInterval)); err != nil {
		panic(err)
	}
}

// periodic executes the provided closure periodically every interval.
// Exits when the context expires.
func periodic(ctx context.Context, interval time.Duration, fn func() error) {
	_ = "STUB: not implemented"
	return
}

// textState creates a rotated state for the text we are displaying.
func textState(text string, capacity, step int) []rune { _ = "STUB: not implemented"; return nil }

// newTextInput creates a new TextInput field that changes the text on the
// SegmentDisplay.
func newTextInput(updateText chan<- string) (*textinput.TextInput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// newSegmentDisplay creates a new SegmentDisplay that initially shows the
// Termdash name. Shows any text that is sent over the channel.
func newSegmentDisplay(ctx context.Context, t terminalapi.Terminal, updateText <-chan string) (*segmentdisplay.SegmentDisplay, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The segment display only knows its capacity after both
// text size and terminal size are known.

// Update the capacity initially the first time the
// terminal reports a non-zero size and then every time the
// terminal resizes.
//
// This is better than updating the capacity on every
// iteration since that leads to edge cases - segment
// display capacity depends on the length of text and here
// we are trying to adjust the text length to the capacity.

// newRollText creates a new Text widget that displays rolling text.
func newRollText(ctx context.Context) (*text.Text, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// newSparkLines creates two new sparklines displaying random values.
func newSparkLines(ctx context.Context) (*sparkline.SparkLine, *sparkline.SparkLine, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// newGauge creates a demo Gauge widget.
func newGauge(ctx context.Context) (*gauge.Gauge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// newDonut creates a demo Donut widget.
func newDonut(ctx context.Context) (*donut.Donut, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// newHeartbeat returns a line chart that displays a heartbeat-like progression.
func newHeartbeat(ctx context.Context) (*linechart.LineChart, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// newBarChart returns a BarcChart that displays random values on multiple bars.
func newBarChart(ctx context.Context) (*barchart.BarChart, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// distance is a thread-safe int value used by the newSince method.
// Buttons write it and the line chart reads it.
type distance struct {
	v  int
	mu sync.Mutex
}

// add adds the provided value to the one stored.
func (d *distance) add(v int) { _ = "STUB: not implemented"; return }

// get returns the current value.
func (d *distance) get() int { _ = "STUB: not implemented"; return 0 }

// newSines returns a line chart that displays multiple sine series and two buttons.
// The left button shifts the second series relative to the first series to
// the left and the right button shifts it to the right.
func newSines(ctx context.Context) (left, right *button.Button, lc *linechart.LineChart, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// diff is the difference a single button press adds or removes to the
// second series.

// setLayout sets the specified layout.
func setLayout(c *container.Container, w *widgets, lt layoutType) error {
	_ = "STUB: not implemented"
	return nil
}

// layoutButtons are buttons that change the layout.
type layoutButtons struct {
	allB  *button.Button
	textB *button.Button
	spB   *button.Button
	lcB   *button.Button
}

// newLayoutButtons returns buttons that dynamically switch the layouts.
func newLayoutButtons(c *container.Container, w *widgets) (*layoutButtons, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// rotateFloats returns a new slice with inputs rotated by step.
// I.e. for a step of one:
//
//	inputs[0] -> inputs[len(inputs)-1]
//	inputs[1] -> inputs[0]
//
// And so on.
func rotateFloats(inputs []float64, step int) []float64 { _ = "STUB: not implemented"; return nil }

// rotateRunes returns a new slice with inputs rotated by step.
// I.e. for a step of one:
//
//	inputs[0] -> inputs[len(inputs)-1]
//	inputs[1] -> inputs[0]
//
// And so on.
func rotateRunes(inputs []rune, step int) []rune { _ = "STUB: not implemented"; return nil }
