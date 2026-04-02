// Solution 6: Timeout Fetch
//
// Key takeaways:
//   - `select` blocks until one of its cases is ready. It's the idiomatic way
//     to multiplex across channel operations.
//   - `time.After(d)` creates a channel that delivers a value after duration d.
//     When used in a select, it acts as a deadline — whichever case fires first wins.
//   - This pattern is fundamental to writing resilient services: never wait forever
//     for an operation that might hang.
//   - Note: the timed-out goroutine still runs to completion (the channel send has
//     no receiver). In production code, use context.Context for cancellation
//     (covered in Problem 13).

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func simulateFetch() string {
	delay := time.Duration(100+rand.Intn(400)) * time.Millisecond
	time.Sleep(delay)
	return fmt.Sprintf("data (fetched in %v)", delay)
}

func main() {
	resultCh := make(chan string, 1)

	go func() {
		resultCh <- simulateFetch()
	}()

	select {
	case result := <-resultCh:
		fmt.Println("success:", result)
	case <-time.After(300 * time.Millisecond):
		fmt.Println("timeout: fetch took too long")
	}
}
