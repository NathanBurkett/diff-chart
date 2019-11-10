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
			name: "Successfully read",
			args: args{
				reader: bytes.NewBuffer([]byte(`12      10      foo/bar.go
9       5       foo/baz.go
24      12      bar/baz.go
`)),
			},
			want: &data_transfer.Diff{
				Insertions: 45,
				Deletions:  27,
				Total:      72,
				Rows: []*data_transfer.DiffRow{
					{
						Insertions: 12,
						Deletions:  10,
						FullPath:   []byte("foo/bar.go"),
						Segments: [][]byte{
							[]byte("foo"),
							[]byte("bar.go"),
						},
					},
					{
						Insertions: 9,
						Deletions:  5,
						FullPath:   []byte("foo/baz.go"),
						Segments: [][]byte{
							[]byte("foo"),
							[]byte("baz.go"),
						},
					},
					{
						Insertions: 24,
						Deletions:  12,
						FullPath:   []byte("bar/baz.go"),
						Segments: [][]byte{
							[]byte("bar"),
							[]byte("baz.go"),
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Unknown column in diff numstat",
			args: args{
				reader: bytes.NewBuffer([]byte(`12      10      foo/bar.go     foo`)),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Cannot parse first column to int",
			args: args{
				reader: bytes.NewBuffer([]byte(`twelve      10      foo/bar.go`)),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Cannot parse second column to int",
			args: args{
				reader: bytes.NewBuffer([]byte(`12      ten      foo/bar.go`)),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cr := input.NewCliDiffNumstatReader(new(data_transfer.Diff))
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
