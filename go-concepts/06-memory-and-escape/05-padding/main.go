// Problem 5: Struct Padding & Alignment
//
// Fields in a struct are laid out in declared order, with PADDING inserted
// to satisfy each field's alignment requirement. Reordering fields can
// shrink a struct considerably.
//
// Tasks:
//   1. Define two structs with the same fields in different orders:
//        type Bad  struct { A bool; B int64; C bool; D int64 }
//        type Good struct { A bool; C bool; B int64; D int64 }
//   2. Print each size with `unsafe.Sizeof`. Bad will be larger because
//      of padding holes after each bool to align the int64s.
//   3. Print Alignof for both bool and int64.
//
// Run:
//   go run .

package main

import "fmt"

// TODO

func main() {
	fmt.Println("Implement me.")
}
