// Problem 7: Deletion During Iteration
//
// Unlike many languages, Go's spec EXPLICITLY allows deleting from a map
// during a `range` loop. The key being deleted will not be visited again
// (it might or might not have been visited already).
//
// Tasks:
//   1. Build a map of integers to "even"/"odd" by computing.
//   2. While ranging, `delete(m, k)` for every "odd" key.
//   3. Print the resulting map.
//
//   Note: ADDING new keys during a range is NOT guaranteed to be visited
//   in this iteration — they might or might not appear. So deletion is
//   safe; insertion of fresh keys is fragile.
//
// Run:
//   go run .

package main

import "fmt"

func main() {
	// TODO
	fmt.Println("Implement me.")
}
