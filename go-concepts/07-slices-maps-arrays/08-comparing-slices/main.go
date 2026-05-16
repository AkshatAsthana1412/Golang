// Problem 8: Comparing Slices
//
// Slices do NOT support `==` (except against nil). Comparing element-wise
// you have several options:
//   - manual loop
//   - reflect.DeepEqual (works on anything; slow & no compile-time guard)
//   - slices.Equal (Go 1.21+, generic, fast, comparable element type)
//   - slices.EqualFunc (custom predicate)
//
// Tasks:
//   1. Try `[]int{1,2,3} == []int{1,2,3}`. Note the compile error.
//   2. Compare with reflect.DeepEqual.
//   3. Compare with slices.Equal.
//   4. Use slices.EqualFunc to compare two []string case-insensitively.
//
// Run:
//   go run .

package main

import "fmt"

func main() {
	// TODO
	fmt.Println("Implement me.")
}
