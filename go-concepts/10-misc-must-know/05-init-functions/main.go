// Problem 5: init() Functions
//
// `init()` functions run automatically before main() and can appear in
// any file (you can have multiple per package). Order rules:
//   - Within a single package: vars initialize in declaration order;
//     `init()`s run in source-file order (alphabetical by filename, then
//     order in file).
//   - Across packages: a package's init runs AFTER all its dependencies'
//     inits have completed.
//
// Tasks:
//   1. Define two `init()` functions in this file. Print the order they
//      run. Show that they execute before main().
//   2. Initialize a package-level var with a function call (e.g., loading
//     a regex). Show that the var is ready by the time main() runs.
//
// Run:
//   go run .

package main

import "fmt"

// TODO

func main() {
	fmt.Println("Implement me.")
}
