// Problem 1: Stack vs Heap
//
// Go's compiler decides per allocation whether a value lives on the
// goroutine's STACK (cheap, freed by frame teardown) or the HEAP (managed
// by the GC). The decision is "escape analysis":
//   - If a value can be proven not to outlive its function, it goes on
//     the stack.
//   - Otherwise (escapes via return, escapes via interface, escapes via
//     stored pointer, ...) it goes on the heap.
//
// You can SEE the compiler's decisions:
//
//   go build -gcflags="-m" .
//
// Tasks:
//   1. Write `localPoint() Point` that returns a Point VALUE.
//   2. Write `escapingPoint() *Point` that returns a pointer to a local.
//   3. Write `viaInterface(p any)` (callers will pass a Point by value;
//      observe escape due to interface boxing).
//   4. From the directory, run `go build -gcflags="-m" ./solution` and
//      compare against your version. Note "moved to heap" vs no message.
//
// Run:
//   go run .
//   go build -gcflags="-m" .

package main

import "fmt"

type Point struct{ X, Y int }

// TODO

func main() {
	fmt.Println("Implement me. (Try: go build -gcflags=\"-m\" .)")
}
