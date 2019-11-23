package algorithm_test

import (
	"github.com/nathanburkett/diff-chart/algorithm"
	"github.com/nathanburkett/diff-chart/datatransfer"
	"reflect"
	"testing"
)

var item1 = datatransfer.DiffRow{
	Insertions: 49,
	Deletions:  0,
	FullPath:   []byte("foo/bar"),
	Segments:   [][]byte{
		[]byte("foo"),
		[]byte("bar"),
	},
}

var item2 = datatransfer.DiffRow{
	Insertions: 25,
	Deletions:  25,
	FullPath:   []byte("foo/baz"),
	Segments:   [][]byte{
		[]byte("foo"),
		[]byte("baz"),
	},
}

var item3 = datatransfer.DiffRow{
	Insertions: 0,
	Deletions:  51,
	FullPath:   []byte("bar/baz"),
	Segments:   [][]byte{
		[]byte("bar"),
		[]byte("baz"),
	},
}

var testBed = &datatransfer.Diff{
	Rows: []*datatransfer.DiffRow{
		&item1,
		&item2,
		&item3,
	},
}

func TestTotalDeltaDescendingSorter_Sort(t *testing.T) {
	type fields struct {
		Diff *datatransfer.Diff
	}
	tests := []struct {
		name   string
		fields fields
		want   *datatransfer.Diff
	}{
		{
			name:   "Descending sort works as expected",
			fields: fields{
				Diff: testBed,
			},
			want:   &datatransfer.Diff{
				Rows: []*datatransfer.DiffRow{
					&item3,
					&item2,
					&item1,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dds := algorithm.NewTotalDeltaDescendingSorter()
			if result, _ := algorithm.Sort(dds, tt.fields.Diff); !reflect.DeepEqual(result, tt.want) {
				t.Errorf("After sort got = %v, want %v", result, tt.want)
			}
		})
	}
}
