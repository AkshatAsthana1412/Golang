// Problem 10: io.Reader / io.Writer Pipeline
//
// io.Reader and io.Writer are the two most-implemented interfaces in Go.
// They're tiny — just `Read(p []byte) (int, error)` and
// `Write(p []byte) (int, error)` — but they compose into anything.
//
// Tasks:
//   1. Use `strings.NewReader` to build an io.Reader over "hello world\n".
//   2. Use `io.Copy(os.Stdout, reader)` to stream it to stdout.
//   3. Build a custom io.Writer `UpperWriter` that uppercases every byte
//      it receives before writing to a wrapped io.Writer.
//   4. Pipe a reader through a `tee` (`io.TeeReader`) into both a
//      `bytes.Buffer` and the UpperWriter wrapping os.Stdout.
//
// Why this matters:
//   - io.Reader/Writer is the canonical demonstration of "small interfaces
//     compose into rich behavior."
//
// Run:
//   go run .

package main

import "fmt"

// TODO: type UpperWriter struct { w io.Writer }
// TODO: Write method

func main() {
	fmt.Println("Implement me.")
}
