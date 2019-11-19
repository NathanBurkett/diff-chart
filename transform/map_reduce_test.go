package transform_test

import (
	"github.com/nathanburkett/diff_table/data_transfer"
	"github.com/nathanburkett/diff_table/transform"
	"reflect"
	"testing"
)

func TestDirectoryDiffMapReducer_Reduce(t *testing.T) {
	diffs := &data_transfer.Diff{
		Insertions: 55,
		Deletions:  110,
		Total:      165,
		Rows: []*data_transfer.DiffRow{
			{
				Insertions: 10,
				Deletions:  20,
				FullPath:   []byte("foo/bar/baz.go"),
				Segments: [][]byte{
					[]byte("foo"),
					[]byte("bar"),
					[]byte("baz.go"),
				},
			},
			{
				Insertions: 10,
				Deletions:  20,
				FullPath:   []byte("foo/baz.go"),
				Segments: [][]byte{
					[]byte("foo"),
					[]byte("baz.go"),
				},
			},
			{
				Insertions: 10,
				Deletions:  20,
				FullPath:   []byte("foo/bar/foo.go"),
				Segments: [][]byte{
					[]byte("foo"),
					[]byte("bar"),
					[]byte("foo.go"),
				},
			},
			{
				Insertions: 10,
				Deletions:  20,
				FullPath:   []byte("foo/foo/foo/baz.go"),
				Segments: [][]byte{
					[]byte("foo"),
					[]byte("foo"),
					[]byte("foo"),
					[]byte("baz.go"),
				},
			},
			{
				Insertions: 10,
				Deletions:  20,
				FullPath:   []byte("bar/foo/baz.go"),
				Segments: [][]byte{
					[]byte("bar"),
					[]byte("foo"),
					[]byte("baz.go"),
				},
			},
			{
				Insertions: 5,
				Deletions:  10,
				FullPath:   []byte("main.go"),
				Segments: [][]byte{
					[]byte("main.go"),
				},
			},
		},
	}
	type fields struct {
		Dirs  int
		Split []byte
	}
	type args struct {
		diff *data_transfer.Diff
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *data_transfer.Diff
	}{
		{
			name: "Maps diffs by directory depth of 1",
			fields: fields{
				Dirs:  1,
				Split: []byte("/"),
			},
			args: args{
				diff: diffs,
			},
			want: &data_transfer.Diff{
				Insertions: 55,
				Deletions:  110,
				Total:      165,
				Rows: []*data_transfer.DiffRow{
					{
						Insertions: 40,
						Deletions:  80,
						FullPath:   []byte("foo"),
						Segments: [][]byte{
							[]byte("foo"),
						},
					},
					{
						Insertions: 10,
						Deletions:  20,
						FullPath:   []byte("bar"),
						Segments: [][]byte{
							[]byte("bar"),
						},
					},
					{
						Insertions: 5,
						Deletions:  10,
						FullPath:   []byte("main.go"),
						Segments: [][]byte{
							[]byte("main.go"),
						},
					},
				},
			},
		},
		{
			name: "Maps diffs by directory depth of 2",
			fields: fields{
				Dirs:  2,
				Split: []byte("/"),
			},
			args: args{
				diff: diffs,
			},
			want: &data_transfer.Diff{
				Insertions: 55,
				Deletions:  110,
				Total:      165,
				Rows: []*data_transfer.DiffRow{
					{
						Insertions: 20,
						Deletions:  40,
						FullPath:   []byte("foo/bar"),
						Segments: [][]byte{
							[]byte("foo"),
							[]byte("bar"),
						},
					},
					{
						Insertions: 10,
						Deletions:  20,
						FullPath:   []byte("foo/baz.go"),
						Segments: [][]byte{
							[]byte("foo"),
							[]byte("baz.go"),
						},
					},
					{
						Insertions: 10,
						Deletions:  20,
						FullPath:   []byte("foo/foo"),
						Segments: [][]byte{
							[]byte("foo"),
							[]byte("foo"),
						},
					},
					{
						Insertions: 10,
						Deletions:  20,
						FullPath:   []byte("bar/foo"),
						Segments: [][]byte{
							[]byte("bar"),
							[]byte("foo"),
						},
					},
					{
						Insertions: 5,
						Deletions:  10,
						FullPath:   []byte("main.go"),
						Segments: [][]byte{
							[]byte("main.go"),
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dd := &transform.DirectoryDiffMapReducer{
				Dirs:  tt.fields.Dirs,
				Split: tt.fields.Split,
			}
			if got, _ := transform.Reduce(dd, tt.args.diff); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDirectoryDiffMapReducer(t *testing.T) {
	type args struct {
		dirs  int
		split []byte
	}
	tests := []struct {
		name string
		args args
		want transform.Reducer
	}{
		{
			name: "Happy Path",
			args: args{
				dirs:  1,
				split: []byte("/"),
			},
			want: &transform.DirectoryDiffMapReducer{
				Dirs:  1,
				Split: []byte("/"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transform.NewDirectoryDiffMapReducer(tt.args.dirs, tt.args.split); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDirectoryDiffMapReducer() = %v, want %v", got, tt.want)
			}
		})
	}
}
