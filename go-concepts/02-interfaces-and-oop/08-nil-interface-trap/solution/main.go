package main

import "fmt"

type MyError struct{ Msg string }

func (e *MyError) Error() string { return e.Msg }

// BUG: returns a TYPED nil. To callers, the returned `error` interface
// has type=*MyError, value=nil — and `err != nil` evaluates to TRUE.
func DoStuffBuggy() error {
	var e *MyError // nil *MyError
	// ... no error happened ...
	return e
}

// FIX: return the bare nil literal. The returned interface is the all-nil
// interface, and `err != nil` is false as expected.
func DoStuffFixed() error {
	var e *MyError
	_ = e
	return nil
}

func main() {
	if err := DoStuffBuggy(); err != nil {
		fmt.Printf("BUG: typed-nil read as error (type=%T, value=%v)\n", err, err)
	}
	if err := DoStuffFixed(); err == nil {
		fmt.Println("FIXED: nil error correctly nil")
	}

	// Workaround if you must keep the typed-nil pattern: explicitly check.
	check := func(err error) {
		if err == nil {
			return
		}
		// One pragmatic guard:
		if me, ok := err.(*MyError); ok && me == nil {
			return
		}
		fmt.Println("real error:", err)
	}
	check(DoStuffBuggy())
}
