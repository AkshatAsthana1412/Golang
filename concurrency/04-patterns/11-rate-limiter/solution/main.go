// Solution 11: Rate Limiter
//
// Key takeaways:
//   - time.Tick(d) returns a channel that delivers a value every d. Receiving
//     from it before each operation naturally throttles throughput to 1/d ops/sec.
//   - A buffered channel pre-filled with tokens acts as a token bucket: receives
//     succeed immediately while tokens are available (burst), then block until
//     the refiller goroutine adds more (steady state).
//   - This is the same idea behind production rate limiters (e.g. golang.org/x/time/rate),
//     just implemented with raw channels for learning purposes.
//   - Note: time.Tick leaks the underlying ticker. In long-running programs,
//     use time.NewTicker and call ticker.Stop() when done.

package main

import (
	"fmt"
	"time"
)

func main() {
	requests := []string{
		"req-1", "req-2", "req-3", "req-4", "req-5",
		"req-6", "req-7", "req-8", "req-9", "req-10",
	}

	// --- Part A: Steady rate (5 per second) ---
	fmt.Println("=== Part A: Steady 5/sec ===")
	limiter := time.Tick(200 * time.Millisecond)

	for _, req := range requests {
		<-limiter
		fmt.Printf("%s  %s\n", time.Now().Format("15:04:05.000"), req)
	}

	// --- Part B: Burst of 3, then 5/sec ---
	fmt.Println("\n=== Part B: Burst 3 + Steady 5/sec ===")
	burstyBucket := make(chan struct{}, 3)

	// Pre-fill: 3 tokens available immediately.
	for range 3 {
		burstyBucket <- struct{}{}
	}

	// Refill at steady rate.
	go func() {
		for range time.Tick(200 * time.Millisecond) {
			select {
			case burstyBucket <- struct{}{}:
			default: // bucket full, discard token
			}
		}
	}()

	for _, req := range requests {
		<-burstyBucket
		fmt.Printf("%s  %s\n", time.Now().Format("15:04:05.000"), req)
	}
}
