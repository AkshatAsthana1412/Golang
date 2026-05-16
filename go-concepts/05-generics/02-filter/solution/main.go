package main

import "fmt"

func Filter[T any](in []T, keep func(T) bool) []T {
	out := make([]T, 0, len(in))
	for _, v := range in {
		if keep(v) {
			out = append(out, v)
		}
	}
	return out
}

func Reduce[T, A any](in []T, init A, f func(A, T) A) A {
	acc := init
	for _, v := range in {
		acc = f(acc, v)
	}
	return acc
}

func main() {
	nums := []int{-2, -1, 0, 1, 2, 3}
	pos := Filter(nums, func(n int) bool { return n > 0 })
	fmt.Println("positives:", pos)

	words := []string{"go", "rust", "python", "c"}
	long := Filter(words, func(s string) bool { return len(s) > 3 })
	fmt.Println("long words:", long)

	sum := Reduce(nums, 0, func(acc, v int) int { return acc + v })
	fmt.Println("sum:", sum)
}
