package main

import "fmt"

// (1) literal, used locally — small + bounded; might stay on stack.
func smallLocal() int {
	s := []int{1, 2, 3}
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

// (2) large make with constant size — Go often heap-allocates large
// backing arrays even if usage is local.
func largeLocal() int {
	s := make([]int, 1024)
	for i := range s {
		s[i] = i
	}
	return s[1023]
}

// (3) dynamic size — unknown at compile time, so escapes.
func dynamic(n int) int {
	s := make([]int, n)
	for i := range s {
		s[i] = i * 2
	}
	return s[n-1]
}

// (4) returned — always escapes.
func returned() []int {
	return []int{10, 20, 30}
}

func main() {
	fmt.Println(smallLocal(), largeLocal(), dynamic(8), returned())
	// Try: go build -gcflags="-m=2" ./solution
}
