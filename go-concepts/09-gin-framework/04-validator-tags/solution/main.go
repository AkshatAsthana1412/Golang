package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignUp struct {
	Email    string `json:"email"    binding:"required,email"`
	Age      int    `json:"age"      binding:"required,gte=18,lte=120"`
	Password string `json:"password" binding:"required,min=8"`
	Role     string `json:"role"     binding:"required,oneof=user admin"`
}

func main() {
	r := gin.Default()

	r.POST("/signup", func(c *gin.Context) {
		var in SignUp
		if err := c.ShouldBindJSON(&in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Don't echo the password in real life — this is a demo.
		c.JSON(http.StatusOK, gin.H{
			"email": in.Email, "age": in.Age, "role": in.Role,
		})
	})

	_ = r.Run(":8080")
}
