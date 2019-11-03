package algorithm

import "github.com/nathanburkett/diff_table/data_transfer"

type TotalDeltaDescendingSorter struct {
	Diff *data_transfer.Diff
}

func (dds TotalDeltaDescendingSorter) Len() int {
	return len(dds.Diff.Rows)
}

func (dds TotalDeltaDescendingSorter) Less(i, j int) bool {
	return dds.Diff.Rows[i].TotalDelta() > dds.Diff.Rows[j].TotalDelta()
}

func (dds TotalDeltaDescendingSorter) Swap(i, j int) {
	dds.Diff.Rows[i], dds.Diff.Rows[j] = dds.Diff.Rows[j], dds.Diff.Rows[i]
}
