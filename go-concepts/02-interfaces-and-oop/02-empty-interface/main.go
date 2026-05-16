// Problem 2: The Empty Interface (`any`)
//
// `interface{}` (aliased as `any` since Go 1.18) is the empty interface —
// every type satisfies it. Use sparingly; you lose static guarantees.
//
// Tasks:
//   1. Write a function `func Describe(x any)` that prints the value AND
//      its concrete type using `%T` and `%v`.
//   2. Call it with: 42, "hello", 3.14, []int{1,2}, nil, struct{X int}{7}.
//   3. Discuss (in a comment) when `any` is appropriate vs when generics
//      (Topic 05) are a better fit.
//
// Run:
//   go run .

package main

import "fmt"

// TODO: Describe

func main() {
	_ = fmt.Sprint("warm fmt")
	fmt.Println("Implement me.")
}
