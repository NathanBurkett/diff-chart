package output_test

import (
	"bytes"
	"fmt"
	"github.com/nathanburkett/diff-chart/output"
	"os"
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
		want    output.Writer
		wantErr bool
	}{
		{
			name:    fmt.Sprintf("\"%s\" yields MarkdownWriter", output.TypeMarkdownCli),
			args:    args{
				flag: output.TypeMarkdownCli,
			},
			want:    &output.MarkdownWriter{
				Writer:   os.Stdout,
				Internal: &bytes.Buffer{},
			},
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
			got, err := output.Make(tt.args.flag)
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
