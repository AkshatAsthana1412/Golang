// Problem 3: []byte ↔ string Conversions
//
// `string(b)` and `[]byte(s)` are CONVERSIONS — they ALLOCATE because
// strings are immutable but []byte is mutable. The runtime can't safely
// share the backing memory in general.
//
// (There are tiny exceptions where the compiler optimizes the cost away —
// e.g., string-to-[]byte feeding directly into `range` — but assume an
// alloc happens.)
//
// Tasks:
//   1. Convert s := "hello" to []byte, mutate it, convert back, and show
//      the original `s` is unchanged (immutability).
//   2. Show that `len(s) == len([]byte(s))` (BYTES, not runes).
//   3. Use `strings.Builder` (next problem) when you'd otherwise loop
//      `s = s + ...`.
//
// Run:
//   go run .

package main

import "fmt"

func main() {
	// TODO
	fmt.Println("Implement me.")
}
