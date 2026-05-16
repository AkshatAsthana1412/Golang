// Problem 3: Slice as Function Argument
//
// Passing a slice to a function passes the HEADER by value, but the
// header's pointer field still references the same backing array. So:
//   - in-place MUTATION through indexes is visible to the caller
//   - re-slicing or appending inside the func is NOT visible (the local
//     header changes, the caller's doesn't)
//
// Tasks:
//   1. Write `Double(s []int)` that does `s[i] *= 2` for each i. Caller's
//      slice changes.
//   2. Write `AppendOne(s []int)` that does `s = append(s, 99)`. Caller's
//      slice does NOT change.
//   3. Write `AppendOneRet(s []int) []int` that appends and RETURNS — the
//      caller must reassign. This is why `slice = append(slice, ...)` is
//      always a re-assignment.
//
// Run:
//   go run .

package main

import "fmt"

// TODO

func main() {
	fmt.Println("Implement me.")
}
