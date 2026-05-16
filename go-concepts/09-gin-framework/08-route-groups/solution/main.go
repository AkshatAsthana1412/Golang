package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Query("admin") != "1" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "admin only"})
			return
		}
		c.Next()
	}
}

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"users": []string{"ada", "bob"}})
			})
			users.POST("", func(c *gin.Context) {
				c.JSON(http.StatusCreated, gin.H{"created": true})
			})
		}

		v1.GET("/orders", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"orders": []int{1, 2, 3}})
		})

		admin := v1.Group("/admin", AdminAuth())
		admin.GET("/stats", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"requests_per_sec": 42})
		})
	}

	_ = r.Run(":8080")
}
