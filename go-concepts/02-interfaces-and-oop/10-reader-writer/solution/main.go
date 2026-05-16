package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

type UpperWriter struct{ w io.Writer }

func (u UpperWriter) Write(p []byte) (int, error) {
	out := make([]byte, len(p))
	for i, b := range p {
		out[i] = byte(unicode.ToUpper(rune(b)))
	}
	// We must report bytes from the ORIGINAL p, not from `out`, per the
	// io.Writer contract: "Write returns the number of bytes written".
	_, err := u.w.Write(out)
	return len(p), err
}

func main() {
	// 1+2: strings.NewReader -> io.Copy -> stdout
	r := strings.NewReader("hello world\n")
	_, _ = io.Copy(os.Stdout, r)

	// 3: UpperWriter
	upper := UpperWriter{w: os.Stdout}
	_, _ = io.Copy(upper, strings.NewReader("uppercased\n"))

	// 4: TeeReader — every byte read goes through the tee to a buffer.
	src := strings.NewReader("piped through tee\n")
	var captured bytes.Buffer
	tee := io.TeeReader(src, &captured)
	_, _ = io.Copy(upper, tee)
	fmt.Printf("tee captured: %q\n", captured.String())
}
