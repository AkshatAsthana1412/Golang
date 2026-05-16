// Problem 3: Type Assertions
//
// A type assertion extracts the underlying concrete value from an interface:
//   v, ok := x.(T)   // safe: ok=false if x is not a T
//   v := x.(T)       // panics if x is not a T
//
// Tasks:
//   1. Build a `[]any` of mixed values: 1, "two", 3.0, "four".
//   2. Sum only the int values using the comma-ok form. Skip the rest.
//   3. Show the panicking form too — wrap it in recover() and print the
//      panic message.
//
// Run:
//   go run .

package main

import "fmt"

func main() {
	var sum int
	arr := []any{1, "two", 3.0, "four", 5, 6}
	for _, item := range arr {
		num, ok := item.(int)
		if ok {
			sum += num
		}
	}
	fmt.Println("Sum of integer items: ", sum)
}
