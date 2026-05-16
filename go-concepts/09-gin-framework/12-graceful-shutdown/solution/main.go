package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/slow", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			c.AbortWithStatusJSON(499, gin.H{"error": "client gone"})
			return
		case <-time.After(3 * time.Second):
		}
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %v", err)
		}
	}()
	log.Println("server up on :8080")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()
	log.Println("signal received; shutting down (5s grace)")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("shutdown error: %v", err)
	} else {
		log.Println("clean shutdown")
	}
}
