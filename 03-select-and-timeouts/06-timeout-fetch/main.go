// Problem 6: Timeout Fetch
//
// Simulate fetching data from a slow API. The fetch takes a random 100–500ms.
// Use `select` to either return the result or time out after 300ms.
//
// Expected behavior:
//   - If the simulated fetch completes within 300ms, print the result.
//   - If it takes longer, print "timeout: fetch took too long".
//   - Run multiple times to see both outcomes.
//
// Hints:
//   - Launch the fetch in a goroutine that sends its result on a channel.
//   - Use select with two cases:
//       case result := <-resultCh:   // fetch completed in time
//       case <-time.After(300ms):    // timeout fired first
//   - `time.After(d)` returns a <-chan time.Time that receives a value after
//     duration d. It acts as a one-shot timer in a select.
//
// Run:
//   go run .

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// simulateFetch pretends to call a slow API.
func simulateFetch() string {
	delay := time.Duration(100+rand.Intn(400)) * time.Millisecond
	time.Sleep(delay)
	return fmt.Sprintf("data (fetched in %v)", delay)
}

func main() {
	// TODO: Launch simulateFetch in a goroutine, sending the result on a channel.
	// Use select to either receive the result or time out after 300ms.
	ch := make(chan string)

	go func() {
		ch <- simulateFetch()
	}()

	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(300 * time.Millisecond):
		fmt.Println("Timeout!")
	}
}
