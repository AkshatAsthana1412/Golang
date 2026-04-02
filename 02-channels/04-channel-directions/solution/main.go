// Solution 4: Channel Directions
//
// Key takeaways:
//   - Channel direction types (<-chan T for receive-only, chan<- T for send-only)
//     enforce at compile time that a function can only read or write a channel.
//     This prevents accidental misuse.
//   - A bidirectional chan T is implicitly convertible to either direction.
//   - Closing a channel signals "no more values." A `for v := range ch` loop
//     will keep receiving until the channel is closed, then exit the loop.
//   - Each stage owns (and closes) its output channel. The producer closes,
//     never the consumer — this is a Go convention.

package main

import "fmt"

func generator(n int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func squarer(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func main() {
	nums := generator(5)
	squares := squarer(nums)

	for v := range squares {
		fmt.Println(v)
	}
}
