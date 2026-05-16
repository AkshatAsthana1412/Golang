// Problem 8: Multi-Errors with `errors.Join` (Go 1.20+)
//
// `errors.Join(err1, err2, ...)` returns a single error whose message
// joins each non-nil error on a newline, AND whose Is/As chain still
// contains every input error. Perfect for validations that should report
// EVERYTHING wrong rather than bail on the first failure.
//
// Tasks:
//   1. Write `ValidateForm(name, email, age string) error` that collects
//      ALL the validation errors (empty name, missing @, non-numeric or
//      negative age) and returns them via errors.Join.
//   2. In main(), call with a fully bad input. Print the joined error.
//      Then call `errors.Is(err, ErrEmptyName)` to confirm the chain still
//      identifies individual errors.
//
// Run:
//   go run .

package main

import "fmt"

// TODO

func main() {
	fmt.Println("Implement me.")
}
