package transform

import (
	"bytes"
	"github.com/nathanburkett/diff-chart/datatransfer"
)

// TypeDirectoryReducer holds flag value for DirectoryDiffMapReducer
const TypeDirectoryReducer = "dir"

// Types varying types of transform.Reducer
var Types = []string{
	TypeDirectoryReducer,
}

// Reducer interface all concrete map reducing structs should implement
type Reducer interface {
	Reduce(diff *datatransfer.Diff) *datatransfer.Diff
}

// Reduce package level map-reduce handler
func Reduce(r Reducer, d *datatransfer.Diff) (*datatransfer.Diff, error) {
	diff := r.Reduce(d)
	return diff, nil
}

// DirectoryDiffMapReducer reducer which condenses diff rows based upon directories
type DirectoryDiffMapReducer struct {
	Dirs  int
	Split []byte
}

// NewDirectoryDiffMapReducer factory func for DirectoryDiffMapReducer
func NewDirectoryDiffMapReducer(dirs int, split []byte) Reducer {
	return &DirectoryDiffMapReducer{
		Dirs:  dirs,
		Split: split,
	}
}

// Reduce map reduce diff rows by directory
func (dd *DirectoryDiffMapReducer) Reduce(diff *datatransfer.Diff) *datatransfer.Diff {
	o := datatransfer.NewDiff()
	o.Insertions = diff.Insertions
	o.Deletions = diff.Deletions
	o.Total = diff.Total

	for _, frmr := range diff.Rows {
		n := dd.Dirs

		if len(frmr.Segments) < n {
			n = len(frmr.Segments)
		}

		path := bytes.Join(frmr.Segments[0:n], dd.Split)

		trgt := o.GetRowByPath(path)
		if trgt == nil {
			trgt = datatransfer.NewDiffRow()
			trgt.FullPath = path
			trgt.Segments = frmr.Segments[0:n]

			o.AddRow(trgt)
		}

		trgt.InheritDeltas(*frmr)
	}

	return o
}
