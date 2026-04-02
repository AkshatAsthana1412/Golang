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

// TODO: Implement racyCounter — demonstrates a data race.
// func racyCounter() int { ... }

// TODO: Implement mutexCounter — uses sync.Mutex for safety.
// func mutexCounter() int { ... }

// TODO: Implement channelCounter — uses channels instead of locks.
// func channelCounter() int { ... }

func main() {
	_ = fmt.Println
	_ = sync.Mutex{}

	// TODO: Run each version and print the result.
	// Observe that racyCounter gives wrong results and triggers -race,
	// while the other two always give 10000.
}
