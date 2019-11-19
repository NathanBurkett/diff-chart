package output

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

var ErrWriterTypeNotFound  = fmt.Errorf("unknown writer type. only types: %s", strings.Join(Types, ", "))

func Make(t string) (Writer, error) {
	var (
		w Writer
		err error
	)

	switch t {
	case TypeMarkdownCli:
		w = NewMarkdownWriter(os.Stdout, &bytes.Buffer{})
	default:
		err = ErrWriterTypeNotFound
	}

	return w, err
}
