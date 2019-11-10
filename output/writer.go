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

type Writer interface {
	Write(diff *data_transfer.Diff) error
}

type MarkdownWriter struct {
	Writer io.Writer
}

func NewMarkdownWriter(w io.Writer) Writer {
	return &MarkdownWriter{
		Writer: w,
	}
}

func (mw *MarkdownWriter) Write(d *data_transfer.Diff) error {
	buf := bytes.NewBuffer([]byte(""))

	for row := range mw.generateRows(d) {
		if _, err := buf.Write(row); err != nil {
			return err
		}
	}

	_, err := mw.Writer.Write(buf.Bytes())

	return err
}

func (mw *MarkdownWriter) generateRows(d *data_transfer.Diff) chan []byte {
	ch := make(chan []byte)

	go func() {
		defer close(ch)

		ch <- []byte("| Directory | +/- | Δ % |\n")
		ch <- []byte("| --- | --- | --- |\n")

		for _, row := range d.Rows {
			percent := float64(row.TotalDelta()) / float64(d.Total) * 100

			buf := bytes.NewBuffer(markdownTableStart)
			buf.Write(mw.unpackRowToBytes(row, percent))
			buf.Write(markdownTableEnd)

			ch <- buf.Bytes()
		}
	}()

	return ch
}

func (mw *MarkdownWriter) unpackRowToBytes(r *data_transfer.DiffRow, f float64) []byte {
	b := [][]byte{
		r.FullPath,
		[]byte(fmt.Sprintf("+%d/-%d", r.Insertions, r.Deletions)),
		[]byte(fmt.Sprintf("%.2f%%", f)),
	}

	return bytes.Join(b, markdownTableSeparator)
}
