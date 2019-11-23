package algorithm_test

import (
	"fmt"
	"github.com/nathanburkett/diff-chart/algorithm"
	"reflect"
	"testing"
)

func TestMake(t *testing.T) {
	type args struct {
		flag string
	}
	tests := []struct {
		name    string
		args    args
		want    algorithm.Sorter
		wantErr bool
	}{
		{
			name:    fmt.Sprintf("\"%s\" yields TotalDeltaDescendingSorter", algorithm.TypeTotalDeltaDesc),
			args:    args{
				flag: algorithm.TypeTotalDeltaDesc,
			},
			want:    &algorithm.TotalDeltaDescendingSorter{},
			wantErr: false,
		},
		{
			name:    "\"foo\" yields error",
			args:    args{
				flag: "foo",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := algorithm.Make(tt.args.flag)
			if (err != nil) != tt.wantErr {
				t.Errorf("Make() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Make() got = %v, want %v", got, tt.want)
			}
		})
	}
}
