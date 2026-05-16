package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		const prefix = "Bearer "
		if !strings.HasPrefix(h, prefix) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
			return
		}
		token := strings.TrimPrefix(h, prefix)
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "empty token"})
			return
		}
		c.Set("userID", "user-"+token)
		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.GET("/public", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	api := r.Group("/api", AuthRequired())
	api.GET("/me", func(c *gin.Context) {
		uid := c.MustGet("userID").(string)
		c.JSON(http.StatusOK, gin.H{"userID": uid})
	})

	_ = r.Run(":8080")
}
