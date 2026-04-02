// Problem 5: Buffered vs Unbuffered Channels
//
// This exercise has THREE parts. Implement each in its own function.
//
// Part A — bufferedDemo():
//   Create a buffered channel with capacity 3.
//   Send 3 values into it WITHOUT any goroutine receiving.
//   Then receive and print all 3 values.
//   This should work fine because the buffer absorbs the sends.
//
// Part B — unbufferedDeadlock():
//   Create an unbuffered channel.
//   Try to send a value on it from the main goroutine (no other goroutine receiving).
//   This WILL deadlock. Observe the runtime error.
//   (Uncomment the call in main to test, then re-comment before moving on.)
//
// Part C — unbufferedFixed():
//   Create an unbuffered channel.
//   Launch a goroutine that receives from it and prints the value.
//   Send a value on the channel from main.
//   This works because the goroutine is ready to receive.
//
// Hints:
//   - Buffered channel:   make(chan int, 3)
//   - Unbuffered channel: make(chan int)
//   - An unbuffered send blocks until someone receives. If nobody ever will → deadlock.
//
// Run:
//   go run .

package main

import "fmt"

// TODO: Implement bufferedDemo
// func bufferedDemo() { ... }

// TODO: Implement unbufferedDeadlock (will panic with "all goroutines are asleep")
// func unbufferedDeadlock() { ... }

// TODO: Implement unbufferedFixed
// func unbufferedFixed() { ... }

func main() {
	_ = fmt.Println // remove once you use fmt

	// TODO: Uncomment these one at a time to observe behavior:
	// bufferedDemo()
	// unbufferedDeadlock()   // WARNING: this will crash — that's the point!
	// unbufferedFixed()
}
