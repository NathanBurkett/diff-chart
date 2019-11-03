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

type DiffRow struct {
	Insertions uint64
	Deletions uint64
	FullPath []byte
	Segments [][]byte
}

func NewDiffRow() *DiffRow {
	return &DiffRow{}
}

func (dr *DiffRow) SetPath(path []byte) {
	dr.FullPath = path
	dr.Segments = bytes.Split(path, DirSeparator)
}

func (dr *DiffRow) TotalDelta() uint64 {
	return dr.Insertions + dr.Deletions
}

