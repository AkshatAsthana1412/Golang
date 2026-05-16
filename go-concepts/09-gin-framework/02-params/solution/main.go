package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/users/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"id": c.Param("id")})
	})

	r.GET("/files/*path", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"path": c.Param("path")})
	})

	r.GET("/search", func(c *gin.Context) {
		q := c.Query("q")
		if q == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "q is required"})
			return
		}
		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if err != nil || limit < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "limit must be a positive integer"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"q": q, "limit": limit})
	})

	_ = r.Run(":8080")
}
