// Problem 7: Detecting Escapes
//
// Tooling cheat sheet:
//
//   go build -gcflags="-m"      # one level of escape decisions
//   go build -gcflags="-m=2"    # verbose, with reasons
//   go build -gcflags="-m -l"   # disable inlining (escape under -m can be
//                                 misleading because inlined funcs may not
//                                 escape; -l shows the un-inlined view)
//
//   go test -benchmem ./...     # see allocs/op in benchmarks
//
// Tasks:
//   1. Write a function that you SUSPECT escapes (return a pointer to a
//      local) and one that doesn't (returns a value).
//   2. Run `go build -gcflags="-m" ./solution` from this dir and read the
//      output. Add a comment in your version with what you saw.
//
// Run:
//   go run .
//   go build -gcflags="-m" .
//   go build -gcflags="-m=2" .

package main

import "fmt"

// TODO

func main() {
	fmt.Println("Implement me.")
}
