package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("s len=%d cap=%d %v\n", len(s), cap(s), s)

	t := s[1:4]
	fmt.Printf("t = s[1:4]  len=%d cap=%d %v\n", len(t), cap(t), t)

	t[0] = 999 // mutates the SAME backing array
	fmt.Println("after t[0] = 999, s =", s)

	// Full slice expression caps `u` at len 3, cap 3.
	u := s[:3:3]
	fmt.Printf("u = s[:3:3] len=%d cap=%d %v\n", len(u), cap(u), u)

	u = append(u, 42) // appending past cap -> new backing array
	fmt.Println("after append, u =", u)
	fmt.Println("s untouched   =", s)
}
