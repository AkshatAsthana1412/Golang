package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Comment struct {
	Body string `json:"body" binding:"required,notblank,max=200"`
}

func main() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("notblank", func(fl validator.FieldLevel) bool {
			return strings.TrimSpace(fl.Field().String()) != ""
		})
	}

	r := gin.Default()

	r.POST("/comments", func(c *gin.Context) {
		var in Comment
		if err := c.ShouldBindJSON(&in); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"ok": true, "comment": in.Body})
	})

	_ = r.Run(":8080")
}
