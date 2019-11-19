package input_test

import (
	"fmt"
	"github.com/nathanburkett/diff_table/data_transfer"
	"github.com/nathanburkett/diff_table/input"
	"reflect"
	"testing"
)

func TestMake(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		args    args
		want    input.DiffReader
		wantErr bool
	}{
		{
			name:    fmt.Sprintf("\"%s\" yields CliDiffNumstatReader", input.TypeGit),
			args:    args{
				t: input.TypeGit,
			},
			want:    &input.CliDiffNumstatReader{
				Output: &data_transfer.Diff{},
			},
			wantErr: false,
		},
		{
			name:    "\"foo\" yields error",
			args:    args{
				t: "foo",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := input.Make(tt.args.t)
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
