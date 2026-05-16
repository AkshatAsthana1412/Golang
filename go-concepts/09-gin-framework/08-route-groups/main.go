// Problem 8: Route Groups
//
// Route groups share a common path prefix and middleware. They keep larger
// servers tidy:
//
//   v1 := r.Group("/api/v1")
//   v1.GET("/users", ...)
//   v1.POST("/users", ...)
//
//   admin := v1.Group("/admin", AdminAuth())   // /api/v1/admin/*
//
// Tasks:
//   1. Build groups for /api/v1/users (GET, POST) and /api/v1/orders
//      (GET).
//   2. Add a nested /api/v1/admin group with its own (stub) AdminAuth
//      middleware that returns 403 unless ?admin=1 is present.
//      Add /api/v1/admin/stats inside.
//
// Verify:
//   curl http://localhost:8080/api/v1/users
//   curl http://localhost:8080/api/v1/orders
//   curl -i http://localhost:8080/api/v1/admin/stats
//   curl -i 'http://localhost:8080/api/v1/admin/stats?admin=1'
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
