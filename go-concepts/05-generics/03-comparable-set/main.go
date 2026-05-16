// Problem 3: The `comparable` Constraint
//
// `comparable` is a built-in constraint satisfied by any type that
// supports `==`. It's required when you use `T` as a map key.
//
// Tasks:
//   1. Implement `Unique[T comparable](in []T) []T` that returns elements
//      in original order, with duplicates removed.
//   2. Implement `Set[T comparable]` as a tiny generic set:
//        - NewSet[T]() *Set[T]
//        - Add(v T)
//        - Has(v T) bool
//        - Len() int
//   3. Use it to count unique words in a slice.
//
// Run:
//   go run .

package main

import "fmt"

// TODO: Unique, Set, methods

func main() {
	fmt.Println("Implement me.")
}
