package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	for run := 1; run <= 3; run++ {
		fmt.Printf("run %d (random):  ", run)
		for k := range m {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Print("sorted        :  ")
	for _, k := range keys {
		fmt.Printf("%s=%d ", k, m[k])
	}
	fmt.Println()
}
