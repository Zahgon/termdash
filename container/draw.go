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

package container

// draw.go contains logic to draw containers and the contained widgets.

import (
	"image"
)

// drawTree draws this container and all of its sub containers.
func drawTree(c *Container) error { _ = "STUB: not implemented"; return nil }

// drawBorder draws the border around the container if requested.
func drawBorder(c *Container) error { _ = "STUB: not implemented"; return nil }

// drawWidget requests the widget to draw on the canvas.
func drawWidget(c *Container) error { _ = "STUB: not implemented"; return nil }

// drawResize draws an unicode character indicating that the size is too small to draw this container.
// Does nothing if the size is smaller than one cell, leaving no space for the character.
func drawResize(c *Container, area image.Rectangle) error { _ = "STUB: not implemented"; return nil }

// drawCont draws the container and its widget.
func drawCont(c *Container) error { _ = "STUB: not implemented"; return nil }
