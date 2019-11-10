package output_test

import (
	"bytes"
	"errors"
	"github.com/nathanburkett/diff_table/data_transfer"
	"github.com/nathanburkett/diff_table/output"
	"io"
	"reflect"
	"testing"
)

type MockWriter struct {
	nValue int
	err    error
}

func (mw MockWriter) Write(p []byte) (int, error) {
	return mw.nValue, mw.err
}

func TestMarkdownWriter_Write(t *testing.T) {
	type fields struct {
		Writer io.Writer
	}
	type args struct {
		d *data_transfer.Diff
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
				Writer: MockWriter{},
			},
			args: args{
				d: &data_transfer.Diff{
					Insertions: 55,
					Deletions:  45,
					Total:      100,
					Rows: []*data_transfer.DiffRow{
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
					err: errors.New("foo"),
				},
			},
			args: args{
				d: &data_transfer.Diff{
					Insertions: 10,
					Deletions:  20,
					Total:      30,
					Rows: []*data_transfer.DiffRow{
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
				Writer: tt.fields.Writer,
			}
			if err := mw.Write(tt.args.d); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConcreteMarkdownWriter_Write(t *testing.T) {
	type fields struct {
		Writer *bytes.Buffer
	}
	type args struct {
		d *data_transfer.Diff
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
				Writer: &bytes.Buffer{},
			},
			args: args{
				d: &data_transfer.Diff{
					Insertions: 55,
					Deletions:  45,
					Total:      100,
					Rows: []*data_transfer.DiffRow{
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
				Writer: w,
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
		Writer io.Writer
	}
	tests := []struct {
		name string
		args args
		want output.Writer
	}{
		{
			name: "Default",
			args: args{
				Writer: MockWriter{},
			},
			want: &output.MarkdownWriter{
				Writer: MockWriter{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := output.NewMarkdownWriter(tt.args.Writer)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMarkdownWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}
