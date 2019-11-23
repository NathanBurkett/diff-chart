package output

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

// ErrWriterTypeNotFound indicates flag and output.Writer pairing not found
var ErrWriterTypeNotFound = fmt.Errorf("unknown writer type. only types: %s", strings.Join(Types, ", "))

// Make returns output.Writer if can be built from flag otherwise returns err
func Make(flag string) (Writer, error) {
	var (
		w Writer
		err error
	)

	switch flag {
	case TypeMarkdownCli:
		w = NewMarkdownWriter(os.Stdout, &bytes.Buffer{})
	default:
		err = ErrWriterTypeNotFound
	}

	return w, err
}
