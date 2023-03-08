package ex90

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestPipes(t *testing.T) {
	// add broken to see it fail
	r := ExecPipe(Generate, Uppercase, Indent("  "))

	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	if err != nil {
		t.Fatal("pipefail:", err)
	}

	if v := buf.String(); v != "  A\n  B\n  C\n" {
		t.Error("got:", v)
	}
}

// first command should generate data to the subsequent ones
func ExecPipe(funcs ...piper) io.Reader {
	var last *io.PipeReader
	for _, fn := range funcs {
		r, w := io.Pipe()
		go fn(last, w)
		last = r
	}
	return last
}

type piper func(*io.PipeReader, *io.PipeWriter)

func Generate(_ *io.PipeReader, w *io.PipeWriter) {
	data := []byte(`a
b
c`)
	w.Write(data)
	w.Close()
}

func Indent(v string) piper {
	return func(r *io.PipeReader, w *io.PipeWriter) {
		s := bufio.NewScanner(r)
		for s.Scan() {
			w.Write([]byte(v))
			w.Write(s.Bytes())
			w.Write([]byte("\n"))
		}
		r.CloseWithError(s.Err())
		w.CloseWithError(s.Err())
	}
}

func Uppercase(r *io.PipeReader, w *io.PipeWriter) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		w.Write(bytes.ToUpper(s.Bytes()))
		w.Write([]byte("\n"))
	}
	r.CloseWithError(s.Err())
	w.CloseWithError(s.Err())
}

func broken(r *io.PipeReader, w *io.PipeWriter) {
	err := fmt.Errorf("broken")
	r.CloseWithError(err)
	w.CloseWithError(err)
}
