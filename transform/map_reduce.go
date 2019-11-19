package transform

import (
	"bytes"
	"github.com/nathanburkett/diff_table/data_transfer"
)

const TypeDirectoryReducer = "dir"

var Types = []string{
	TypeDirectoryReducer,
}

type Reducer interface {
	Reduce(diff *data_transfer.Diff) *data_transfer.Diff
}

func Reduce(r Reducer, d *data_transfer.Diff) (*data_transfer.Diff, error) {
	diff := r.Reduce(d)
	return diff, nil
}

type DirectoryDiffMapReducer struct {
	Dirs  int
	Split []byte
}

func NewDirectoryDiffMapReducer(dirs int, split []byte) Reducer {
	return &DirectoryDiffMapReducer{
		Dirs:  dirs,
		Split: split,
	}
}

func (dd *DirectoryDiffMapReducer) Reduce(diff *data_transfer.Diff) *data_transfer.Diff {
	o := data_transfer.NewDiff()
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
			trgt = data_transfer.NewDiffRow()
			trgt.FullPath = path
			trgt.Segments = frmr.Segments[0:n]

			o.AddRow(trgt)
		}

		trgt.InheritDeltas(*frmr)
	}

	return o
}
