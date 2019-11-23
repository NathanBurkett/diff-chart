package input

import (
	"fmt"
	"github.com/nathanburkett/diff-chart/datatransfer"
	"strings"
)

// ErrReaderTypeNotFound indicates flag and input.DiffReader pairing not found
var ErrReaderTypeNotFound = fmt.Errorf("unknown reader type. only types: %s", strings.Join(Types, ", "))

// Make returns input.DiffReader if can be built from flag otherwise returns err
func Make(flag string) (DiffReader, error) {
	var (
		dr DiffReader
		err error
	)

	switch flag {
	case TypeGit:
		dr = NewCliDiffNumstatReader(new(datatransfer.Diff))
	default:
		err = ErrReaderTypeNotFound
	}

	return dr, err
}
