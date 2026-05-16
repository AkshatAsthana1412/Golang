package main

import "fmt"

func Describe(x any) {
	fmt.Printf("type=%T value=%v\n", x, x)
}

func main() {
	Describe(42)
	Describe("hello")
	Describe(3.14)
	Describe([]int{1, 2})
	Describe(nil)
	Describe(struct{ X int }{7})

	// When to use `any`:
	//   - Logging / fmt-like APIs
	//   - JSON unmarshal targets when shape is unknown
	//   - Heterogeneous containers (map[string]any for config)
	// When to prefer generics:
	//   - Containers/algorithms parameterized by ONE type (Stack[T], Map[K,V])
	//   - When you'd otherwise use type assertions for static safety
}
