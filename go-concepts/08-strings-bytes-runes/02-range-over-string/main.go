// Problem 2: range over a string
//
// `for i, r := range s` decodes UTF-8 rune by rune:
//   - `i` is the BYTE index of where each rune starts
//   - `r` is the rune (int32) value
//
// Contrast with `for i := 0; i < len(s); i++` which iterates BYTES.
//
// Tasks:
//   1. range over "héllo" and print byte index + rune for each step.
//   2. Iterate the same string with `for i := 0; i < len(s); i++` and
//      print s[i] (a byte) at each step. Observe the difference in
//      iteration count.
//
// Run:
//   go run .

package main

import "fmt"

func main() {
	// TODO
	fmt.Println("Implement me.")
}
