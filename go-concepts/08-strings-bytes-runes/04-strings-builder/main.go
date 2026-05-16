// Problem 4: strings.Builder
//
// `strings.Builder` accumulates writes into an internal []byte buffer and
// produces the final string with `String()` — a single allocation at the
// end (or zero, if you preallocate with `Grow(n)`).
//
// Tasks:
//   1. Use a strings.Builder to join 1..1000 with commas. (For interview
//      flavor, try the same with naive `s += fmt.Sprint(i)+","` and see
//      the impact in `time` if you're curious.)
//   2. Use `b.Grow(estimatedBytes)` to preallocate.
//   3. Note that you must NOT copy a strings.Builder — it has a noCopy
//      sentinel. If you do, you'll get a vet warning.
//
// Run:
//   go run .

package main

import "fmt"

func main() {
	// TODO
	fmt.Println("Implement me.")
}
