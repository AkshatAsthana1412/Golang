// Problem 2: The append Aliasing Bug
//
// `append` returns a new slice header. If the backing array has room
// (len < cap), the new header points to the same array. If you forget
// that, two slices can stomp on each other.
//
// Classic bug:
//
//   parent := make([]int, 0, 10)
//   parent = append(parent, 1, 2, 3)
//   childA := append(parent, 100)   // shares backing array!
//   childB := append(parent, 200)   // ALSO shares; overwrites the 100
//
// Tasks:
//   1. Reproduce the above. Print parent, childA, childB.
//   2. Show that the issue goes away if you DEEP COPY parent before
//      appending, OR if you use the full-slice form `parent[:i:i]`.
//
// Run:
//   go run .

package main

import "fmt"

func main() {
	// TODO
	fmt.Println("Implement me.")
}
