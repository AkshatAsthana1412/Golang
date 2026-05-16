package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const maxUpload = 5 << 20 // 5 MiB

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = maxUpload

	r.POST("/upload", func(c *gin.Context) {
		fh, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if fh.Size > maxUpload {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "file too large", "max": maxUpload})
			return
		}

		if err := os.MkdirAll("uploads", 0o755); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		dst := filepath.Join("uploads", filepath.Base(fh.Filename))
		if err := c.SaveUploadedFile(fh, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"filename": fh.Filename,
			"size":     fh.Size,
			"path":     dst,
		})
	})

	_ = r.Run(":8080")
}
