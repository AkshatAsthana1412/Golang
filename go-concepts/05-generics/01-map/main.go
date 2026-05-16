// Problem 1: Generic Map
//
// Go got generics in 1.18. Type parameters go in `[T any]` after the
// function name:
//
//   func Map[T, U any](in []T, f func(T) U) []U { ... }
//
// Tasks:
//   1. Implement `Map[T, U any](in []T, f func(T) U) []U`.
//   2. Use it to:
//      - turn []int{1,2,3} into []int{1,4,9}
//      - turn []string{"go","is","fast"} into []int (lengths)
//
// Run:
//   go run .

package main

import "fmt"

// TODO: Map

func main() {
	fmt.Println("Implement me.")
}
