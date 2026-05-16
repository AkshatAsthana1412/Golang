package main

import (
	"errors"
	"fmt"
)

var ErrDB = errors.New("db failure")

type QueryError struct {
	Query string
	Err   error // the wrapped cause
}

func (e *QueryError) Error() string {
	return fmt.Sprintf("query %q: %v", e.Query, e.Err)
}

// Implementing Unwrap lets errors.Is/As walk the chain.
func (e *QueryError) Unwrap() error { return e.Err }

func runQuery(sql string) error {
	return &QueryError{Query: sql, Err: ErrDB}
}

func service() error {
	if err := runQuery("SELECT 1"); err != nil {
		return fmt.Errorf("service: %w", err)
	}
	return nil
}

func main() {
	err := service()
	fmt.Println("err =", err)

	fmt.Println("Is ErrDB?", errors.Is(err, ErrDB))

	var qerr *QueryError
	if errors.As(err, &qerr) {
		fmt.Printf("As *QueryError: Query=%q\n", qerr.Query)
	}

	// Wrong target type: errors.As returns false.
	var p *fmt.Stringer // pointer to a different interface type
	_ = p
}
