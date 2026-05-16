package main

import (
	"errors"
	"fmt"
	"strings"
)

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation: %s: %s", e.Field, e.Message)
}

func ValidateUser(name, email string) error {
	if strings.TrimSpace(name) == "" {
		return &ValidationError{Field: "name", Message: "must not be empty"}
	}
	if !strings.Contains(email, "@") {
		return &ValidationError{Field: "email", Message: "missing @"}
	}
	return nil
}

func main() {
	cases := []struct{ name, email string }{
		{"", "ada@example.com"},
		{"Bob", "bob_at_example.com"},
		{"Cara", "cara@example.com"},
	}
	for _, c := range cases {
		err := ValidateUser(c.name, c.email)
		if err == nil {
			fmt.Printf("ok: %+v\n", c)
			continue
		}
		var vErr *ValidationError
		if errors.As(err, &vErr) {
			fmt.Printf("INVALID field=%s msg=%s\n", vErr.Field, vErr.Message)
		} else {
			fmt.Println("other error:", err)
		}
	}
}
