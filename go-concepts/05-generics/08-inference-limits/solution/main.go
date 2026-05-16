package main

import "fmt"

func Zero[T any]() T {
	var z T
	return z
}

func New[T any]() *T {
	return new(T)
}

func main() {
	// `Zero()` would not compile — T cannot be inferred from nothing.
	z := Zero[int]()
	fmt.Printf("Zero[int] = %v (%T)\n", z, z)

	p := New[string]()
	*p = "hello"
	fmt.Println("*New[string] =", *p)

	// Inference DOES work where the target type provides T:
	var n int = Zero[int]()
	_ = n
}
