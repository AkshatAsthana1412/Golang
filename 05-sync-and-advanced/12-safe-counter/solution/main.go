// Solution 12: Safe Counter
//
// Key takeaways:
//   - A data race occurs when two goroutines access the same variable concurrently
//     and at least one access is a write. `go run -race` detects this at runtime.
//   - sync.Mutex provides mutual exclusion: only one goroutine can hold the lock
//     at a time. Simple and efficient for protecting shared state.
//   - The channel approach avoids shared state entirely: goroutines communicate
//     increments via a channel, and a single goroutine owns the counter. This
//     follows the Go proverb "share memory by communicating."
//   - Which to use? Mutex is simpler for protect-a-variable cases. Channels are
//     better when the interaction is naturally about communication or pipelines.

package main

import (
	"fmt"
	"sync"
)

const (
	numGoroutines = 10
	numIncrements = 1000
)

func racyCounter() int {
	counter := 0
	var wg sync.WaitGroup

	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range numIncrements {
				counter++ // DATA RACE: unsynchronized read-modify-write
			}
		}()
	}

	wg.Wait()
	return counter
}

func mutexCounter() int {
	counter := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range numIncrements {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	return counter
}

func channelCounter() int {
	ch := make(chan int, numGoroutines)
	var wg sync.WaitGroup

	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range numIncrements {
				ch <- 1
			}
		}()
	}

	// Close channel once all senders are done.
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Single goroutine owns the counter — no race possible.
	total := 0
	for v := range ch {
		total += v
	}
	return total
}

func main() {
	fmt.Println("=== Racy Counter (expect wrong result, triggers -race) ===")
	fmt.Printf("result: %d  (expected: %d)\n\n", racyCounter(), numGoroutines*numIncrements)

	fmt.Println("=== Mutex Counter ===")
	fmt.Printf("result: %d  (expected: %d)\n\n", mutexCounter(), numGoroutines*numIncrements)

	fmt.Println("=== Channel Counter ===")
	fmt.Printf("result: %d  (expected: %d)\n", channelCounter(), numGoroutines*numIncrements)
}
