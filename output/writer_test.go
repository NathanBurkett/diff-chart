package output_test

import (
	"bytes"
	"github.com/nathanburkett/diff_table/data_transfer"
	"github.com/nathanburkett/diff_table/output"
	"testing"
)

type MockWriter struct {
	WriteErr error
}

func (mw *MockWriter) Write(p []byte) (n int, err error) {
	if mw.WriteErr != nil {
		return 0, mw.WriteErr
	}

	return 0, nil
}

func TestWrite(t *testing.T) {
	type args struct {
		diff *data_transfer.Diff
	}
	tests := []struct {
		name       string
		args       args
		wantWriter string
		wantErr    bool
	}{
		{
			name:       "Happy path",
			args:       args{
				diff: &data_transfer.Diff{
					Rows: []*data_transfer.DiffRow{
						{
							Insertions: 10,
							Deletions:  20,
							FullPath:   []byte("foo/bar"),
							Segments:   [][]byte{
								[]byte("foo"),
								[]byte("bar"),
							},
						},
						{
							Insertions: 30,
							Deletions:  10,
							FullPath:   []byte("foo/baz.md"),
							Segments:   [][]byte{
								[]byte("foo"),
								[]byte("baz.md"),
							},
						},
						{
							Insertions: 15,
							Deletions:  15,
							FullPath:   []byte("bar/baz.txt"),
							Segments:   [][]byte{
								[]byte("bar"),
								[]byte("baz.txt"),
							},
						},
					},
				},
			},
			wantWriter: `| Directory | +/- | Δ % |
| --- | --- | --- |
| foo/bar | +10/-20 | 30.00% |
| foo/baz.md | +30/-10 | 40.00% |
| bar/baz.txt | +15/-15 | 30.00% |
`,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			err := output.Write(writer, tt.args.diff)
			if (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("Write() gotWriter = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
