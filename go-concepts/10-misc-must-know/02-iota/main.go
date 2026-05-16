// Problem 2: iota Patterns
//
// `iota` is a counter that starts at 0 in each `const ( ... )` block and
// increments by 1 per ConstSpec. Combine with expressions for richer
// patterns:
//
//   const (
//     Sunday = iota // 0
//     Monday        // 1
//     ...
//   )
//
//   const (
//     KB = 1 << (10 * (iota + 1)) // 1024
//     MB                          // 1048576
//     GB                          // ...
//   )
//
// Tasks:
//   1. Define a `Weekday` enum with iota.
//   2. Define byte-size constants KB, MB, GB, TB using bit-shift + iota.
//   3. Define a bit-flag set with 1 << iota:
//        ReadFlag, WriteFlag, ExecFlag (1, 2, 4)
//      Combine with bitwise-or and test with `&`.
//
// Run:
//   go run .

package main

import "fmt"

func main() {
	// TODO
	fmt.Println("Implement me.")
}
