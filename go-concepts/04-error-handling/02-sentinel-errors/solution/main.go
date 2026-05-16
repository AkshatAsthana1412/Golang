package main

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound   = errors.New("not found")
	ErrPermission = errors.New("permission denied")
)

type Store struct {
	data    map[string]string
	allowed map[string]bool
}

func (s *Store) Get(user, key string) (string, error) {
	if !s.allowed[user] {
		return "", fmt.Errorf("user %q: %w", user, ErrPermission)
	}
	v, ok := s.data[key]
	if !ok {
		return "", fmt.Errorf("key %q: %w", key, ErrNotFound)
	}
	return v, nil
}

func main() {
	s := &Store{
		data:    map[string]string{"greeting": "hello"},
		allowed: map[string]bool{"ada": true},
	}

	for _, c := range []struct{ user, key string }{
		{"ada", "greeting"},
		{"ada", "missing"},
		{"bob", "greeting"},
	} {
		v, err := s.Get(c.user, c.key)
		switch {
		case errors.Is(err, ErrNotFound):
			fmt.Printf("not found: %v\n", err)
		case errors.Is(err, ErrPermission):
			fmt.Printf("forbidden: %v\n", err)
		case err != nil:
			fmt.Printf("other: %v\n", err)
		default:
			fmt.Printf("ok: %s\n", v)
		}
	}
}
