package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func WriteGreeting(w io.Writer, name string) error {
	_, err := fmt.Fprintf(w, "Hello, %s!\n", name)
	return err
}

type Counter struct{ n int }

func NewCounter() *Counter { return &Counter{} } // returns concrete type

func (c *Counter) Inc()         { c.n++ }
func (c *Counter) Value() int   { return c.n }
func (c *Counter) Reset()       { c.n = 0 }

func main() {
	_ = WriteGreeting(os.Stdout, "Ada")

	var buf bytes.Buffer
	_ = WriteGreeting(&buf, "Bob")
	fmt.Printf("captured: %q\n", buf.String())

	c := NewCounter()
	c.Inc()
	c.Inc()
	c.Inc()
	c.Reset() // available because we returned a concrete type
	c.Inc()
	fmt.Println("counter:", c.Value())
}
