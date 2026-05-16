package main

import (
	"fmt"
	"unicode"
)

type EmptyInputError struct{}

func (EmptyInputError) Error() string { return "empty input" }

func ParseDigits(s string) (int, error) {
	if s == "" {
		return 0, EmptyInputError{}
	}
	sum := 0
	for i, r := range s {
		if !unicode.IsDigit(r) {
			return 0, fmt.Errorf("non-digit %q at index %d", r, i)
		}
		sum += int(r - '0')
	}
	return sum, nil
}

func main() {
	for _, in := range []string{"", "12345", "12a4"} {
		v, err := ParseDigits(in)
		if err != nil {
			fmt.Printf("ParseDigits(%q) -> err: %v\n", in, err)
			continue
		}
		fmt.Printf("ParseDigits(%q) -> %d\n", in, v)
	}
}
