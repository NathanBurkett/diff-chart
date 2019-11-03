package output

import (
	"bytes"
	"fmt"
	"github.com/nathanburkett/diff_table/data_transfer"
	"io"
)

var (
	markdownTableStart     = []byte("| ")
	markdownTableSeparator = []byte(" | ")
	markdownTableEnd       = []byte(" |\n")
)

func Write(writer io.Writer, diff *data_transfer.Diff) error {
	var (
		grandTotal     uint64
		byteSlice []byte
	)

	buf := bytes.NewBuffer(byteSlice)
	buf.Write([]byte("| Directory | +/- | Δ % |\n"))
	buf.Write([]byte("| --- | --- | --- |\n"))

	for _, row := range diff.Rows {
		grandTotal = grandTotal + row.TotalDelta()
	}

	for _, row := range diff.Rows {
		percent := float64(row.TotalDelta()) / float64(grandTotal) * 100
		if err := writeDiffRow(buf, row, percent); err != nil {
			return err
		}
	}

	_, err := writer.Write(buf.Bytes())

	return err
}

func writeDiffRow(buf *bytes.Buffer, row *data_transfer.DiffRow, percent float64) error {
	rowBytes := [][]byte{
		markdownTableStart,
		row.FullPath,
		markdownTableSeparator,
		[]byte(fmt.Sprintf("+%d/-%d", row.Insertions, row.Deletions)),
		markdownTableSeparator,
		[]byte(fmt.Sprintf("%.2f%%", percent)),
		markdownTableEnd,
	}

	for _, byteRow := range rowBytes {
		if _, err := buf.Write(byteRow); err != nil {
			return err
		}
	}

	return nil
}
