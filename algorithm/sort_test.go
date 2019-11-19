package algorithm_test

import (
	"github.com/nathanburkett/diff_table/algorithm"
	"github.com/nathanburkett/diff_table/data_transfer"
	"reflect"
	"testing"
)

var item1 = data_transfer.DiffRow{
	Insertions: 49,
	Deletions:  0,
	FullPath:   []byte("foo/bar"),
	Segments:   [][]byte{
		[]byte("foo"),
		[]byte("bar"),
	},
}

var item2 = data_transfer.DiffRow{
	Insertions: 25,
	Deletions:  25,
	FullPath:   []byte("foo/baz"),
	Segments:   [][]byte{
		[]byte("foo"),
		[]byte("baz"),
	},
}

var item3 = data_transfer.DiffRow{
	Insertions: 0,
	Deletions:  51,
	FullPath:   []byte("bar/baz"),
	Segments:   [][]byte{
		[]byte("bar"),
		[]byte("baz"),
	},
}

var testBed = &data_transfer.Diff{
	Rows: []*data_transfer.DiffRow{
		&item1,
		&item2,
		&item3,
	},
}

func TestTotalDeltaDescendingSorter_Sort(t *testing.T) {
	type fields struct {
		Diff *data_transfer.Diff
	}
	tests := []struct {
		name   string
		fields fields
		want   *data_transfer.Diff
	}{
		{
			name:   "Descending sort works as expected",
			fields: fields{
				Diff: testBed,
			},
			want:   &data_transfer.Diff{
				Rows: []*data_transfer.DiffRow{
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
