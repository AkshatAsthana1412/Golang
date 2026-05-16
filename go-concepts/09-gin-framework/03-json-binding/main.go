// Problem 3: JSON Binding
//
// Gin's `c.ShouldBindJSON(&dst)` parses the request body into a struct
// using tagged fields. Returns an error if JSON is malformed or required
// fields are missing.
//
// `c.Bind*` (no Should) auto-aborts and writes a 400 on error — convenient
// for prototypes, less so when you want to control the error shape.
//
// Tasks:
//   1. Define `type CreateUser struct { Name string `+"`"+`json:"name"`+"`"+`; Email string `+"`"+`json:"email"`+"`"+` }`.
//   2. POST /users — bind body. On success, echo with HTTP 201.
//      On invalid JSON, return 400 with the error message.
//
// Verify:
//   curl -i -X POST http://localhost:8080/users \
//     -H "Content-Type: application/json" \
//     -d '{"name":"Ada","email":"ada@example.com"}'
//
//   curl -i -X POST http://localhost:8080/users \
//     -H "Content-Type: application/json" \
//     -d 'not json'
//
// Run:
//   go run .

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	_ = r
	// TODO
}
