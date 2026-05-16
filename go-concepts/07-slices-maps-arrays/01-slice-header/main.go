// Problem 1: Slice Header Anatomy
//
// A slice is NOT an array. It's a 3-word VALUE:
//   - ptr: pointer to the first element of the backing array
//   - len: number of elements visible
//   - cap: capacity (distance from ptr to end of backing array)
//
// Slicing (s[i:j]) does NOT copy — it produces a new header that
// points into the SAME backing array.
//
// Tasks:
//   1. Build `s := []int{0,1,2,3,4,5}`. Print len(s) and cap(s).
//   2. `t := s[1:4]`. Print len(t), cap(t), then mutate t[0] and observe
//      that s changes too.
//   3. `u := s[:3:3]` — this is the "full slice expression" that limits
//      capacity. Show that appending to `u` does NOT clobber s.
//
// Run:
//   go run .

package main

import "fmt"

func main() {
	// TODO
	fmt.Println("Implement me.")
}
