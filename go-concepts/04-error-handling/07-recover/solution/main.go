package main

import (
	"errors"
	"fmt"
)

func safeCall(fn func()) (recovered interface{}) {
	defer func() {
		recovered = recover()
	}()
	fn()
	return nil
}

func main() {
	r1 := safeCall(func() { panic("boom") })
	fmt.Printf("r1: %T %v\n", r1, r1)

	r2 := safeCall(func() {
		panic(errors.New("structured"))
	})
	if err, ok := r2.(error); ok {
		fmt.Printf("r2 as error: %v\n", err)
	}

	// Goroutine pitfall — a panic inside `go fn()` is NOT caught by a
	// recover in the parent goroutine. Each goroutine needs its own
	// defer/recover (or wrap the goroutine entry point).
	//   go func() {
	//     defer func() { recover() }()
	//     fn()
	//   }()
}
