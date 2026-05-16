package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestTimer() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next() // run downstream
		// Code here runs AFTER the handler returns.
		log.Printf("%s %s -> %d in %s",
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			time.Since(start).Round(time.Microsecond),
		)
	}
}

func main() {
	r := gin.New()       // start with no defaults
	r.Use(gin.Recovery()) // keep the panic recovery
	r.Use(RequestTimer())

	r.GET("/slow", func(c *gin.Context) {
		time.Sleep(200 * time.Millisecond)
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	_ = r.Run(":8080")
}
