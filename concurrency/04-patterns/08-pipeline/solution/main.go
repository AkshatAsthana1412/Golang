// Solution 8: Pipeline
//
// Key takeaways:
//   - A pipeline is a series of stages connected by channels. Each stage is a
//     goroutine that reads from an input channel, processes data, and writes to
//     an output channel.
//   - Backpressure is automatic: if the consumer (downstream) is slow, the
//     producer (upstream) blocks on its channel send. No explicit flow control
//     code is needed.
//   - Each stage owns and closes its output channel. The `range` loop in the
//     next stage terminates when the channel is closed — this is how shutdown
//     propagates through the pipeline.
//   - Pipelines compose naturally: you can add, remove, or reorder stages by
//     changing how the channels are wired in main.

package main

import "fmt"

func generate(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func filterOdd(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			if v%2 != 0 {
				out <- v
			}
		}
		close(out)
	}()
	return out
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ch := filterOdd(square(generate(nums)))

	for v := range ch {
		fmt.Println(v)
	}
}
