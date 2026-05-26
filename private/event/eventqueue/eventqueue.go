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

// Package eventqueue provides an unboud FIFO queue of events.
package eventqueue

import (
	"context"
	"sync"

	"github.com/mum4k/termdash/terminal/terminalapi"
)

// node is a single data item on the queue.
type node struct {
	prev  *node
	next  *node
	event terminalapi.Event
}

// Unbound is an unbound FIFO queue of terminal events.
// Unbound must not be copied, pass it by reference only.
// This implementation is thread-safe.
type Unbound struct {
	first *node
	last  *node
	// mu protects first and last.
	mu sync.Mutex

	// cond is used to notify any callers waiting on a call to Pull().
	cond *sync.Cond

	// condMU protects cond.
	condMU sync.RWMutex

	// done is closed when the queue isn't needed anymore.
	done chan struct{}
}

// New returns a new Unbound queue of terminal events.
// Call Close() when done with the queue.
func New() *Unbound { _ = "STUB: not implemented"; return nil }

// Stops when Close() is called.

// wake periodically wakes up all goroutines waiting at Pull() so they can
// check if the context expired.
func (u *Unbound) wake() { _ = "STUB: not implemented"; return }

// Empty determines if the queue is empty.
func (u *Unbound) Empty() bool { _ = "STUB: not implemented"; return false }

// empty determines if the queue is empty.
func (u *Unbound) empty() bool { _ = "STUB: not implemented"; return false }

// Push pushes an event onto the queue.
func (u *Unbound) Push(e terminalapi.Event) { _ = "STUB: not implemented"; return }

// push is the implementation of Push.
// Caller must hold u.mu.
func (u *Unbound) push(e terminalapi.Event) { _ = "STUB: not implemented"; return }

// Pop pops an event from the queue. Returns nil if the queue is empty.
func (u *Unbound) Pop() terminalapi.Event {
	_ = "STUB: not implemented"
	return *new(terminalapi.Event)
}

// Pull is like Pop(), but blocks until an item is available or the context
// expires. Returns a nil event if the context expired.
func (u *Unbound) Pull(ctx context.Context) terminalapi.Event {
	_ = "STUB: not implemented"
	return *new(terminalapi.Event)
}

// Close should be called when the queue isn't needed anymore.
func (u *Unbound) Close() {
	_ = "STUB: not implemented"

	// Throttled is an unbound and throttled FIFO queue of terminal events.
	// Throttled must not be copied, pass it by reference only.
	// This implementation is thread-safe.
	return
}

type Throttled struct {
	queue *Unbound
	max   int
}

// NewThrottled returns a new Throttled queue of terminal events.
//
// This queue scans the queue content on each Push call and won't Push the
// event if there already is a continuous chain of exactly the same events
// en queued. The argument maxRep specifies the maximum number of repetitive
// events.
//
// Call Close() when done with the queue.
func NewThrottled(maxRep int) *Throttled { _ = "STUB: not implemented"; return nil }

// Empty determines if the queue is empty.
func (t *Throttled) Empty() bool { _ = "STUB: not implemented"; return false }

// Push pushes an event onto the queue.
func (t *Throttled) Push(e terminalapi.Event) { _ = "STUB: not implemented"; return }

// Drop the repetitive event.

// Pop pops an event from the queue. Returns nil if the queue is empty.
func (t *Throttled) Pop() terminalapi.Event {
	_ = "STUB: not implemented"
	return *

	// Pull is like Pop(), but blocks until an item is available or the context
	// expires. Returns a nil event if the context expired.
	new(terminalapi.Event)
}

func (t *Throttled) Pull(ctx context.Context) terminalapi.Event {
	_ = "STUB: not implemented"
	return *

	// Close should be called when the queue isn't needed anymore.
	new(terminalapi.Event)
}

func (t *Throttled) Close() { _ = "STUB: not implemented"; return }
