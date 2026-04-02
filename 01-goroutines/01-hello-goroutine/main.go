// Problem 1: Hello Goroutine
//
// Launch 5 goroutines, each printing "Hello from goroutine N" (where N is 0..4).
//
// Observe:
//   - The output order is non-deterministic across runs.
//   - If you remove the time.Sleep at the bottom, some (or all) goroutines may
//     never print because main exits before they get a chance to run.
//
// Hints:
//   - Use the `go` keyword to launch a goroutine: go func() { ... }()
//   - The main function IS a goroutine. When it returns, the program exits
//     immediately — it does NOT wait for other goroutines to finish.
//   - time.Sleep is a hacky way to wait; Problem 2 introduces the proper approach.
//
// Run:
//   go run .
//   go run .   (again — notice different ordering)

package main

import (
	"fmt"
	"time"
)

func main() {
	// TODO: Launch 5 goroutines. Each should print "Hello from goroutine N".

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("Hello from goroutine", i)
		}()
	}

	// Temporary sleep so main doesn't exit before goroutines finish.
	// (This is intentionally fragile — Problem 2 shows the right way.)
	time.Sleep(1 * time.Second)
}
