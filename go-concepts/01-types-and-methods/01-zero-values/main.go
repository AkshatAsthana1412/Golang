// Problem 1: Zero Values
//
// Every variable in Go has a well-defined zero value when declared without
// an initializer. Predicting them — and knowing which types' zero value is
// USABLE — is bread-and-butter senior knowledge.
//
// Tasks:
//   1. Declare (without initializing) one variable of each type below and
//      print its zero value: int, float64, bool, string, *int, []int,
//      map[string]int, chan int, error, struct{ Name string; Age int }.
//   2. For each, print whether the zero value is "usable" or "panics on use".
//      For map and chan specifically, *try* writing to the nil zero value
//      inside a recover() block so you can OBSERVE the panic.
//
// Why this matters:
//   - A struct's zero value should ideally be useful (sync.Mutex is the
//     canonical example — `var m sync.Mutex` is ready to lock).
//   - A nil map cannot be written to; a nil slice CAN be appended to.
//
// Run:
//   go run .

package main

import "fmt"

func main() {
	// TODO: declare and print zero values for the listed types.
	// TODO: demonstrate the nil-map write panic (caught with recover).
	fmt.Println("Implement me.")
}
