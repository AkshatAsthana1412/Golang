// Problem 1: defer Order & Loop Pitfall
//
// Defers run LIFO when the surrounding function returns. Two famous gotchas:
//
//   (a) Arguments are EVALUATED at the defer site, not at execution time.
//   (b) Defer runs at FUNCTION exit — looping `defer f.Close()` inside a
//       for-loop holds every file open until the function returns.
//
// Tasks:
//   1. Print 0..3 and `defer fmt.Println(i)` each. Show that the prints
//      come back as 3,2,1,0 (LIFO) AND each captured the value of `i`
//      at the time of the defer (not the final loop value).
//   2. Same loop but `defer func() { fmt.Println(i) }()` — now i is
//      captured BY REFERENCE in the closure. (Pre-Go 1.22 behavior is
//      that you'd see 4,4,4,4 because the loop var was shared. From Go
//      1.22, each iteration has its own `i`.)
//   3. Show the file-handle-leak pattern (process N filenames opening
//      each in a loop, deferring Close). Discuss in a comment that you
//      should wrap the body in an inner func or use an explicit Close.
//
// Run:
//   go run .

package main

import "fmt"

func main() {
	// TODO
	fmt.Println("Implement me.")
}
