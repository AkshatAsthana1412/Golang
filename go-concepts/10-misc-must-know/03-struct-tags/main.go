// Problem 3: Struct Tags & Reflection
//
// Struct tags are string metadata attached to fields, read at runtime
// via the `reflect` package. The standard library's encoding/json reads
// `json:"..."` tags; many other libs follow the convention.
//
// Tasks:
//   1. Define:
//        type User struct {
//          Name  string `json:"name" db:"user_name"`
//          Email string `json:"email,omitempty" db:"email"`
//          age   int    // unexported, no tag — won't be serialized
//        }
//   2. Use `reflect.TypeOf(User{})` to print each field's name, tag.Get("json"),
//      and tag.Get("db").
//   3. Marshal a User with empty Email and verify "email" is omitted.
//
// Run:
//   go run .

package main

import "fmt"

func main() {
	// TODO
	fmt.Println("Implement me.")
}
