// Problem 8: Function Types as First-Class Values
//
// Functions in Go are first-class values: you can assign them to variables,
// pass them as arguments, return them, and store them in maps/slices.
// This unlocks callbacks, strategy patterns, and a Go-flavored functional
// style — without needing classes or lambdas-as-objects.
//
// Tasks:
//   1. Define a type alias `type BinOp func(a, b int) int`.
//   2. Build a `map[string]BinOp` containing "+", "-", "*", "/".
//   3. Write `apply(op string, a, b int) (int, error)` that looks up the
//      operator in the map and returns an error if it's missing.
//   4. Demonstrate METHOD VALUES vs METHOD EXPRESSIONS:
//        type Greeter struct{ Name string }
//        func (g Greeter) Hello(s string) string { ... }
//
//        g := Greeter{Name: "Ada"}
//        f1 := g.Hello                    // method VALUE: bound to g
//        f2 := Greeter.Hello              // method EXPRESSION: receiver is 1st arg
//
//        f1("hi")            // -> "Ada says hi"
//        f2(g, "hi")         // -> "Ada says hi"
//
// Run:
//   go run .

package main

import "fmt"

// TODO: type BinOp func(a, b int) int
// TODO: ops map and apply()
// TODO: Greeter type + Hello method

func main() {
	fmt.Println("Implement me.")
}
