package algorithm

import (
	"fmt"
	"strings"
)

var ErrSortTypeNotFound = fmt.Errorf("unknown sort type. only types: %s", strings.Join(Types, ", "))

func Make(t string) (Sorter, error) {
	var (
		s Sorter
		err error
	)

	switch t {
	case TypeTotalDeltaDesc:
		s = NewTotalDeltaDescendingSorter()
	default:
		err = ErrSortTypeNotFound
	}

	return s, err
}
