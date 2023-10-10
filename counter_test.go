package metrica

import (
	"sync"
	"testing"
)

func TestAtomicCounter(t *testing.T) {
	var wg sync.WaitGroup
	c := NewAtomicCounter()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			c.Inc(1)
			defer wg.Done()
		}()
	}
	wg.Wait()

	assert(t, 100, int(c.Value()))
}
