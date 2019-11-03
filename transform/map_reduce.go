package transform

import (
	"bytes"
	"github.com/nathanburkett/diff_table/data_transfer"
)

type MapReducer func (diff *data_transfer.Diff, dirLevels int) *data_transfer.Diff

func MapReduceDiffByDirectory(diff *data_transfer.Diff, dirLevels int) *data_transfer.Diff {
	output := data_transfer.NewDiff()

	container := map[string][]*data_transfer.DiffRow{}

	for _, row := range diff.Rows {
		numSegments := dirLevels

		if len(row.Segments) < dirLevels {
			numSegments = len(row.Segments)
		}

		segments := row.Segments[0:numSegments]

		fullPath := bytes.Join(segments, data_transfer.DirSeparator)
		fullPathString := string(fullPath)

		container[fullPathString] = append(container[fullPathString], row)
	}

	for fullPath, rows := range container {
		diffRow := data_transfer.NewDiffRow()
		diffRow.SetPath([]byte(fullPath))

		for _, row := range rows {
			diffRow.Insertions = diffRow.Insertions + row.Insertions
			diffRow.Deletions = diffRow.Deletions + row.Deletions
		}

		output.AddRow(diffRow)
	}

	return output
}
