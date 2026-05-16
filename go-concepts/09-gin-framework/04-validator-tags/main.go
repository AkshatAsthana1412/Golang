// Problem 4: Validator Tags
//
// Gin uses go-playground/validator under the hood. You declare validation
// rules right on struct fields:
//
//   type SignUp struct {
//     Email    string `binding:"required,email"`
//     Age      int    `binding:"required,gte=18,lte=120"`
//     Password string `binding:"required,min=8"`
//     Role     string `binding:"required,oneof=user admin"`
//   }
//
// `c.ShouldBindJSON` runs the rules and returns a validator error.
//
// Tasks:
//   1. Define SignUp as above.
//   2. POST /signup — bind, return:
//        - 400 with the validator error on failure
//        - 200 + the validated data on success
//
// Verify:
//   curl -i -X POST http://localhost:8080/signup \
//     -H 'Content-Type: application/json' \
//     -d '{"email":"bad","age":15,"password":"short","role":"god"}'
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
