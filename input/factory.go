package input

import (
	"fmt"
	"github.com/nathanburkett/diff_table/data_transfer"
	"strings"
)

var ErrReaderTypeNotFound = fmt.Errorf("unknown reader type. only types: %s", strings.Join(Types, ", "))

func Make(t string) (DiffReader, error) {
	var (
		dr DiffReader
		err error
	)

	switch t {
	case TypeGit:
		dr = NewCliDiffNumstatReader(new(data_transfer.Diff))
	default:
		err = ErrReaderTypeNotFound
	}

	return dr, err
}
