// Problem 6: Struct Comparability
//
// A struct value supports `==` if and only if EVERY field's type is
// comparable. Slices, maps, and functions are NOT comparable, so any
// struct containing them is also not comparable.
//
// Tasks:
//   1. Define `Point { X, Y int }` and confirm `Point{1,2} == Point{1,2}`.
//   2. Define `Tags { values []string }` and observe that `Tags{} == Tags{}`
//      fails to compile.
//   3. Define a struct that embeds an `interface{ Foo() }`. Comparing two
//      values is a RUNTIME comparison and PANICS if the dynamic type is
//      not comparable. Demonstrate by stuffing a slice-bearing type into
//      the interface and comparing.
//
// Why this matters:
//   - Map keys must be comparable, so this rule decides whether your type
//     can be used as a map key.
//   - The interface comparison panic is a real-world gotcha — interview
//     gold.
//
// Run:
//   go run .

package main

import "fmt"

type Point struct{ X, Y int }

// TODO: type Tags ...
// TODO: an interface-bearing struct that may panic on ==

func main() {
	fmt.Println(Point{1, 2} == Point{1, 2})
	// TODO: rest of tasks
	fmt.Println("Implement me.")
}
