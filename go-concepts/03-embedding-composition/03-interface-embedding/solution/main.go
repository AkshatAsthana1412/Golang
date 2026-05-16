package main

import (
	"errors"
	"fmt"
	"io"
)

type Reader interface {
	Read(p []byte) (int, error)
}

type Closer interface {
	Close() error
}

type ReadCloser interface {
	Reader
	Closer
}

type FakeReadCloser struct {
	data   []byte
	pos    int
	closed bool
}

func (f *FakeReadCloser) Read(p []byte) (int, error) {
	if f.closed {
		return 0, errors.New("read on closed reader")
	}
	if f.pos >= len(f.data) {
		return 0, io.EOF
	}
	n := copy(p, f.data[f.pos:])
	f.pos += n
	return n, nil
}

func (f *FakeReadCloser) Close() error {
	f.closed = true
	return nil
}

func main() {
	var rc ReadCloser = &FakeReadCloser{data: []byte("hello, world!")}

	buf := make([]byte, 4)
	for {
		n, err := rc.Read(buf)
		if n > 0 {
			fmt.Printf("read %d: %q\n", n, buf[:n])
		}
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Println("error:", err)
			break
		}
	}
	_ = rc.Close()
}
