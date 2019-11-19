package input

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/nathanburkett/diff_table/data_transfer"
	"io"
	"strconv"
)

const TypeGit = "cli"

var Types = []string{
	TypeGit,
}

type DiffReader interface {
	Generate(scn *bufio.Scanner) chan [][]byte
	TransformToDiffRow(fields [][]byte) (*data_transfer.DiffRow, error)
}

var (
	UnexpectedCliDiffColumn = errors.New("unexpected diff report column")
)

type CliDiffNumstatReader struct {
	Output *data_transfer.Diff
}

func NewCliDiffNumstatReader(diff *data_transfer.Diff) *CliDiffNumstatReader {
	return &CliDiffNumstatReader{
		Output: diff,
	}
}

func Read(drdr DiffReader, rdr io.Reader) (*data_transfer.Diff, error) {
	diff := new(data_transfer.Diff)
	scn := bufio.NewScanner(rdr)

	for cols := range drdr.Generate(scn) {
		row, err := drdr.TransformToDiffRow(cols)
		if err != nil {
			return diff, err
		}

		diff.AddRow(row)
	}

	return diff, nil
}

func (cr *CliDiffNumstatReader) Generate(scanner *bufio.Scanner) chan [][]byte {
	ch := make(chan [][]byte)

	go func() {
		defer close(ch)
		for scanner.Scan() {
			ch <- bytes.Fields(scanner.Bytes())
		}
	}()

	return ch
}

func (cr *CliDiffNumstatReader) TransformToDiffRow(cols [][]byte) (*data_transfer.DiffRow, error) {
	row := data_transfer.NewDiffRow()

	for i, col := range cols {
		switch i {
		case 0:
			if err := cr.coerceBytesToUint64(col, &row.Insertions); err != nil {
				return nil, err
			}
		case 1:
			if err := cr.coerceBytesToUint64(col, &row.Deletions); err != nil {
				return nil, err
			}
		case 2:
			row.FullPath = col
			row.Segments = bytes.Split(col, data_transfer.DirSeparator)
		default:
			return nil, UnexpectedCliDiffColumn
		}
	}

	return row, nil
}

func (cr *CliDiffNumstatReader) coerceBytesToUint64(b []byte, trgt *uint64) error {
	buf := bytes.NewBuffer(b)

	// todo - remove conversion to string then uint (endian byte order)
	val, err := strconv.ParseUint(buf.String(), 10, 0)
	if err != nil {
		return err
	}

	*trgt = val

	return nil
}
