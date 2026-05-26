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

// hv_line_graph.go helps to keep track of locations where lines cross.

import (
	"image"

	"github.com/mum4k/termdash/linestyle"
)

// hVLineEdge is an edge between two points on the graph.
type hVLineEdge struct {
	// from is the starting node of this edge.
	// From is guaranteed to be less than to.
	from image.Point

	// to is the ending point of this edge.
	to image.Point
}

// newHVLineEdge returns a new edge between the two points.
func newHVLineEdge(from, to image.Point) hVLineEdge {
	_ = "STUB: not implemented"
	return *new(hVLineEdge)
}

// hVLineNode represents one node in the graph.
// I.e. one cell.
type hVLineNode struct {
	// p is the point where this node is.
	p image.Point

	// edges are the edges between this node and the surrounding nodes.
	// The code only supports horizontal and vertical lines so there can only
	// ever be edges to nodes on these planes.
	edges map[hVLineEdge]bool
}

// newHVLineNode creates a new newHVLineNode.
func newHVLineNode(p image.Point) *hVLineNode { _ = "STUB: not implemented"; return nil }

// hasDown determines if this node has an edge to the one below it.
func (n *hVLineNode) hasDown() bool { _ = "STUB: not implemented"; return false }

// hasUp determines if this node has an edge to the one above it.
func (n *hVLineNode) hasUp() bool { _ = "STUB: not implemented"; return false }

// hasLeft determines if this node has an edge to the next node on the left.
func (n *hVLineNode) hasLeft() bool { _ = "STUB: not implemented"; return false }

// hasRight determines if this node has an edge to the next node on the right.
func (n *hVLineNode) hasRight() bool { _ = "STUB: not implemented"; return false }

// rune, given the selected line style returns the correct line character to
// represent this node.
// Only handles nodes with two or more edges, as returned by multiEdgeNodes().
func (n *hVLineNode) rune(ls linestyle.LineStyle) (rune, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// hVLineGraph represents lines on the canvas as a bidirectional graph of
// nodes. Helps to determine the characters that should be used where multiple
// lines cross.
type hVLineGraph struct {
	nodes map[image.Point]*hVLineNode
}

// newHVLineGraph creates a new hVLineGraph.
func newHVLineGraph() *hVLineGraph { _ = "STUB: not implemented"; return nil }

// getOrCreateNode gets an existing or creates a new node for the point.
func (g *hVLineGraph) getOrCreateNode(p image.Point) *hVLineNode {
	_ = "STUB: not implemented"
	return nil
}

// addLine adds a line to the graph.
// This adds edges between all the points on the line.
func (g *hVLineGraph) addLine(line *hVLine) { _ = "STUB: not implemented"; return }

// multiEdgeNodes returns all nodes that have more than one edge.  These are
// the nodes where we might need to use different line characters to represent
// the crossing of multiple lines.
func (g *hVLineGraph) multiEdgeNodes() []*hVLineNode { _ = "STUB: not implemented"; return nil }
