package data_transfer

import (
	"bytes"
)

var DirSeparator = []byte("/")

type Diff struct {
	Rows []*DiffRow
}

func NewDiff() *Diff {
	return &Diff{}
}

func (d *Diff) AddRow(row *DiffRow) {
	d.Rows = append(d.Rows, row)
}

func (d *Diff) HasRowWithFullPath(FullPath []byte) bool {
	for _, row := range d.Rows {
		if bytes.Compare(row.FullPath, FullPath) == 0 {
			return true
		}
	}

	return false
}

func (d *Diff) GetRowByFullpath(FullPath []byte) *DiffRow {
	for _, row := range d.Rows {
		if bytes.Compare(row.FullPath, FullPath) == 0 {
			return row
		}
	}

	return nil
}

type DiffRow struct {
	Insertions uint64
	Deletions uint64
	FullPath []byte
	Segments [][]byte
}

func NewDiffRow() *DiffRow {
	return &DiffRow{}
}

func (dr *DiffRow) TotalDelta() uint64 {
	return dr.Insertions + dr.Deletions
}

