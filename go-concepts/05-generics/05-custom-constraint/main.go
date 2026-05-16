// Problem 5: Custom Constraint Interfaces & Type Sets
//
// You can write your own constraints by defining an interface that lists
// allowed types via "type sets":
//
//   type Number interface {
//     ~int | ~int32 | ~int64 | ~float32 | ~float64
//   }
//
// The `~T` form means "any type whose UNDERLYING type is T" — so a custom
// `type Celsius float64` also satisfies the constraint.
//
// Tasks:
//   1. Define a `Number` constraint as above.
//   2. Implement `Sum[T Number](xs []T) T`.
//   3. Define `type Celsius float64` and confirm `Sum([]Celsius{...})`
//      compiles AND returns a Celsius (not a float64).
//
// Run:
//   go run .

package main

import "fmt"

// TODO: Number, Sum, Celsius

func main() {
	fmt.Println("Implement me.")
}
