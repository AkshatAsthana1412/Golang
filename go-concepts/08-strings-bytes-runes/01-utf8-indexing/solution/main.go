package main

import (
	"fmt"
	"unicode/utf8"
)

func runeAt(s string, i int) rune {
	count := 0
	for _, r := range s {
		if count == i {
			return r
		}
		count++
	}
	return utf8.RuneError
}

func main() {
	s := "héllo世界"
	fmt.Println("len(string)  =", len(s))
	fmt.Println("len([]rune)  =", len([]rune(s)))
	fmt.Println("rune count   =", utf8.RuneCountInString(s))

	// Indexing returns a byte. s[1] is the FIRST byte of "é"'s 2-byte UTF-8.
	fmt.Printf("s[1] as %%d: %d, as %%c: %c\n", s[1], s[1])

	// Decoding properly:
	for i := 0; i < utf8.RuneCountInString(s); i++ {
		fmt.Printf("rune %d: %c\n", i, runeAt(s, i))
	}
}
