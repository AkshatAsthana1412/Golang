// Problem 11: File Upload (multipart)
//
// Gin exposes:
//   c.FormFile("file")    -> *multipart.FileHeader for one file
//   c.MultipartForm()     -> for multiple files
//   c.SaveUploadedFile(fh, "/dst/path")
//
// Tasks:
//   1. POST /upload — accept one form file under field "file".
//   2. Save it under ./uploads/<filename> (create the dir if needed).
//   3. Return JSON with filename and saved size.
//   4. Reject files larger than 5 MiB with 413.
//
// Verify:
//   curl -i -F 'file=@README.md' http://localhost:8080/upload
//
// Run:
//   go run .

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	_ = r
	// TODO
}
