package main

import (
	"testing"
	"io/ioutil"
	"bytes"
)

func TestCopy(t *testing.T) {
	type args struct {
		fromPath string
		toPath   string
		offset   int64
		limit    int64
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantPath string
	}{
		{
			name: "offset0_limit0",
			args: args{
				fromPath: "./testdata/input.txt",
				toPath:   "./testdata/out_offset0_limit0_result.txt",
				offset:   0,
				limit:    0,
			},
			wantErr: false,
			wantPath:   "./testdata/out_offset0_limit0.txt",

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Copy(tt.args.fromPath, tt.args.toPath, tt.args.offset, tt.args.limit); (err != nil) != tt.wantErr {
				t.Fatalf("Copy() error = %v, wantErr %v", err, tt.wantErr)
			}

			origin, err := ioutil.ReadFile(tt.wantPath)
			if err != nil {
				t.Fatalf(".ReadAll(tt.wantPath) error = %v, want nil", err)
			}
			result, err := ioutil.ReadFile(tt.args.toPath)
			if err != nil {
				t.Fatalf(".ReadAll(tt.args.toPath) error = %v, want nil", err)
			}
			if !bytes.Equal(origin, result) {
				t.Fatalf("files are not equal")
			}
		})
	}
}
