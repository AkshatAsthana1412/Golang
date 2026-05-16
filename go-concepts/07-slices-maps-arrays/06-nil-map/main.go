// Problem 6: Nil Map — Reads vs Writes
//
// A nil map READS as if it's empty: `v, ok := m[key]` returns the zero
// value of V and ok=false. No panic.
//
// A nil map WRITE panics: "assignment to entry in nil map".
//
// Tasks:
//   1. `var m map[string]int`. Print `m["nope"]` and `m == nil`.
//   2. `_, ok := m["nope"]` — print `ok`.
//   3. Try `m["nope"] = 1` inside a recover() and print the panic msg.
//   4. After making the map (`m = map[string]int{}`), the same write works.
//
// Run:
//   go run .

package main

import "fmt"

func main() {
	// TODO
	fmt.Println("Implement me.")
}
