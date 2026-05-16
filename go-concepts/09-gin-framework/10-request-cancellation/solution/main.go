package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/work", func(c *gin.Context) {
		ctx := c.Request.Context()
		for step := 1; step <= 5; step++ {
			select {
			case <-ctx.Done():
				log.Printf("client cancelled at step %d: %v", step, ctx.Err())
				c.AbortWithStatusJSON(499, gin.H{"error": "client cancelled"})
				return
			case <-time.After(1 * time.Second):
				log.Printf("step %d done", step)
			}
		}
		c.JSON(http.StatusOK, gin.H{"steps": 5})
	})

	_ = r.Run(":8080")
}
