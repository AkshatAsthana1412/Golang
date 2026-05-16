// Problem 2: Path & Query Params
//
// Gin's router supports:
//   - PATH params:  /users/:id          ->  c.Param("id")
//   - WILDCARD:     /files/*path        ->  c.Param("path") (includes leading /)
//   - QUERY:        /search?q=foo       ->  c.Query("q")  /  c.DefaultQuery("q","fallback")
//
// Tasks:
//   1. GET /users/:id           — return {"id": <id>} as JSON.
//   2. GET /files/*path         — return {"path": <path>}.
//   3. GET /search              — read q (required), limit (default 10).
//      Return 400 if q is empty.
//
// Verify:
//   curl http://localhost:8080/users/42
//   curl http://localhost:8080/files/etc/hosts
//   curl http://localhost:8080/search?q=go
//   curl -i http://localhost:8080/search
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
