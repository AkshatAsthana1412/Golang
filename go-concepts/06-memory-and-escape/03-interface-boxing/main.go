// Problem 3: Interface Boxing
//
// When you store a value into an `interface{}` (or `any`), Go creates an
// "interface header" with two words: a TYPE pointer and a DATA pointer.
// Because the data pointer must point to the underlying value, that value
// typically gets COPIED to the heap (boxed).
//
// Tasks:
//   1. Write `printAny(x any)` — anything passed becomes boxed.
//   2. Compare to a generic `printT[T any](x T)` — Go's generics use
//      monomorphization-with-shape-stenciling, and small primitive args
//      can stay on the stack.
//   3. Run with `-gcflags="-m"` and observe.
//   4. Discuss in a comment: this is one reason logging hot paths can
//      surprise you with allocations.
//
// Run:
//   go run .
//   go build -gcflags="-m" .

package main

import "fmt"

// TODO

func main() {
	fmt.Println("Implement me.")
}
