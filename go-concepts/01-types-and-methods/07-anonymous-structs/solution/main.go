package main

import (
	"encoding/json"
	"fmt"
)

func square(n int) int { return n * n }

func main() {
	cases := []struct {
		in   int
		want int
	}{
		{2, 4}, {3, 9}, {4, 16}, {0, 0}, {-3, 9},
	}
	for _, tc := range cases {
		got := square(tc.in)
		status := "ok"
		if got != tc.want {
			status = "FAIL"
		}
		fmt.Printf("[%s] square(%d) = %d, want %d\n", status, tc.in, got, tc.want)
	}

	body := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{Name: "Ada", Email: "ada@example.com"}

	data, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data)) // {"name":"Ada","email":"ada@example.com"}
}
