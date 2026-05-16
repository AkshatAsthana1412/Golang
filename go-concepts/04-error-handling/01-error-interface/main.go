// Problem 1: The `error` Interface
//
// `error` is just a tiny built-in interface:
//
//   type error interface { Error() string }
//
// Anything with an `Error() string` method satisfies it. There is nothing
// magical about it.
//
// Tasks:
//   1. Define a type `EmptyInputError struct{}` with `Error() string`.
//   2. Write `func ParseDigits(s string) (int, error)` that returns an
//      `EmptyInputError` when s is empty, otherwise sums each digit's value.
//   3. Demonstrate the canonical `if err != nil { ... }` pattern in main().
//
// Run:
//   go run .

package main

import "fmt"

// TODO: EmptyInputError, ParseDigits

func main() {
	fmt.Println("Implement me.")
}
