package main

import "fmt"

type City struct{ Name string }

func (c City) String() string { return "City(" + c.Name + ")" }

func Stringify(x any) string {
	switch v := x.(type) {
	case int:
		return fmt.Sprintf("%d", v)
	case int64:
		return fmt.Sprintf("%d", v)
	case float32:
		return fmt.Sprintf("%.2f", v)
	case float64:
		return fmt.Sprintf("%.2f", v)
	case string:
		return fmt.Sprintf("%q", v)
	case []byte:
		return string(v)
	case fmt.Stringer:
		// `case fmt.Stringer` matches any concrete type with a String() method
		// that wasn't matched by an earlier (more specific) case.
		return v.String()
	default:
		return fmt.Sprintf("?%T?", v)
	}
}

func main() {
	values := []any{
		42, int64(1 << 40), float32(2.5), 3.14159,
		"hello", []byte("bytes"), City{Name: "Tokyo"},
		struct{ X int }{1},
	}
	for _, v := range values {
		fmt.Println(Stringify(v))
	}
}
