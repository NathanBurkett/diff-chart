package transform

import (
	"bytes"
	"github.com/nathanburkett/diff_table/data_transfer"
)

type MapReducer func (diff *data_transfer.Diff, dirLevels int) *data_transfer.Diff

func MapReduceDiffByDirectory(diff *data_transfer.Diff, dirLevels int) *data_transfer.Diff {
	output := data_transfer.NewDiff()

	for _, row := range diff.Rows {
		numSegments := dirLevels

		if len(row.Segments) < dirLevels {
			numSegments = len(row.Segments)
		}

		segments := row.Segments[0:numSegments]

		fullPath := bytes.Join(segments, data_transfer.DirSeparator)

		var diffRow *data_transfer.DiffRow

		if !output.HasRowWithFullPath(fullPath) {
			diffRow = data_transfer.NewDiffRow()
			diffRow.SetPath([]byte(fullPath))
			output.AddRow(diffRow)
		} else {
			diffRow = output.GetRowByFullpath(fullPath)
		}

		diffRow.Insertions = diffRow.Insertions + row.Insertions
		diffRow.Deletions = diffRow.Deletions + row.Deletions
	}

	return output
}
