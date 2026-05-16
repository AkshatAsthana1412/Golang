package main

import (
	"fmt"
	"regexp"
)

var emailRe = regexp.MustCompile(`^[^@]+@[^@]+\.[^@]+$`)

func init() {
	fmt.Println("init #1")
}

func init() {
	fmt.Println("init #2")
}

func main() {
	fmt.Println("main starts")
	fmt.Println("emailRe ready:", emailRe.MatchString("ada@example.com"))
}
