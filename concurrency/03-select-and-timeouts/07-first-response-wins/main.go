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
//
// Concepts:
// - What happens if res is unbuffered?
// If you change the code to res := make(chan string):
// 3 goroutines are started: Each one calls simulateReplica and then tries to send its result to res <- ....
// main waits for one: The fmt.Println(<-res) line blocks until the fastest replica finishes.
// main finishes and exits: Once main receives that first result, it proceeds. In this simple script, the program ends.
// In a real long-running server, main would move on to other tasks.
// The "Slow" Goroutines are stuck: The other 2 replicas were slightly slower. When they finish their work and try to send their result (res <- ...), there is no longer anyone listening (receiving) on that channel.
//
// Perpetual Blocking: Because an unbuffered channel requires a sender and a receiver to be present at the same time, those 2 "slow" goroutines will block on that send line forever.
// They can never finish their function and disappear.

// - What does "goroutine leak" mean?
// A goroutine leak occurs when you start a goroutine that is intended to be temporary, but it never terminates.
// It stays "alive" in memory for the entire duration of your program's life.

// Why it's bad:
// Memory Usage: Each goroutine has its own stack (starting at ~2KB). If you have a leak in a high-traffic web server (e.g., leaking 2 goroutines every request), you will eventually run out of RAM and the process will crash (OOM - Out of Memory).
// Resource Exhaustion: Leaked goroutines might hold onto file descriptors, database connections, or other resources that never get released.
// Garbage Collection: The Go Garbage Collector cannot clean up a goroutine that is blocked waiting on a channel. It assumes the goroutine is still doing "work."

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

	// Query all replicas concurrently.
	// Accept only the first response and print it.
	res := make(chan string, 3)
	for _, replica := range replicas {
		go func() {
			res <- simulateReplica(replica)
		}()
	}

	fmt.Println(<-res)
}
