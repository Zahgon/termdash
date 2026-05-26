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

// Binary formdemo creates a form that accepts text inputs and supports
// keyboard navigation.
package main

import (
	"context"
	"time"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/textinput"
)

// buttonChunks creates the text chunks for a button from the provided text.
func buttonChunks(text string) []*button.TextChunk { _ = "STUB: not implemented"; return nil }

// form contains the elements of a text input form.
type form struct {
	// userInput is a text input that accepts user name.
	userInput *textinput.TextInput
	// uidInput is a text input that accepts UID.
	uidInput *textinput.TextInput
	// gidInput is a text input that accepts GID.
	gidInput *textinput.TextInput
	// homeInput is a text input that accepts path to the home folder.
	homeInput *textinput.TextInput

	// submitB is a button that submits the form.
	submitB *button.Button
	// cancelB is a button that exist the application.
	cancelB *button.Button
}

// newForm returns a new form instance.
// The cancel argument is a function that terminates the application when called.
func newForm(cancel context.CancelFunc) (*form, error) { _ = "STUB: not implemented"; return nil, nil }

// formLayout updates the container into a layout with text inputs and buttons.
func formLayout(c *container.Container, f *form) error { _ = "STUB: not implemented"; return nil }

// submitLayout updates the container into a layout that displays the submitted data.
// The cancel argument is a function that terminates Termdash when called.
func submitLayout(c *container.Container, f *form, cancel context.CancelFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func main() {
	t, err := tcell.New()
	if err != nil {
		panic(err)
	}
	defer t.Close()

	ctx, cancel := context.WithCancel(context.Background())
	c, err := container.New(t, container.ID("root"))
	if err != nil {
		panic(err)
	}

	f, err := newForm(cancel)
	if err != nil {
		panic(err)
	}
	f.submitB.SetCallback(func() error {
		return submitLayout(c, f, cancel)
	})
	if err := formLayout(c, f); err != nil {
		panic(err)
	}

	if err := termdash.Run(ctx, t, c, termdash.RedrawInterval(100*time.Millisecond)); err != nil {
		panic(err)
	}
}
