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

// Binary segmentdisplaydemo shows the functionality of a segment display.
package main

import (
	"context"
	"time"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/segmentdisplay"
)

// clock displays the current time on the segment display.
// Exists when the context expires.
func clock(ctx context.Context, sd *segmentdisplay.SegmentDisplay) {
	_ = "STUB: not implemented"
	return
}

// rotate returns a new slice with inputs rotated by step.
// I.e. for a step of one:
//
//	inputs[0] -> inputs[len(inputs)-1]
//	inputs[1] -> inputs[0]
//
// And so on.
func rotate(inputs []rune, step int) []rune { _ = "STUB: not implemented"; return nil }

// rollText rolls a text across the segment display.
// Exists when the context expires.
func rollText(ctx context.Context, sd *segmentdisplay.SegmentDisplay) {
	_ = "STUB: not implemented"
	return
}

func main() {
	t, err := tcell.New()
	if err != nil {
		panic(err)
	}
	defer t.Close()

	ctx, cancel := context.WithCancel(context.Background())
	clockSD, err := segmentdisplay.New()
	if err != nil {
		panic(err)
	}
	go clock(ctx, clockSD)

	rollingSD, err := segmentdisplay.New()
	if err != nil {
		panic(err)
	}
	go rollText(ctx, rollingSD)

	c, err := container.New(
		t,
		container.Border(linestyle.Light),
		container.BorderTitle("PRESS Q TO QUIT"),
		container.SplitHorizontal(
			container.Top(
				container.PlaceWidget(rollingSD),
			),
			container.Bottom(
				container.PlaceWidget(clockSD),
			),
			container.SplitPercent(40),
		),
	)
	if err != nil {
		panic(err)
	}

	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' {
			cancel()
		}
	}

	if err := termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(quitter), termdash.RedrawInterval(1*time.Second)); err != nil {
		panic(err)
	}
}
