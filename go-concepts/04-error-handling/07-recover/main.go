// Problem 7: Recover in Defer
//
// `recover()` only does anything when called inside a DEFERRED function
// during a panic. It captures the panic value, stops the unwind, and
// resumes normal control flow at the caller.
//
// Common pattern: a top-level handler in an HTTP middleware that converts
// panics into 500-error responses without crashing the server.
//
// Tasks:
//   1. Write `func safeCall(fn func()) (recovered interface{})` that runs
//      fn and returns whatever was passed to panic, or nil if fn returned
//      normally.
//   2. Demonstrate it with a function that does `panic("boom")`.
//   3. Demonstrate it with a function that explicitly panics with a
//      typed value: `panic(errors.New("structured"))` and recover it
//      back as an `error` via type assertion.
//   4. Note in a comment: recover in a NESTED function call won't catch
//      a panic in the OUTER goroutine — each goroutine needs its own
//      defer/recover.
//
// Run:
//   go run .

package main

import "fmt"

// TODO

func main() {
	fmt.Println("Implement me.")
}
