// Problem 8: Type Inference & Its Limits
//
// Go infers type parameters from FUNCTION ARGUMENTS, not from return types
// or variable annotations. There are still cases where you must spell out
// the type parameter.
//
// Tasks:
//   1. `Map[T, U any]` — inferable from args (we did this in #1). Confirm
//      that `Map([]int{...}, fn)` works without explicit types.
//   2. Define `Zero[T any]() T` — returns the zero value for T.
//      Show that you MUST call it as `Zero[int]()` — the compiler can't
//      infer T from a return type alone.
//   3. Define `New[T any]() *T` similarly. Confirm you must write `New[int]()`.
//   4. Bonus: pass a result of `Zero[int]()` into a context that expects
//      `int` and observe that *now* inference DOES work because the
//      target type is known.
//
// Run:
//   go run .

package main

import "fmt"

// TODO: Zero, New

func main() {
	fmt.Println("Implement me.")
}
