// Problem 4: Type Switches
//
// A type switch is the idiomatic way to branch on the dynamic type of an
// interface value:
//
//   switch v := x.(type) {
//   case int:    ...
//   case string: ...
//   case []byte, []int: ... // multi-type case: v is still `any` inside
//   default:     ...
//   }
//
// Tasks:
//   Implement `Stringify(x any) string` that returns:
//     - int / int64:        decimal representation
//     - float32 / float64:  with 2 decimals
//     - string:             quoted with %q
//     - []byte:             string conversion
//     - fmt.Stringer:       call x.String()
//     - default:            fmt.Sprintf("?%T?", x)
//
// Run:
//   go run .

package main

import "fmt"

func Stringify(x any) string {
	switch v := x.(type) {
	case int:
		return fmt.Sprintf("%d", v)
	case int64:
		return fmt.Sprintf("%d", v)
	case float32:
		return fmt.Sprintf("%.2f", v)
	case string:
		return fmt.Sprintf("%q", v)
	case []byte:
		return string(v)
	default:
		return fmt.Sprintf("?%T?", v)
	}
}

// TODO: a type that implements fmt.Stringer to test the case

func main() {
	values := []any{
		42, int64(1 << 40), float32(2.5), 3.14159,
		"hello", []byte("bytes"),
		struct{ X int }{1},
	}
	for _, v := range values {
		fmt.Println(Stringify(v))
	}
}
