package output

import (
	"bytes"
	"fmt"
	"github.com/nathanburkett/diff-chart/datatransfer"
	"io"
)

// TypeMarkdownCli holds flag value for MarkdownWriter type
const TypeMarkdownCli = "github-md"

// Types varying types of output.Writer
var Types = []string{
	TypeMarkdownCli,
}

var (
	markdownTableStart     = []byte("| ")
	markdownTableSeparator = []byte(" | ")
	markdownTableEnd       = []byte(" |\n")
)

// Writer interface that all concrete output writing structs should implement
type Writer interface {
	Write(diff *datatransfer.Diff) error
}

// InternalWriter writer which can write to a data stream and another io.Writer
type InternalWriter interface {
	io.Writer
	io.WriterTo
	Bytes() []byte
}

// MarkdownWriter struct responsible for outputting markdown
type MarkdownWriter struct {
	Writer   io.Writer
	Internal InternalWriter
}

// NewMarkdownWriter factory func for MarkdownWriter
func NewMarkdownWriter(w io.Writer, int InternalWriter) Writer {
	return &MarkdownWriter{
		Writer:   w,
		Internal: int,
	}
}

// Write write diff to select medium
func (mw *MarkdownWriter) Write(d *datatransfer.Diff) error {
	for row := range mw.generateRows(d) {
		if _, err := mw.Internal.Write(row); err != nil {
			return err
		}
	}

	if _, err := mw.Internal.WriteTo(mw.Writer); err != nil {
		return err
	}

	_, err := mw.Writer.Write(mw.Internal.Bytes())

	return err
}

func (mw *MarkdownWriter) generateRows(d *datatransfer.Diff) chan []byte {
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

func (mw *MarkdownWriter) unpackRowToBytes(r *datatransfer.DiffRow, f float64) []byte {
	b := [][]byte{
		r.FullPath,
		[]byte(fmt.Sprintf("+%d/-%d", r.Insertions, r.Deletions)),
		[]byte(fmt.Sprintf("%.2f%%", f)),
	}

	return bytes.Join(b, markdownTableSeparator)
}
