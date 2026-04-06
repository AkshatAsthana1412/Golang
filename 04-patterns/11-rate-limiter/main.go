// Problem 11: Rate Limiter
//
// Process a stream of incoming "requests" but limit throughput.
//
// Part A — Steady rate:
//   Process requests at most 5 per second using time.Tick.
//
// Part B — Burst + steady rate:
//   Allow an initial burst of up to 3 requests, then fall back to the steady
//   5-per-second rate. Use a buffered channel pre-filled with tokens as a
//   token bucket.
//
// Hints:
//   - time.Tick(200ms) returns a channel that delivers a value every 200ms
//     (= 5 per second).
//   - For the token bucket: create a buffered channel of cap 3, pre-fill it
//     with 3 values, and refill it with a goroutine using time.Tick. Before
//     processing each request, receive from the bucket (blocks if empty).
//   - Print a timestamp with each request so you can see the timing.
//
// Run:
//   go run .

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

	// TODO Part A: Process requests at a steady 5/sec rate.
	// For each request, wait for a tick, then print the request with time.Now().
	start_time := time.Now()
	// for _, request := range requests {
	// 	<-time.Tick(200 * time.Millisecond)
	// 	fmt.Printf("%s  %s  %s\n", time.Now().Format("15:04:05.000"), request, time.Since(start_time))
	// }

	// TODO Part B: Allow an initial burst of 3, then 5/sec steady state.
	// Create a buffered channel (cap 3), pre-fill with tokens, refill via tick.
	burstyBucket := make(chan struct{}, 3)

	// prefill 3 tokens
	for range 3 {
		burstyBucket <- struct{}{}
	}

	// go routine to refill burstyBucket, because this is what the main goroutine is blocking on, so we have to simulate the ticker with it.
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()
	go func() {
		for range ticker.C {
			select {
			case burstyBucket <- struct{}{}:
			default:
			}
		}
	}()

	for _, request := range requests {
		<-burstyBucket
		fmt.Printf("%s  %s  %s\n", time.Now().Format("15:04:05.000"), request, time.Since(start_time))
	}
}
