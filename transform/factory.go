package transform

import (
	"fmt"
	"strings"
)

// ErrReducerTypeNotFound indicates flag and transform.Reducers pairing not found
var ErrReducerTypeNotFound = fmt.Errorf("unknown reducer type. only types: %s", strings.Join(Types, ", "))

// Make returns transform.Reducer if can be built from flag otherwise returns err
func Make(flag string) (Reducer, error) {
	var (
		r Reducer
		err error
	)

	switch flag {
	case TypeDirectoryReducer:
		r = NewDirectoryDiffMapReducer(1, []byte("/"))
	default:
		err = ErrReducerTypeNotFound
	}

	return r, err
}
