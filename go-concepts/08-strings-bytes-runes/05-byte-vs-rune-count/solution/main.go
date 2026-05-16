package main

import (
	"fmt"
	"unicode/utf8"
)

func IsAscii(s string) bool {
	for _, r := range s {
		if r >= 128 {
			return false
		}
	}
	return true
}

func main() {
	s := "naïve world 🌍"
	fmt.Println("bytes :", len(s))
	fmt.Println("runes :", utf8.RuneCountInString(s))

	fmt.Println("hello ascii?", IsAscii("hello"))
	fmt.Println("café  ascii?", IsAscii("café"))
}
