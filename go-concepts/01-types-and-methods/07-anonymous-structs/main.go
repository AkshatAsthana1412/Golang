// Problem 7: Anonymous Structs
//
// An anonymous struct is a struct type declared inline, without a name.
// They're ideal for table-driven tests, ad-hoc JSON payloads, and one-off
// return values where defining a named type would be ceremony for nothing.
//
// Tasks:
//   1. Build a slice of anonymous-struct test cases:
//        []struct{ in int; want int }{
//          {2, 4}, {3, 9}, {4, 16},
//        }
//      and verify `square(in) == want` for each.
//   2. Define an anonymous-struct LITERAL to send as a JSON request body
//      to a hypothetical endpoint:
//        body := struct {
//          Name  string `json:"name"`
//          Email string `json:"email"`
//        }{ "Ada", "ada@example.com" }
//      Marshal it with encoding/json and print the JSON.
//
// Run:
//   go run .

package main

import (
	"encoding/json"
	"fmt"
)

func square(n int) int { return n * n }

func main() {
	// TODO: table-driven test loop
	items := []struct {
		in   int
		want int
	}{
		{2, 4}, {3, 9}, {4, 16},
	}
	for _, item := range items {
		fmt.Printf("%d^2 == %d? %t\n", item.in, item.want, item.want == item.in*item.in)
	}
	// TODO: anonymous struct + json.Marshal
	body := struct {
		Name    string `json:"name"`
		EmailID string `json:"emailId"`
	}{
		"Ada", "ada.883@gmail.com",
	}
	data, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
