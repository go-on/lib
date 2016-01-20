package mustache

import (
	"fmt"
	"io"
)

type Replacer interface {
	Get(string) interface{} // might be []string, []interface, string, bool, int, Stringer, map[string]interface
}

type Template struct {
	Open         []rune
	Close        []rune
	BytesWritten int
	BytesRead    int
	err          error
	src          io.Reader
	dest         io.Writer
	repl         Replacer
	buffsize     int
	buff         []byte
	placeholder  string
	replaceNow   bool
}

func (t *Template) Merge() (err error) {
	// var err error
	var n int
	for err == nil {
		n, err = t.src.Read(t.buff)
		t.BytesRead += n
		errWrite := t.Write(t.buff[:n])
		if errWrite != nil {
			return errWrite
		}

	}

	if err != io.EOF {
		return err
	}

	return nil
}

func (t *Template) Write(bf []byte) error {
	rest := t.Replace(bf)
	if len(rest) > 0 {
		n, err := t.dest.Write(rest)
		t.BytesWritten += n
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *Template) Replace(in []byte) (rest []byte) {
	// todo: check if we have a current placeholder and if so call the replacer

	if t.replaceNow && t.placeholder != "" {
		replacement := t.repl.Get(t.placeholder)

		switch v := replacement.(type) {
		case string:
			n, err := t.dest.Write([]byte(v))
			t.BytesWritten += n
			if err != nil {
				t.err = err
			}
		case []byte:
			n, err := t.dest.Write(v)
			t.BytesWritten += n
			if err != nil {
				t.err = err
			}
		case fmt.Stringer:
			n, err := t.dest.Write([]byte(v.String()))
			t.BytesWritten += n
			if err != nil {
				t.err = err
			}
		case bool: // todo: skip section for false and show it for true
		case []interface{}: // todo: iterate
		case map[string]interface{}:
		default:
		}

	}
	// TODO return the rest
	return
}

func New(src io.Reader, dest io.Writer, repl Replacer) *Template {
	return &Template{
		Open:     []rune{'{', '{'},
		Close:    []rune{'}', '}'},
		buffsize: 1024,
		src:      src,
		dest:     dest,
		buff:     make([]byte, 1024),
	}
}

func Merge(src io.Reader, dest io.Writer, repl Replacer) (written int, err error) {
	return
}
