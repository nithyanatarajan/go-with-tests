package synch_test

import (
	"sync"
	"testing"

	"github.com/nithyanatarajan/go-with-tests/synch"
	. "github.com/nithyanatarajan/go-with-tests/test"
)

func TestCounter(t *testing.T) {
	t.Run("should set counter to 3 when incremented by 3", func(t *testing.T) {
		counter := synch.NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("should increment safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := synch.NewCounter()

		var wg sync.WaitGroup
		// Tell the WaitGroup to expect 1000 goroutines to call wg.Done()
		// to signal their completion.
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				// Decrement the WaitGroup counter when the function completes
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t *testing.T, got *synch.Counter, want int) {
	AssertEqual(t, got.Value(), want)
}
