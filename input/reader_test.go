package input_test

import (
	"bytes"
	"github.com/nathanburkett/diff_table/data_transfer"
	"github.com/nathanburkett/diff_table/input"
	"io"
	"reflect"
	"testing"
)

func TestCliDiffNumstatReader_Read(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *data_transfer.Diff
		wantErr bool
	}{
		{
			name:    "Successfully read",
			args:    args{
				reader: bytes.NewBuffer([]byte(`12      10      foo/bar.go
9       5       foo/baz.go
24      12      bar/baz.go
`)),
			},
			want: &data_transfer.Diff{
				Rows: []*data_transfer.DiffRow{
					{
						Insertions: 12,
						Deletions:  10,
						FullPath:   []byte("foo/bar.go"),
						Segments:   [][]byte{
							[]byte("foo"),
							[]byte("bar.go"),
						},
					},
					{
						Insertions: 9,
						Deletions:  5,
						FullPath:   []byte("foo/baz.go"),
						Segments:   [][]byte{
							[]byte("foo"),
							[]byte("baz.go"),
						},
					},
					{
						Insertions: 24,
						Deletions:  12,
						FullPath:   []byte("bar/baz.go"),
						Segments:   [][]byte{
							[]byte("bar"),
							[]byte("baz.go"),
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name:    "Unknown column in diff numstat",
			args:    args{
				reader: bytes.NewBuffer([]byte(`12      10      foo/bar.go     foo`)),
			},
			want: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cr := &input.CliDiffNumstatReader{}
			got, err := cr.Read(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCliDiffNumstatReader(t *testing.T) {
	tests := []struct {
		name string
		want *input.CliDiffNumstatReader
	}{
		{
			name: "*CliDiffNumstatReader successfully created",
			want: &input.CliDiffNumstatReader{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := input.NewCliDiffNumstatReader(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCliDiffNumstatReader() = %v, want %v", got, tt.want)
			}
		})
	}
}
