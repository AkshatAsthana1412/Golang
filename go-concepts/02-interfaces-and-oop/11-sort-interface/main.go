// Problem 11: sort.Interface
//
// Before generics, the standard way to sort a custom slice type was to
// implement the three-method `sort.Interface`:
//
//   Len() int
//   Less(i, j int) bool
//   Swap(i, j int)
//
// Modern Go ALSO offers `sort.Slice(x, less)` and `slices.SortFunc` that
// don't require an interface. Both styles are still common in the wild;
// know both.
//
// Tasks:
//   Given `type Person struct { Name string; Age int }`:
//   1. Implement `ByAge []Person` with sort.Interface and sort it.
//   2. Re-sort the same slice by Name with `sort.Slice`.
//   3. Re-sort with `slices.SortFunc` using a `cmp.Compare`-style function.
//
// Run:
//   go run .

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// TODO: ByAge type and three methods

func main() {
	fmt.Println("Implement me.")
}
