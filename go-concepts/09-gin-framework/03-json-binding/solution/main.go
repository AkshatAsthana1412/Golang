package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	r := gin.Default()

	r.POST("/users", func(c *gin.Context) {
		var in CreateUser
		if err := c.ShouldBindJSON(&in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"created": true,
			"user":    in,
		})
	})

	_ = r.Run(":8080")
}
