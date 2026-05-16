// Problem 4: Slice Backing-Array Escape
//
// A slice value is a 3-word header (ptr, len, cap). The header itself is
// cheap. The BACKING ARRAY is what may live on the stack or heap.
//
// Tasks:
//   1. Build a small fixed-size slice with a literal: `s := []int{1,2,3}`
//      and use it locally. Confirm with -gcflags="-m" that it does not
//      escape (it might still — depends on inlining).
//   2. `make([]int, 1024)` and use it locally — likely heap-allocated due
//      to size threshold.
//   3. `make([]int, n)` where `n` is unknown at compile time — escapes
//      (size unknown at compile time).
//   4. Returning the slice -> always escapes.
//
// Run:
//   go run .
//   go build -gcflags="-m=2" .

package main

import "fmt"

// TODO

func main() {
	fmt.Println("Implement me.")
}
