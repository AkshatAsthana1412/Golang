package main

import "fmt"

func main() {
	dst := make([]int, 5)
	n := copy(dst, []int{1, 2, 3})
	fmt.Printf("copy(5-dst, 3-src) -> n=%d dst=%v\n", n, dst)

	dst2 := make([]int, 2)
	n = copy(dst2, []int{1, 2, 3, 4, 5})
	fmt.Printf("copy(2-dst, 5-src) -> n=%d dst=%v\n", n, dst2)

	src := []int{10, 20, 30}
	dup := make([]int, len(src))
	copy(dup, src)
	dup[0] = 999
	fmt.Println("src =", src, " (untouched)")
	fmt.Println("dup =", dup)

	b := make([]byte, 8)
	n = copy(b, "hello")
	fmt.Printf("copy(b, \"hello\") -> n=%d b=%q\n", n, b[:n])
}
