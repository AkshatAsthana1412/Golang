// Problem 3: Custom Error Types
//
// Sometimes the caller needs MORE than a sentinel — they need structured
// data about the failure. Define a struct that satisfies `error`.
//
// Tasks:
//   1. Define:
//        type ValidationError struct {
//          Field   string
//          Message string
//        }
//      Implement Error() string -> "validation: <field>: <message>".
//
//   2. Write `func ValidateUser(name, email string) error` that returns
//      a *ValidationError pointer when something is wrong.
//
//   3. In main(), use `errors.As` to extract the *ValidationError from a
//      returned error and print its Field/Message separately.
//
// Run:
//   go run .

package main

import "fmt"

// TODO

func main() {
	fmt.Println("Implement me.")
}
