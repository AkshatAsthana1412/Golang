package main

import (
	"cmp"
	"fmt"
)

func Min[T cmp.Ordered](xs ...T) T {
	if len(xs) == 0 {
		var zero T
		return zero
	}
	m := xs[0]
	for _, v := range xs[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

func Max[T cmp.Ordered](xs ...T) T {
	if len(xs) == 0 {
		var zero T
		return zero
	}
	m := xs[0]
	for _, v := range xs[1:] {
		if v > m {
			m = v
		}
	}
	return m
}

func Clamp[T cmp.Ordered](v, lo, hi T) T {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

func main() {
	fmt.Println(Min(3, 1, 4, 1, 5, 9, 2)) // 1
	fmt.Println(Max("go", "rust", "c"))    // rust
	fmt.Println(Clamp(150, 0, 100))         // 100
	fmt.Println(Clamp(-1.5, 0.0, 1.0))      // 0
}
