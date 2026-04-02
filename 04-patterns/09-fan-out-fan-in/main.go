// Problem 9: Fan-Out / Fan-In
//
// Extend the pipeline idea: instead of one squarer, run 3 squarer workers in
// parallel (fan-out), then merge their outputs into a single channel (fan-in).
//
//   generate(1..20) ──┬── squarer ──┐
//                     ├── squarer ──┼── merge ── print
//                     └── squarer ──┘
//
// Steps:
//   1. Reuse the `generate` function from Problem 8.
//   2. Launch 3 `square` goroutines, all reading from the SAME input channel.
//      (Multiple goroutines can safely receive from one channel — Go handles this.)
//   3. Write a `merge` function that takes multiple <-chan int and returns a single
//      <-chan int that outputs all values from all inputs.
//   4. Print all merged results.
//
// Hints:
//   - For fan-out: create the input channel once, start N goroutines that each
//     range over it. The runtime distributes values across receivers.
//   - For fan-in (merge): for each input channel, launch a goroutine that forwards
//     values to a shared output channel. Use a WaitGroup to close the output
//     channel once ALL input channels are drained.
//   - Output order will be non-deterministic — that's expected.
//
// Run:
//   go run .

package main

import (
	"fmt"
	"sync"
)

// TODO: Implement generate — sends nums onto channel, then closes it.
// func generate(nums []int) <-chan int { ... }

// TODO: Implement square — reads from in, sends v*v on output, then closes output.
// func square(in <-chan int) <-chan int { ... }

// TODO: Implement merge — combines multiple channels into one.
// func merge(channels ...<-chan int) <-chan int { ... }

func main() {
	_ = fmt.Println
	_ = sync.WaitGroup{}

	// TODO:
	// 1. Generate numbers 1..20
	// 2. Fan-out: launch 3 square workers reading from the same channel
	// 3. Fan-in: merge the 3 output channels
	// 4. Print all results
}
