package main

import "fmt"

func main() {
	parent := make([]int, 0, 10)
	parent = append(parent, 1, 2, 3) // len=3, cap=10

	childA := append(parent, 100) // len=4 — same backing array
	childB := append(parent, 200) // ALSO len=4, OVERWROTE childA[3]

	fmt.Println("parent =", parent)
	fmt.Println("childA =", childA, " <- expected [1 2 3 100]")
	fmt.Println("childB =", childB)

	// Fix #1: full-slice form caps the parent's view at its current length,
	// forcing the next append to grow.
	p2 := append(parent[:len(parent):len(parent)], 100)
	q2 := append(parent[:len(parent):len(parent)], 200)
	fmt.Println("\nWith full-slice fix:")
	fmt.Println("p2 =", p2)
	fmt.Println("q2 =", q2)

	// Fix #2: copy first.
	cp := make([]int, len(parent))
	copy(cp, parent)
	p3 := append(cp, 100)
	fmt.Println("p3 =", p3, "(parent untouched)")
}
