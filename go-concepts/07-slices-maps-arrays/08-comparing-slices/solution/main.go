package main

import (
	"fmt"
	"reflect"
	"slices"
	"strings"
)

func main() {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}

	// a == b // ❌ compile: slice can only be compared to nil

	fmt.Println("DeepEqual:", reflect.DeepEqual(a, b))
	fmt.Println("slices.Equal:", slices.Equal(a, b))

	x := []string{"Go", "Rust", "C"}
	y := []string{"go", "rust", "c"}
	eq := slices.EqualFunc(x, y, strings.EqualFold)
	fmt.Println("case-insensitive equal:", eq)
}
