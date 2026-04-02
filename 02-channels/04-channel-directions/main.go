// Problem 4: Channel Directions
//
// Build a two-stage number processing chain:
//   1. A `generator` function that produces integers 1..n on a channel.
//   2. A `squarer` function that reads from an input channel, squares each value,
//      and sends the result on an output channel.
//
// Wire them together in main and print each squared value.
//
// Expected output for n=5:
//   1
//   4
//   9
//   16
//   25
//
// Hints:
//   - Use channel direction types in the function signatures:
//       func generator(n int) <-chan int          // returns receive-only channel
//       func squarer(in <-chan int) <-chan int     // takes receive-only, returns receive-only
//   - Inside generator, launch a goroutine that sends 1..n then closes the channel.
//   - Inside squarer, launch a goroutine that ranges over `in`, squares, sends, then closes output.
//   - In main, range over the final channel to print results.
//   - Closing a channel causes `range` over it to terminate — this is how the
//     pipeline shuts down cleanly.
//
// Run:
//   go run .

package main

import "fmt"

// TODO: Implement generator — returns a <-chan int that produces 1..n, then closes.
// func generator(n int) <-chan int { ... }
func generator(n int) <-chan int {
	nums := make(chan int, 10)
	go func() {
		defer close(nums)
		for i := range n {
			nums <- i
		}
	}()
	return nums
}

// TODO: Implement squarer — reads from in, squares each value, sends on output, then closes.
// func squarer(in <-chan int) <-chan int { ... }
func squarer(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num * num
		}
	}()
	return out
}

func main() {

	for num := range squarer(generator(5)) {
		fmt.Println(num)
	}
	// for sq := range squarer(generator(5)) {
	// 	fmt.Println(sq)
	// }
}
