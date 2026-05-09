// Solution 9: Fan-Out / Fan-In
//
// Key takeaways:
//   - Fan-out: multiple goroutines read from the same channel. The Go runtime
//     delivers each value to exactly one receiver — this is safe and requires
//     no extra synchronization.
//   - Fan-in: a merge function combines multiple channels into one. Each input
//     channel gets a forwarding goroutine; a WaitGroup ensures the output
//     channel is closed only after ALL inputs are drained.
//   - This pattern is ideal when a pipeline stage is CPU-bound and you want to
//     parallelize it across cores.
//   - Output order is non-deterministic because the workers race to process items.

package main

import (
	"fmt"
	"sync"
)

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

func merge(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}()
	}

	// Close the output channel once all forwarding goroutines are done.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	nums := make([]int, 20)
	for i := range nums {
		nums[i] = i + 1
	}

	// Stage 1: generate
	in := generate(nums)

	// Stage 2: fan-out — 3 workers reading from the same channel
	sq1 := square(in)
	sq2 := square(in)
	sq3 := square(in)

	// Stage 3: fan-in — merge all worker outputs
	for v := range merge(sq1, sq2, sq3) {
		fmt.Println(v)
	}
}
