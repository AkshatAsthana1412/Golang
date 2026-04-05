// Problem 8: Pipeline
//
// Build a 3-stage processing pipeline connected by channels:
//
//   generate(numbers) → square() → filterOdd() → print results
//
// Stage 1 — generate(nums []int) <-chan int
//   Sends each number from the slice onto a channel, then closes it.
//
// Stage 2 — square(in <-chan int) <-chan int
//   Reads each value from in, squares it, sends on output, then closes output.
//
// Stage 3 — filterOdd(in <-chan int) <-chan int
//   Reads each value from in, sends only odd values, then closes output.
//
// In main, feed []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} and print the results.
//
// Expected output (odd squares of 1-10):
//   1
//   9
//   25
//   49
//   81
//
// Hints:
//   - Each stage is a function that returns a <-chan int (like Problem 4).
//   - Each function internally launches a goroutine to do the work.
//   - Channels provide natural backpressure: if a downstream stage is slow,
//     the upstream stage blocks on send until the downstream is ready.
//
// Run:
//   go run .

package main

import "fmt"

// generate — sends all nums onto a channel, then closes it.
func generate(nums []int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, num := range nums {
			ch <- num
		}
	}()
	return ch
}

// square — reads from in, sends v*v, then closes output.
func square(in <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for num := range in {
			ch <- num * num
		}
	}()
	return ch
}

// filterOdd — reads from in, sends only odd values, then closes output.
func filterOdd(in <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for sq := range in {
			if sq%2 == 1 {
				ch <- sq
			}
		}
	}()
	return ch
}

func main() {
	feed := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for n := range filterOdd(square(generate(feed))) {
		fmt.Println(n)
	}
}
