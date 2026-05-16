// Problem 5: Byte / Rune / Grapheme Counts
//
// Three different counts for "naïve world 🌍":
//   - len(s)                              -> bytes
//   - utf8.RuneCountInString(s)           -> Unicode code points
//   - "graphemes" (user-perceived chars)  -> requires golang.org/x/text/unicode/norm or similar
//
// Tasks:
//   1. Print all three counts (skip the grapheme count — beyond stdlib).
//   2. Define a helper `IsAscii(s string) bool` returning true iff all
//      runes are < 128. Show it for "hello" (true) and "café" (false).
//
// Run:
//   go run .

package main

import "fmt"

func main() {
	// TODO
	fmt.Println("Implement me.")
}
