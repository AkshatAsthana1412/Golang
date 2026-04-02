// Problem 3: Ping Pong
//
// Two goroutines ("ping" and "pong") pass a ball back and forth.
// The ball is an integer counter that increments with each pass.
// The game ends when the counter reaches 10.
//
// Expected output (exact):
//   ping: 1
//   pong: 2
//   ping: 3
//   pong: 4
//   ...
//   pong: 10
//
// Hints:
//   - Use an unbuffered channel for each direction (ping→pong and pong→ping),
//     or a single channel if you prefer.
//   - An unbuffered channel send blocks until the other side receives — this is
//     what enforces the strict alternation.
//   - One goroutine must initiate by sending the first value.
//
// Run:
//   go run .

package main

import "fmt"

func main() {
	// TODO: Create channel(s) and launch two goroutines (ping and pong).
	// Each receives the ball, increments it, prints its role and the count,
	// and sends the ball to the other side.
	// Stop when the count reaches 10.

	_ = fmt.Println // remove once you use fmt
}
