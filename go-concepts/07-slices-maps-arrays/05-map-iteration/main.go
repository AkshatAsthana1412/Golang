// Problem 5: Map Iteration Order
//
// Go intentionally RANDOMIZES the order of `for k, v := range m`. This
// catches code that accidentally relies on iteration order — a portability
// hazard.
//
// Tasks:
//   1. Build `m := map[string]int{"a":1,"b":2,"c":3,"d":4}`.
//   2. Run `range m` three times in a row, printing keys. Observe that the
//      order may differ between runs.
//   3. To get a stable order, copy keys to a slice and sort it:
//        keys := make([]string, 0, len(m))
//        for k := range m { keys = append(keys, k) }
//        sort.Strings(keys)
//      and iterate by keys.
//
// Run:
//   go run .

package main

import "fmt"

func main() {
	// TODO
	fmt.Println("Implement me.")
}
