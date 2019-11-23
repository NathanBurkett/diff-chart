package algorithm

import (
	"fmt"
	"strings"
)

// ErrSortTypeNotFound indicates flag and algorithm.Sorter pairing not found
var ErrSortTypeNotFound = fmt.Errorf("unknown sort type. only types: %s", strings.Join(Types, ", "))

// Make returns algorithm.Sorter if can be built from flag otherwise returns err
func Make(flag string) (Sorter, error) {
	var (
		s Sorter
		err error
	)

	switch flag {
	case TypeTotalDeltaDesc:
		s = NewTotalDeltaDescendingSorter()
	default:
		err = ErrSortTypeNotFound
	}

	return s, err
}
