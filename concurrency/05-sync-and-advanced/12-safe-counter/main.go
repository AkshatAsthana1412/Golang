// Problem 12: Safe Counter
//
// Multiple goroutines increment a shared counter 1000 times each.
// Implement THREE versions to explore different synchronization strategies.
//
// Part A — Racy (intentionally broken):
//   Increment a plain int from 10 goroutines, 1000 times each.
//   Run with `go run -race .` to see the race detector complain.
//   The final count will likely be less than 10000.
//
// Part B — Mutex:
//   Protect the counter with sync.Mutex so increments are atomic.
//   The final count should always be exactly 10000.
//
// Part C — Channel:
//   Instead of a shared variable, use a channel: goroutines send +1 on a channel,
//   and a single goroutine accumulates the total. No lock needed.
//
// Hints:
//   - Part A: just `counter++` from each goroutine — no protection.
//   - Part B: `mu.Lock(); counter++; mu.Unlock()` (or defer mu.Unlock()).
//   - Part C: make(chan int), each goroutine sends 1, a collector reads all values.
//
// Run:
//   go run -race .

package main

import (
	"fmt"
	"sync"
)

const (
	numGoroutines = 10
	numIncrements = 1000
)

// TODO: Implement racyCounter — demonstrates a data race.
func racyCounter() int {
	wg := sync.WaitGroup{}
	counter := 0
	for range numGoroutines {
		wg.Add(1)
		go func() {
			for range numIncrements {
				counter++
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return counter
}

// TODO: Implement mutexCounter — uses sync.Mutex for safety.
func mutexCounter() int {
	wg := sync.WaitGroup{}
	var mu sync.Mutex
	counter := 0
	for range numGoroutines {
		wg.Add(1)
		go func() {
			for range numIncrements {
				mu.Lock()
				counter++
				mu.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return counter
}

// TODO: Implement channelCounter — uses channels instead of locks.
func channelCounter() int {
	out := make(chan int)
	wg := sync.WaitGroup{}
	counter := 0
	// launch the goroutines to send 1 to the out channel
	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range numIncrements {
				out <- 1
			}
		}()
	}

	// close the channel once all senders are done
	go func() {
		wg.Wait()
		close(out)
	}()

	// collect the results, main goroutine blocks here until the out channel is closed
	for v := range out {
		counter += v
	}

	return counter
}

func main() {
	_ = sync.Mutex{}
	// TODO: Run each version and print the result.
	// Observe that racyCounter gives wrong results and triggers -race,
	// while the other two always give 10000.
	fmt.Printf("Racy counter value: %d\n", racyCounter())
	fmt.Printf("Mutex Counter value: %d\n", mutexCounter())
	fmt.Printf("Mutex Counter value: %d\n", channelCounter())
}
