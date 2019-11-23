package input_test

import (
	"bytes"
	"github.com/nathanburkett/diff-chart/datatransfer"
	"github.com/nathanburkett/diff-chart/input"
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
		want    *datatransfer.Diff
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
			want: &datatransfer.Diff{
				Insertions: 45,
				Deletions:  27,
				Total:      72,
				Rows: []*datatransfer.DiffRow{
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
			want:    &datatransfer.Diff{},
			wantErr: true,
		},
		{
			name: "Cannot parse first column to int",
			args: args{
				reader: bytes.NewBuffer([]byte(`twelve      10      foo/bar.go`)),
			},
			want:    &datatransfer.Diff{},
			wantErr: true,
		},
		{
			name: "Cannot parse second column to int",
			args: args{
				reader: bytes.NewBuffer([]byte(`12      ten      foo/bar.go`)),
			},
			want:    &datatransfer.Diff{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cr := input.NewCliDiffNumstatReader(new(datatransfer.Diff))
			got, err := input.Read(cr, tt.args.reader)
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
