package output_test

import (
	"bytes"
	"errors"
	"github.com/nathanburkett/diff-chart/datatransfer"
	"github.com/nathanburkett/diff-chart/output"
	"io"
	"reflect"
	"testing"
)

type MockWriter struct {
	NWriteValue int
	WriteErr    error
}

func (mw MockWriter) Write(p []byte) (int, error) {
	return mw.NWriteValue, mw.WriteErr
}

type MockInternalWriter struct {
	NWriteValue   int
	WriteErr      error
	NWriteToValue int64
	WriteToErr    error
	BytesVal      []byte
}

func (miw MockInternalWriter) Write(p []byte) (int, error) {
	return miw.NWriteValue, miw.WriteErr
}

func (miw MockInternalWriter) WriteTo(io.Writer) (n int64, err error) {
	return miw.NWriteToValue, miw.WriteToErr
}

func (miw MockInternalWriter) Bytes() []byte {
	return miw.BytesVal
}

func TestMarkdownWriter_Write(t *testing.T) {
	type fields struct {
		Writer   io.Writer
		Internal output.InternalWriter
	}
	type args struct {
		d *datatransfer.Diff
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Happy path",
			fields: fields{
				Writer:   &bytes.Buffer{},
				Internal: &bytes.Buffer{},
			},
			args: args{
				d: &datatransfer.Diff{
					Insertions: 55,
					Deletions:  45,
					Total:      100,
					Rows: []*datatransfer.DiffRow{
						{
							Insertions: 10,
							Deletions:  20,
							FullPath:   []byte("foo/bar"),
							Segments: [][]byte{
								[]byte("foo"),
								[]byte("bar"),
							},
						},
						{
							Insertions: 30,
							Deletions:  10,
							FullPath:   []byte("foo/baz.md"),
							Segments: [][]byte{
								[]byte("foo"),
								[]byte("baz.md"),
							},
						},
						{
							Insertions: 15,
							Deletions:  15,
							FullPath:   []byte("bar/baz.txt"),
							Segments: [][]byte{
								[]byte("bar"),
								[]byte("baz.txt"),
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Writer produces error",
			fields: fields{
				Writer: MockWriter{
					WriteErr: errors.New("foo"),
				},
				Internal: &bytes.Buffer{},
			},
			args: args{
				d: &datatransfer.Diff{
					Insertions: 10,
					Deletions:  20,
					Total:      30,
					Rows: []*datatransfer.DiffRow{
						{
							Insertions: 10,
							Deletions:  20,
							FullPath:   []byte("foo/bar"),
							Segments: [][]byte{
								[]byte("foo"),
								[]byte("bar"),
							},
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Internal produces error on row generation",
			fields: fields{
				Writer: &bytes.Buffer{},
				Internal: MockInternalWriter{
					NWriteValue:   0,
					WriteErr:      bytes.ErrTooLarge,
					NWriteToValue: 0,
					WriteToErr:    nil,
				},
			},
			args: args{
				d: &datatransfer.Diff{
					Insertions: 10,
					Deletions:  20,
					Total:      30,
					Rows: []*datatransfer.DiffRow{
						{
							Insertions: 10,
							Deletions:  20,
							FullPath:   []byte("foo/bar"),
							Segments: [][]byte{
								[]byte("foo"),
								[]byte("bar"),
							},
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Internal produces error on WriteTo",
			fields: fields{
				Writer: &bytes.Buffer{},
				Internal: MockInternalWriter{
					NWriteValue:   0,
					WriteErr:      nil,
					NWriteToValue: 0,
					WriteToErr:    errors.New("unknown error"),
				},
			},
			args: args{
				d: &datatransfer.Diff{
					Insertions: 10,
					Deletions:  20,
					Total:      30,
					Rows: []*datatransfer.DiffRow{
						{
							Insertions: 10,
							Deletions:  20,
							FullPath:   []byte("foo/bar"),
							Segments: [][]byte{
								[]byte("foo"),
								[]byte("bar"),
							},
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := &output.MarkdownWriter{
				Writer:   tt.fields.Writer,
				Internal: tt.fields.Internal,
			}
			if err := mw.Write(tt.args.d); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConcreteMarkdownWriter_Write(t *testing.T) {
	type fields struct {
		Writer   *bytes.Buffer
		Internal *bytes.Buffer
	}
	type args struct {
		d *datatransfer.Diff
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantErr   bool
		wantBytes []byte
	}{
		{
			name: "Happy path",
			fields: fields{
				Writer:   &bytes.Buffer{},
				Internal: &bytes.Buffer{},
			},
			args: args{
				d: &datatransfer.Diff{
					Insertions: 55,
					Deletions:  45,
					Total:      100,
					Rows: []*datatransfer.DiffRow{
						{
							Insertions: 10,
							Deletions:  20,
							FullPath:   []byte("foo/bar"),
							Segments: [][]byte{
								[]byte("foo"),
								[]byte("bar"),
							},
						},
						{
							Insertions: 30,
							Deletions:  10,
							FullPath:   []byte("foo/baz.md"),
							Segments: [][]byte{
								[]byte("foo"),
								[]byte("baz.md"),
							},
						},
						{
							Insertions: 15,
							Deletions:  15,
							FullPath:   []byte("bar/baz.txt"),
							Segments: [][]byte{
								[]byte("bar"),
								[]byte("baz.txt"),
							},
						},
					},
				},
			},
			wantErr: false,
			wantBytes: []byte(`| Directory | +/- | Δ % |
| --- | --- | --- |
| foo/bar | +10/-20 | 30.00% |
| foo/baz.md | +30/-10 | 40.00% |
| bar/baz.txt | +15/-15 | 30.00% |
`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := tt.fields.Writer
			mw := &output.MarkdownWriter{
				Writer:   w,
				Internal: tt.fields.Internal,
			}
			if err := mw.Write(tt.args.d); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
			if result := w.Bytes(); bytes.Equal(result, tt.wantBytes) == false {
				t.Errorf("Actualy bytes = %v, want bytes %v", result, tt.wantErr)
			}
		})
	}
}

func TestNewMarkdownWriter(t *testing.T) {
	type args struct {
		Writer   io.Writer
		Internal output.InternalWriter
	}
	tests := []struct {
		name string
		args args
		want output.Writer
	}{
		{
			name: "Default",
			args: args{
				Writer:   MockWriter{},
				Internal: MockInternalWriter{},
			},
			want: &output.MarkdownWriter{
				Writer:   MockWriter{},
				Internal: MockInternalWriter{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := output.NewMarkdownWriter(tt.args.Writer, tt.args.Internal)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMarkdownWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}
