// Problem 4: cmp.Ordered Constraint
//
// `cmp.Ordered` (Go 1.21) is the constraint for any type that supports
// `<`, `<=`, `>=`, `>` — that's all integer kinds, all float kinds, and
// strings.
//
// Tasks:
//   1. Implement `Min[T cmp.Ordered](xs ...T) T` and Max similarly.
//   2. Use them with []int and []string.
//   3. Implement `Clamp[T cmp.Ordered](v, lo, hi T) T`.
//
// Run:
//   go run .

package main

import "fmt"

// TODO: Min, Max, Clamp

func main() {
	fmt.Println("Implement me.")
}
