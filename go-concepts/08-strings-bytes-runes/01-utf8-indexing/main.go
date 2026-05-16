// Problem 1: UTF-8 Indexing Pitfall
//
// In Go, `string` is an immutable sequence of BYTES (not runes/chars).
// Indexing s[i] returns a byte, even if the string contains multi-byte
// UTF-8 characters.
//
//   s := "héllo"
//   len(s)   // 6 — "é" is 2 bytes in UTF-8
//   s[1]     // first byte of "é", not the character
//
// Tasks:
//   1. Print len("héllo世界") and len([]rune("héllo世界")). Note the difference.
//   2. Print s[1] in two ways: %d and %c. Observe that you've printed
//      part of a multi-byte sequence.
//   3. Build a function `runeAt(s string, i int) rune` that returns the
//      i-th RUNE (decoded character) using utf8.DecodeRuneInString in a
//      loop, or by converting to []rune (slow but simple).
//
// Run:
//   go run .

package main

import "fmt"

func main() {
	// TODO
	fmt.Println("Implement me.")
}
