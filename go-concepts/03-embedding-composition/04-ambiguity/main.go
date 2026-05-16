// Problem 4: Promotion Ambiguity & Shadowing
//
// If two embedded types both contribute a method or field with the same
// NAME at the same depth, Go does NOT auto-pick. The selector becomes
// ambiguous unless the outer type SHADOWS it with its own definition.
//
// Tasks:
//   1. Define A with `Name() string` and B with `Name() string`.
//   2. Define `type Outer struct { A; B }`.
//   3. Try calling `Outer{}.Name()` — observe the compile error
//      "ambiguous selector".
//   4. Resolve by either:
//        a. calling `o.A.Name()` / `o.B.Name()` explicitly, OR
//        b. defining `func (o Outer) Name() string { return o.A.Name() }`
//           which shadows the ambiguity.
//
// Run:
//   go run .

package main

import "fmt"

// TODO: A, B, Outer, demonstrate ambiguity + resolution

func main() {
	fmt.Println("Implement me.")
}
