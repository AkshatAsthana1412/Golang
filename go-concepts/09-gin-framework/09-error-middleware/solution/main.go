package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ErrNotFound = errors.New("not found")

type BadInput struct{ Field string }

func (b *BadInput) Error() string { return "bad input: " + b.Field }

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) == 0 || c.Writer.Written() {
			return
		}
		err := c.Errors.Last().Err

		switch {
		case errors.Is(err, ErrNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			var bad *BadInput
			if errors.As(err, &bad) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "field": bad.Field})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal"})
		}
	}
}

func main() {
	r := gin.Default()
	r.Use(ErrorMiddleware())

	r.GET("/missing", func(c *gin.Context) { _ = c.Error(ErrNotFound) })
	r.GET("/invalid", func(c *gin.Context) { _ = c.Error(&BadInput{Field: "name"}) })
	r.GET("/boom", func(c *gin.Context) { _ = c.Error(errors.New("kaboom")) })

	_ = r.Run(":8080")
}
