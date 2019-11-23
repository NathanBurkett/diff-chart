package algorithm

import (
	"github.com/nathanburkett/diff-chart/datatransfer"
	"sort"
)

// TypeTotalDeltaDesc holds flag value for TotalDeltaDescendingSorter type
const TypeTotalDeltaDesc = "delta"

// Types varying types of algorithm.Sorter instances that can be built
var Types = []string{
	TypeTotalDeltaDesc,
}

// Sorter interface that all concrete sorting algorithm structs should implement
type Sorter interface {
	sort.Interface
	setDiff(*datatransfer.Diff)
}

// TotalDeltaDescendingSorter algo struct which prioritizes directories by largest diff
type TotalDeltaDescendingSorter struct {
	Diff *datatransfer.Diff
}

// NewTotalDeltaDescendingSorter factory for TotalDeltaDescendingSorter
func NewTotalDeltaDescendingSorter() Sorter {
	return &TotalDeltaDescendingSorter{}
}

// Sort package level algorithmic sorting handler
func Sort(s Sorter, d *datatransfer.Diff) (*datatransfer.Diff, error) {
	s.setDiff(d)
	sort.Sort(s)

	return d, nil
}

// Len is the number of elements in the TotalDeltaDescendingSorter collection.
func (dds TotalDeltaDescendingSorter) Len() int {
	return len(dds.Diff.Rows)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (dds TotalDeltaDescendingSorter) Less(i, j int) bool {
	return dds.Diff.Rows[i].TotalDelta() > dds.Diff.Rows[j].TotalDelta()
}

// Swap swaps the elements with indexes i and j.
func (dds TotalDeltaDescendingSorter) Swap(i, j int) {
	dds.Diff.Rows[i], dds.Diff.Rows[j] = dds.Diff.Rows[j], dds.Diff.Rows[i]
}

func (dds *TotalDeltaDescendingSorter) setDiff(d *datatransfer.Diff) {
	dds.Diff = d
}
