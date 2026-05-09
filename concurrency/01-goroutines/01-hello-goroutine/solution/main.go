// Solution 1: Hello Goroutine
//
// Key takeaways:
//   - `go f()` launches f in a new goroutine — a lightweight thread managed by
//     the Go runtime, not an OS thread.
//   - Goroutines run concurrently. The scheduler decides the execution order,
//     which is why output appears in a different order each run.
//   - When main() returns, ALL goroutines are killed immediately. The
//     time.Sleep here is a deliberate hack to demonstrate the problem —
//     the next exercise introduces sync.WaitGroup for proper synchronization.

package main

import (
	"fmt"
	"time"
)

func main() {
	for i := range 5 {
		go func() {
			fmt.Printf("Hello from goroutine %d\n", i)
		}()
	}

	// Without this sleep, main exits instantly and you'll see zero or partial output.
	// Try commenting it out to verify!
	time.Sleep(1 * time.Second)
}
