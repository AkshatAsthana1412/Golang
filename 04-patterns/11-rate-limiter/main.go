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
	"math/rand"
	"sync"
	"time"
)

func processRequest(req string) {
	start_time := time.Now()
	time.Sleep(time.Duration(200+10*rand.Intn(50)) * time.Millisecond)
	fmt.Printf("Processed: %s successfully! time elapsed: %s\n", req, time.Since(start_time).Round(time.Millisecond))
}

func main() {
	requests := []string{
		"req-1", "req-2", "req-3", "req-4", "req-5",
		"req-6", "req-7", "req-8", "req-9", "req-10",
	}

	// TODO Part A: Process requests at a steady 5/sec rate.
	// For each request, wait for a tick, then print the request with time.Now().
	// for _, req := range requests {
	// 	<-ticker.C
	// 	fmt.Println("Processing ", req, "at ", time.Since(start_time).Round(time.Millisecond))
	// }

	bursty_bucket := make(chan struct{}, 3)
	for range 3 {
		bursty_bucket <- struct{}{}
	}

	start_time := time.Now()
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()
	var wg sync.WaitGroup

	go func() {
		for range ticker.C {
			select {
			case bursty_bucket <- struct{}{}:
			default:
			}

		}
	}()

	for _, req := range requests {
		wg.Add(1)
		<-bursty_bucket
		fmt.Printf("%s, triggered at %s\n", req, time.Since(start_time).Round(time.Millisecond))
		go func() {
			defer wg.Done()
			processRequest(req)
		}()
	}
	wg.Wait()
}
