package transform

import (
	"bytes"
	"github.com/nathanburkett/diff_table/data_transfer"
)

type MapReducer func (diff *data_transfer.Diff, dirLevels int) *data_transfer.Diff

func MapReduceDiffByDirectory(diff *data_transfer.Diff, dirs int, split []byte) *data_transfer.Diff {
	output := data_transfer.NewDiff()

	for _, frmr := range diff.Rows {
		num := dirs

		if len(frmr.Segments) < dirs {
			num = len(frmr.Segments)
		}

		path := bytes.Join(frmr.Segments[0:num], split)

		trgt := output.GetRowByPath(path)
		if trgt == nil {
			trgt = data_transfer.NewDiffRow()
			trgt.FullPath = path
			trgt.Segments = frmr.Segments[0:num]

			output.AddRow(trgt)
		}

		trgt.InheritDeltas(*frmr)
	}

	return output
}
