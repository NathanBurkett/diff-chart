package data_transfer

import (
	"bytes"
)

var DirSeparator = []byte("/")

type Diff struct {
	Insertions uint64
	Deletions  uint64
	Total      uint64
	Rows       []*DiffRow
}

func NewDiff() *Diff {
	return &Diff{}
}

func (d *Diff) AddRow(row *DiffRow) {
	d.Insertions = d.Insertions + row.Insertions
	d.Deletions = d.Deletions + row.Deletions
	d.Total = d.Insertions + d.Deletions

	d.Rows = append(d.Rows, row)
}

func (d *Diff) GetRowByPath(p []byte) *DiffRow {
	for _, row := range d.Rows {
		if bytes.Compare(row.FullPath, p) == 0 {
			return row
		}
	}

	return nil
}

type DiffRow struct {
	Insertions uint64
	Deletions  uint64
	FullPath   []byte
	Segments   [][]byte
}

func NewDiffRow() *DiffRow {
	return &DiffRow{}
}

func (dr *DiffRow) TotalDelta() uint64 {
	return dr.Insertions + dr.Deletions
}

func (dr *DiffRow) InheritDeltas(frmr DiffRow) {
	dr.Insertions = dr.Insertions + frmr.Insertions
	dr.Deletions = dr.Deletions + frmr.Deletions
}
