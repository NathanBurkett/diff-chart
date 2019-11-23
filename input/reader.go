package input

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/nathanburkett/diff-chart/datatransfer"
	"io"
	"strconv"
)

// TypeGit holds flag for CliDiffNumstatReader
const TypeGit = "cli"

// Types varying types of input.DiffReader instances that can be built
var Types = []string{
	TypeGit,
}

// DiffReader interface all concrete input diff readers should implement
type DiffReader interface {
	Generate(scn *bufio.Scanner) chan [][]byte
	TransformToDiffRow(fields [][]byte) (*datatransfer.DiffRow, error)
}

var (
	// ErrUnexpectedCliDiffColumn indicates an unexpected cli diff column
	ErrUnexpectedCliDiffColumn = errors.New("unexpected diff report column")
)

// CliDiffNumstatReader reader which consumes cli diff command output
type CliDiffNumstatReader struct {
	Output *datatransfer.Diff
}

// NewCliDiffNumstatReader factory func for CliDiffNumstatReader
func NewCliDiffNumstatReader(diff *datatransfer.Diff) DiffReader {
	return &CliDiffNumstatReader{
		Output: diff,
	}
}

// Read package level diff reading handler
func Read(drdr DiffReader, rdr io.Reader) (*datatransfer.Diff, error) {
	diff := new(datatransfer.Diff)
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

// Generate generates slice of byte slices which correspond to columnar
// cli git diff report values
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

// TransformToDiffRow transforms slice of byte slices to DiffRow struct
func (cr *CliDiffNumstatReader) TransformToDiffRow(cols [][]byte) (*datatransfer.DiffRow, error) {
	row := datatransfer.NewDiffRow()

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
			row.Segments = bytes.Split(col, datatransfer.DirSeparator)
		default:
			return nil, ErrUnexpectedCliDiffColumn
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
