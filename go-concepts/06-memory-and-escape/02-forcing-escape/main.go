// Problem 2: Forcing an Escape
//
// Easy ways to force a heap allocation (good for tests, bad in production):
//   - return a pointer to a local
//   - assign a local's address into a closure that escapes (returned)
//   - store into a slice/map/channel that escapes
//
// Tasks:
//   1. Write three variants of `makeBuf()`:
//        a. returns []byte by value (escape: yes, slice header points to
//           a backing array; if it's small the array may still be stack
//           if the slice itself doesn't escape — but if returned, it does)
//        b. returns a closure capturing a local []byte
//        c. stores a local []byte into a package-level []*[]byte (escapes)
//   2. Run `go build -gcflags="-m=2" .` and read the output.
//
// Run:
//   go run .

package main

import "fmt"

// TODO

func main() {
	fmt.Println("Implement me.")
}
