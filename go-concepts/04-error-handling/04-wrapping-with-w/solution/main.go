package main

import (
	"errors"
	"fmt"
)

var ErrTimeout = errors.New("timeout")

func fetchUser(id int) error {
	return fmt.Errorf("user %d: %w", id, ErrTimeout)
}

func loadProfile(id int) error {
	if err := fetchUser(id); err != nil {
		return fmt.Errorf("loadProfile: %w", err)
	}
	return nil
}

// Anti-version using %v — does NOT preserve identity.
func fetchUserBuggy(id int) error {
	return fmt.Errorf("user %d: %v", id, ErrTimeout)
}

func main() {
	err := loadProfile(42)
	fmt.Println("err =", err)
	fmt.Println("is ErrTimeout?", errors.Is(err, ErrTimeout))

	bad := fetchUserBuggy(99)
	fmt.Println("buggy err =", bad)
	fmt.Println("is ErrTimeout (with non-wrap verb)?", errors.Is(bad, ErrTimeout))
}
