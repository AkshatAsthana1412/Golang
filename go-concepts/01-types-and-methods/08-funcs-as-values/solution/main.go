package main

import (
	"errors"
	"fmt"
)

type BinOp func(a, b int) int

var ops = map[string]BinOp{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return a / b },
}

func apply(op string, a, b int) (int, error) {
	fn, ok := ops[op]
	if !ok {
		return 0, fmt.Errorf("unknown op %q", op)
	}
	if op == "/" && b == 0 {
		return 0, errors.New("divide by zero")
	}
	return fn(a, b), nil
}

type Greeter struct{ Name string }

func (g Greeter) Hello(s string) string {
	return fmt.Sprintf("%s says %s", g.Name, s)
}

func main() {
	for _, op := range []string{"+", "-", "*", "/", "%"} {
		r, err := apply(op, 6, 3)
		fmt.Printf("apply(%q, 6, 3) -> %d, err=%v\n", op, r, err)
	}

	g := Greeter{Name: "Ada"}

	// Method VALUE — receiver `g` is captured.
	f1 := g.Hello
	fmt.Println(f1("hi"))

	// Method EXPRESSION — receiver becomes the first parameter.
	f2 := Greeter.Hello
	fmt.Println(f2(g, "hi"))
}
