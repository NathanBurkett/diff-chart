package transform

import (
	"fmt"
	"strings"
)

var ErrReducerTypeNotFound = fmt.Errorf("unknown reducer type. only types: %s", strings.Join(Types, ", "))

func Make(t string) (Reducer, error) {
	var (
		r Reducer
		err error
	)

	switch t {
	case TypeDirectoryReducer:
		r = NewDirectoryDiffMapReducer(1, []byte("/"))
	default:
		err = ErrReducerTypeNotFound
	}

	return r, err
}
