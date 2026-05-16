// Problem 4: Wrapping with %w
//
// `fmt.Errorf("...: %w", err)` produces a NEW error that WRAPS the
// original. `errors.Is` and `errors.As` walk the wrap chain.
//
// Compare to `%v` which just embeds the message text — losing the
// programmatic identity of the wrapped error.
//
// Tasks:
//   1. Define `var ErrTimeout = errors.New("timeout")`.
//   2. `fetchUser` returns `fmt.Errorf("user %d: %w", id, ErrTimeout)`.
//   3. `loadProfile` calls fetchUser and re-wraps: `fmt.Errorf("loadProfile: %w", err)`.
//   4. In main(), confirm that `errors.Is(returnedErr, ErrTimeout)` is
//      true even after two layers of wrapping.
//   5. Re-implement using `%v` instead of `%w`. Show that errors.Is
//      now returns false.
//
// Run:
//   go run .

package main

import "fmt"

// TODO

func main() {
	fmt.Println("Implement me.")
}
