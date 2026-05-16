// Problem 5: errors.Is vs errors.As
//
// `errors.Is(err, target)` — true if `err` (or anything in its wrap chain)
// IS the target sentinel value. Use when target is a singleton like
// io.EOF or a sentinel `var Err...`.
//
// `errors.As(err, &target)` — true if any error in the chain has the same
// concrete type as `*target`; assigns it. Use when target is a STRUCT
// type with extra fields.
//
// Tasks:
//   1. Build a chain that contains BOTH a sentinel `ErrDB` and a custom
//      `*QueryError{ Query string }` error.
//   2. Show that `errors.Is(err, ErrDB)` returns true.
//   3. Show that `errors.As(err, &qerr)` extracts the *QueryError and
//      its Query field.
//
// To make wrapping work for the custom type, implement `Unwrap() error`
// on *QueryError that returns the wrapped sentinel.
//
// Run:
//   go run .

package main

import "fmt"

// TODO

func main() {
	fmt.Println("Implement me.")
}
