package main

import "fmt"

func main() {
	mixed := []any{1, "two", 3.0, "four", 5}

	sum := 0
	for _, x := range mixed {
		if n, ok := x.(int); ok {
			sum += n
		}
	}
	fmt.Println("int sum:", sum) // 6

	// Panicking form, caught with recover.
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("recovered: %v\n", r)
			}
		}()
		var any any = "not an int"
		_ = any.(int) // panics: interface conversion
	}()
}
