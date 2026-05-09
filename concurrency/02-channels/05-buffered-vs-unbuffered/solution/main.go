// Solution 5: Buffered vs Unbuffered Channels
//
// Key takeaways:
//   - Buffered channels decouple sender and receiver in time (up to the buffer
//     capacity). Sends succeed immediately as long as the buffer isn't full.
//   - Unbuffered channels require both sides to be ready at the same instant.
//     A send on an unbuffered channel BLOCKS until another goroutine does a
//     corresponding receive (and vice versa).
//   - If the Go runtime detects that all goroutines are blocked (no progress
//     possible), it reports "fatal error: all goroutines are asleep - deadlock!"
//   - Buffered channels are useful when you know a bounded amount of work will be
//     produced before it's consumed. Unbuffered channels are useful when you want
//     synchronization guarantees (handoff semantics).

package main

import "fmt"

func bufferedDemo() {
	fmt.Println("=== Buffered Demo ===")
	ch := make(chan int, 3)

	// All three sends succeed immediately — buffer absorbs them.
	ch <- 10
	ch <- 20
	ch <- 30

	// Now drain them.
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func unbufferedDeadlock() {
	fmt.Println("=== Unbuffered Deadlock ===")
	ch := make(chan int)

	// This blocks forever: no goroutine will ever receive from ch.
	// The runtime detects the deadlock and panics.
	ch <- 42
}

func unbufferedFixed() {
	fmt.Println("=== Unbuffered Fixed ===")
	ch := make(chan int)

	go func() {
		val := <-ch
		fmt.Println("received:", val)
	}()

	// This send now succeeds because the goroutine above is (or will be) ready to receive.
	ch <- 42
}

func main() {
	bufferedDemo()
	// unbufferedDeadlock() // uncomment to see the deadlock panic
	unbufferedFixed()
}
