// Problem 3: Interface Embedding
//
// Already met briefly in Topic 02 #7. Here, get hands-on:
//
// Tasks:
//   1. Replicate the standard library's:
//        Reader { Read([]byte) (int, error) }
//        Closer { Close() error }
//        ReadCloser interface { Reader; Closer }
//   2. Define `FakeReadCloser{ data []byte; closed bool }` and implement
//      Read + Close so it satisfies ReadCloser.
//   3. In main(), assign it to a `ReadCloser` variable, read all bytes
//      with `io.ReadAll` — actually, since we made our own interface, just
//      loop reading into a 4-byte buffer until io.EOF and Close() at the
//      end.
//
// Run:
//   go run .

package main

import "fmt"

// TODO: Reader, Closer, ReadCloser, FakeReadCloser + methods

func main() {
	fmt.Println("Implement me.")
}
