package transform_test

import (
	"github.com/nathanburkett/diff_table/data_transfer"
	"github.com/nathanburkett/diff_table/transform"
	"reflect"
	"testing"
)

func TestMapReduceDiffByDirectory(t *testing.T) {
	inputDiffs := &data_transfer.Diff{
		Rows: []*data_transfer.DiffRow{
			{
				Insertions: 10,
				Deletions:  20,
				FullPath:   []byte("foo/bar/baz.go"),
				Segments:   [][]byte{
					[]byte("foo"),
					[]byte("bar"),
					[]byte("baz.go"),
				},
			},
			{
				Insertions: 10,
				Deletions:  20,
				FullPath:   []byte("foo/baz.go"),
				Segments:   [][]byte{
					[]byte("foo"),
					[]byte("baz.go"),
				},
			},
			{
				Insertions: 10,
				Deletions:  20,
				FullPath:   []byte("foo/bar/foo.go"),
				Segments:   [][]byte{
					[]byte("foo"),
					[]byte("bar"),
					[]byte("foo.go"),
				},
			},
			{
				Insertions: 10,
				Deletions:  20,
				FullPath:   []byte("foo/foo/foo/baz.go"),
				Segments:   [][]byte{
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
				Segments:   [][]byte{
					[]byte("bar"),
					[]byte("foo"),
					[]byte("baz.go"),
				},
			},
		},
	}

	type args struct {
		diff      *data_transfer.Diff
		dirLevels int
		split []byte
	}
	tests := []struct {
		name string
		args args
		want *data_transfer.Diff
	}{
		{
			name: "Maps diffs by directory depth of 1",
			args: args{
				diff:      inputDiffs,
				dirLevels: 1,
				split: []byte("/"),
			},
			want: &data_transfer.Diff{
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
						Segments:   [][]byte{
							[]byte("bar"),
						},
					},
				},
			},
		},
		{
			name: "Maps diffs by directory depth of 2",
			args: args{
				diff:      inputDiffs,
				dirLevels: 2,
				split: []byte("/"),
			},
			want: &data_transfer.Diff{
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
						Segments:   [][]byte{
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
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transform.MapReduceDiffByDirectory(tt.args.diff, tt.args.dirLevels, tt.args.split); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapReduceDiffByDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}
