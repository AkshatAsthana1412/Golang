package main

import "fmt"

func main() {
	m := map[int]string{}
	for i := 1; i <= 8; i++ {
		if i%2 == 0 {
			m[i] = "even"
		} else {
			m[i] = "odd"
		}
	}

	for k, v := range m {
		if v == "odd" {
			delete(m, k) // safe to delete during range
		}
	}

	fmt.Println(m) // only even keys remain
}
