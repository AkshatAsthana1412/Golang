// Problem 8: The Nil Interface Trap
//
// An interface value has TWO words: a TYPE and a VALUE.
// An interface is `nil` ONLY when BOTH are nil.
//
// If you store a TYPED nil pointer into an interface, the interface's type
// word is non-nil, so `iface == nil` returns FALSE — even though the
// underlying value is a nil pointer.
//
// Tasks:
//   1. Define a `MyError` type with `func (e *MyError) Error() string`.
//   2. Write `func DoStuff() error` that has a local `var e *MyError` and
//      RETURNS `e` (a typed nil). Show that the caller's `if err != nil`
//      branch is unexpectedly taken.
//   3. Fix the function — return the bare `nil` literal (or `error(nil)`)
//      when there is no error.
//
// Why this matters:
//   - Real-world bug source. Many Go projects have CI lints for this.
//
// Run:
//   go run .

package main

import "fmt"

// TODO: MyError + Error() method
// TODO: DoStuff (buggy and fixed)

func main() {
	fmt.Println("Implement me.")
}
