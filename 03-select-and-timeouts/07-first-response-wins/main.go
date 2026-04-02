// Problem 7: First Response Wins
//
// Send the same request to 3 simulated "replicas" (goroutines), each with
// random latency. Accept whichever one responds first and discard the rest.
//
// Expected behavior:
//   - Print which replica responded and how long it took.
//   - Only one result is printed even though 3 replicas were queried.
//
// Hints:
//   - Create a single result channel (buffered with cap 3 so slow goroutines
//     don't leak — their sends succeed into the buffer even after main moves on).
//   - Launch 3 goroutines, each calling simulateReplica with a different name.
//   - Use a single <-resultCh receive (or a select) to get the first result.
//
// Run:
//   go run .

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// simulateReplica simulates a replica with random latency.
func simulateReplica(name string) string {
	delay := time.Duration(50+rand.Intn(450)) * time.Millisecond
	time.Sleep(delay)
	return fmt.Sprintf("%s responded in %v", name, delay)
}

func main() {
	replicas := []string{"replica-A", "replica-B", "replica-C"}

	// TODO: Query all replicas concurrently.
	// Accept only the first response and print it.

	_ = replicas        // remove once you use replicas
	_ = fmt.Println     // remove once you use fmt
	_ = simulateReplica // remove once you use simulateReplica
}
