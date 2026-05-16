// Problem 6: Pointer vs Value Receiver — Cost
//
// "Always use pointer receivers" is a lie. Tradeoffs:
//
//   VALUE receiver:
//     + No indirection at call site
//     + Method body works on a stack-resident copy (often)
//     + Safe to use from concurrent contexts (each caller has a copy)
//     - Copies the receiver on every call (cost = sizeof receiver)
//
//   POINTER receiver:
//     + No copy
//     + Required for mutation
//     - Adds an indirection
//     - Can force the receiver to escape to heap
//
// Tasks:
//   1. Define a `BigStruct` with, say, 16 int64 fields (128 bytes).
//   2. Define a tiny `SmallStruct` with one int.
//   3. Add a value-receiver `(s SmallStruct) Sum() int` and time a tight
//      loop that calls it a million times. (Use time.Now()/Since.)
//   4. Add the same but with BigStruct (value receiver). Note that the
//      gap is real but for tiny structs the value receiver is fine.
//   5. Use `b.Sum()` via a pointer for BigStruct and compare timings.
//
// Run:
//   go run .

package main

import "fmt"

// TODO

func main() {
	fmt.Println("Implement me.")
}
