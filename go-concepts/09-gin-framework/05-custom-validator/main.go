// Problem 5: Custom Validator
//
// You can register custom validation rules with the underlying validator:
//
//   import "github.com/go-playground/validator/v10"
//   v, _ := binding.Validator.Engine().(*validator.Validate)
//   v.RegisterValidation("notblank", func(fl validator.FieldLevel) bool {
//     return strings.TrimSpace(fl.Field().String()) != ""
//   })
//
// Tasks:
//   1. Register a `notblank` validator that rejects empty/whitespace strings.
//   2. Define `type Comment struct { Body string `+"`"+`binding:"required,notblank,max=200"`+"`"+` }`.
//   3. POST /comments — bind. Return 400 on failure, 200 on success.
//
// Verify:
//   curl -i -X POST http://localhost:8080/comments -H 'Content-Type: application/json' -d '{"body":"   "}'
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
