package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrEmptyName    = errors.New("name is empty")
	ErrInvalidEmail = errors.New("email missing @")
	ErrInvalidAge   = errors.New("age must be a positive integer")
)

func ValidateForm(name, email, age string) error {
	var errs []error
	if strings.TrimSpace(name) == "" {
		errs = append(errs, ErrEmptyName)
	}
	if !strings.Contains(email, "@") {
		errs = append(errs, ErrInvalidEmail)
	}
	if n, err := strconv.Atoi(age); err != nil || n <= 0 {
		errs = append(errs, fmt.Errorf("%w (got %q)", ErrInvalidAge, age))
	}
	return errors.Join(errs...)
}

func main() {
	err := ValidateForm("", "no-at", "-3")
	if err != nil {
		fmt.Println("validation failed:")
		fmt.Println(err)
	}

	fmt.Println("---")
	fmt.Println("Is ErrEmptyName?   ", errors.Is(err, ErrEmptyName))
	fmt.Println("Is ErrInvalidEmail?", errors.Is(err, ErrInvalidEmail))
	fmt.Println("Is ErrInvalidAge?  ", errors.Is(err, ErrInvalidAge))

	if err := ValidateForm("Ada", "ada@x", "30"); err == nil {
		fmt.Println("ok form")
	}
}
