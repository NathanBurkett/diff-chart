package input

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/nathanburkett/diff_table/data_transfer"
	"io"
	"strconv"
)

var (
	UnexpectedDiffColumn = errors.New("unexpected diff report column")
)

type CliDiffNumstatReader struct{}

func NewCliDiffNumstatReader() *CliDiffNumstatReader {
	return &CliDiffNumstatReader{}
}

func (cr *CliDiffNumstatReader) Read(reader io.Reader) (*data_transfer.Diff, error) {
	scanner := bufio.NewScanner(reader)

	output := data_transfer.NewDiff()

	for scanner.Scan() {
		cols := bytes.Fields(scanner.Bytes())
		row := data_transfer.NewDiffRow()

		for i, col := range cols {
			buf := bytes.NewBuffer(col)

			switch i {
			case 0:
				ins, err := cr.parseByteBufferToUint64(buf)
				if err != nil {
					return output, err
				}

				row.Insertions = ins
			case 1:
				del, err := cr.parseByteBufferToUint64(buf)
				if err != nil {
					return output, err
				}

				row.Deletions = del
			case 2:
				row.FullPath = col
				row.Segments = bytes.Split(col, data_transfer.DirSeparator)
			case 3:
				return nil, UnexpectedDiffColumn
			}
		}

		output.AddRow(row)
	}

	return output, nil
}

func (cr *CliDiffNumstatReader) parseByteBufferToUint64(buf *bytes.Buffer) (uint64, error) {
	return strconv.ParseUint(buf.String(), 10, 0)
}
