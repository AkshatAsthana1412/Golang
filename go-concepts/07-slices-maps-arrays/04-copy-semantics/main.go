// Problem 4: copy() Semantics
//
// Built-in `copy(dst, src)` copies min(len(dst), len(src)) elements and
// returns that count. It does NOT grow `dst`. To duplicate a slice you
// must size the destination first.
//
// Tasks:
//   1. `dst := make([]int, 5)` and `copy(dst, []int{1,2,3})`. Print dst
//      and the return value.
//   2. `dst := make([]int, 2)` and copy a 5-element source. Show that
//      only the first 2 are copied.
//   3. Cleanly duplicate a slice: `dup := make([]T, len(src)); copy(dup, src)`.
//      Then mutate dup and confirm src is untouched.
//   4. Bonus: `copy(b, "hello")` where b is []byte works because of a
//      special case for byte slice <- string.
//
// Run:
//   go run .

package main

import "fmt"

func main() {
	// TODO
	fmt.Println("Implement me.")
}
