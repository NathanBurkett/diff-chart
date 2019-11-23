package datatransfer

import (
	"bytes"
)

var (
	// DirSeparator byte slice representation of directory separator
	DirSeparator = []byte("/")
)

// Diff git diff statistic data transfer object
type Diff struct {
	Insertions uint64
	Deletions  uint64
	Total      uint64
	Rows       []*DiffRow
}

// NewDiff factory func for Diff struct
func NewDiff() *Diff {
	return &Diff{}
}

// AddRow add DiffRow to Diff
func (d *Diff) AddRow(row *DiffRow) {
	d.Insertions = d.Insertions + row.Insertions
	d.Deletions = d.Deletions + row.Deletions
	d.Total = d.Insertions + d.Deletions

	d.Rows = append(d.Rows, row)
}

// GetRowByPath find DiffRow by byte slice path
func (d *Diff) GetRowByPath(p []byte) *DiffRow {
	for _, row := range d.Rows {
		if bytes.Compare(row.FullPath, p) == 0 {
			return row
		}
	}

	return nil
}

// DiffRow representation of diff statistics for given file or path
type DiffRow struct {
	Insertions uint64
	Deletions  uint64
	FullPath   []byte
	Segments   [][]byte
}

// NewDiffRow factory func for DiffRow struct
func NewDiffRow() *DiffRow {
	return &DiffRow{}
}

// TotalDelta calculate total delta for DiffRow
func (dr *DiffRow) TotalDelta() uint64 {
	return dr.Insertions + dr.Deletions
}

// InheritDeltas inherit insertions and deletions from other DiffRow
func (dr *DiffRow) InheritDeltas(frmr DiffRow) {
	dr.Insertions = dr.Insertions + frmr.Insertions
	dr.Deletions = dr.Deletions + frmr.Deletions
}
