// Solution 7: First Response Wins
//
// Key takeaways:
//   - The "hedged request" or "redundant request" pattern: send the same request
//     to multiple backends and use whichever responds first. This reduces tail
//     latency at the cost of extra work.
//   - Using a buffered channel (cap = number of goroutines) ensures that the
//     "losing" goroutines can still send their result without blocking forever.
//     Without the buffer, those goroutines would leak (blocked on send with no
//     receiver). This is a common gotcha.
//   - A bare `<-ch` receive (without select) works perfectly here since we
//     only care about the first value. Using select would be needed if we also
//     wanted a timeout.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func simulateReplica(name string) string {
	delay := time.Duration(50+rand.Intn(450)) * time.Millisecond
	time.Sleep(delay)
	return fmt.Sprintf("%s responded in %v", name, delay)
}

func main() {
	replicas := []string{"replica-A", "replica-B", "replica-C"}

	// Buffered so losing goroutines don't block forever on send.
	resultCh := make(chan string, len(replicas))

	for _, name := range replicas {
		go func() {
			resultCh <- simulateReplica(name)
		}()
	}

	// Only take the first result.
	first := <-resultCh
	fmt.Println("winner:", first)
}
