// Package metrica provides the AtomicCounter implementation, that is used to increment the counter for concurrent requests.
package metrica

import (
	"sync/atomic"
)

// AtomicCounter is a struct to be used as a counter for concurrent requests.
type AtomicCounter struct {
	count int64
}

// NewAtomicCounter returns a new AtomicCounter.
func NewAtomicCounter() *AtomicCounter {
	return &AtomicCounter{}
}

// Inc increments the counter by v.
func (c *AtomicCounter) Inc(v int64) {
	atomic.AddInt64(&c.count, v)
}

// Value returns the current value of the counter.
func (c *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.count)
}
