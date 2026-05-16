package main

import "fmt"

func Map[T, U any](in []T, f func(T) U) []U {
	out := make([]U, len(in))
	for i, v := range in {
		out[i] = f(v)
	}
	return out
}

func main() {
	squared := Map([]int{1, 2, 3}, func(n int) int { return n * n })
	fmt.Println(squared) // [1 4 9]

	lengths := Map([]string{"go", "is", "fast"}, func(s string) int { return len(s) })
	fmt.Println(lengths) // [2 2 4]

	// Type inference picks T from the slice and U from the func return.
	// You can be explicit when needed: Map[string, int](...)
}
