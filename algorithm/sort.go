package algorithm

import (
	"github.com/nathanburkett/diff_table/data_transfer"
	"sort"
)

const TypeTotalDeltaDesc = "delta"

var Types = []string{
	TypeTotalDeltaDesc,
}

type Sorter interface {
	sort.Interface
	setDiff(*data_transfer.Diff)
}

type TotalDeltaDescendingSorter struct {
	Diff *data_transfer.Diff
}

func NewTotalDeltaDescendingSorter() Sorter {
	return &TotalDeltaDescendingSorter{}
}

func Sort(s Sorter, d *data_transfer.Diff) (*data_transfer.Diff, error) {
	s.setDiff(d)
	sort.Sort(s)

	return d, nil
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

func (dds *TotalDeltaDescendingSorter) setDiff(d *data_transfer.Diff) {
	dds.Diff = d
}
