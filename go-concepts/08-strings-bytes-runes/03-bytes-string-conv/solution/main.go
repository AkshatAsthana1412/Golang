package main

import "fmt"

func main() {
	s := "hello"

	b := []byte(s)
	b[0] = 'H'
	fmt.Println("mutated bytes:", string(b)) // Hello
	fmt.Println("original s   :", s)         // hello — strings are immutable

	fmt.Println("len byte ==:", len(s) == len([]byte(s)))

	// Naive concat in a loop (AVOID for hot paths):
	out := ""
	for i := 0; i < 5; i++ {
		out += "x" // creates a new string each iteration
	}
	fmt.Println("naive concat:", out)

	// See next problem: strings.Builder is the efficient way.
}
