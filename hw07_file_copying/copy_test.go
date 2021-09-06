package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestCopy(t *testing.T) {
	fromPath := "./testdata/input.txt"
	tmpDir := os.TempDir()
	type args struct {
		fromPath string
		toPath   string
		offset   int64
		limit    int64
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		wantPath string
	}{
		{
			name: "offset0_limit0",
			args: args{
				fromPath: fromPath,
				toPath:   "out_offset0_limit0_result.txt",
				offset:   0,
				limit:    0,
			},
			wantErr:  false,
			wantPath: "./testdata/out_offset0_limit0.txt",
		},
		{
			name: "offset0_limit10",
			args: args{
				fromPath: fromPath,
				toPath:   "out_offset0_limit10_result.txt",
				offset:   0,
				limit:    10,
			},
			wantErr:  false,
			wantPath: "./testdata/out_offset0_limit10.txt",
		},
		{
			name: "offset0_limit1000",
			args: args{
				fromPath: fromPath,
				toPath:   "out_offset0_limit1000_result.txt",
				offset:   0,
				limit:    1_000,
			},
			wantErr:  false,
			wantPath: "./testdata/out_offset0_limit1000.txt",
		},
		{
			name: "offset0_limit10000",
			args: args{
				fromPath: fromPath,
				toPath:   "out_offset0_limit10000_result.txt",
				offset:   0,
				limit:    10_000,
			},
			wantErr:  false,
			wantPath: "./testdata/out_offset0_limit10000.txt",
		},
		{
			name: "offset100_limit1000",
			args: args{
				fromPath: fromPath,
				toPath:   "out_offset100_limit1000_result.txt",
				offset:   100,
				limit:    1_000,
			},
			wantErr:  false,
			wantPath: "./testdata/out_offset100_limit1000.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			toPath := path.Join(tmpDir, tt.args.toPath)
			if err := Copy(tt.args.fromPath, toPath, tt.args.offset, tt.args.limit); (err != nil) != tt.wantErr {
				t.Fatalf("Copy() error = %v, wantErr %v", err, tt.wantErr)
			}

			origin, err := ioutil.ReadFile(tt.wantPath)
			if err != nil {
				t.Fatalf(".ReadFile(tt.wantPath) error = %v, want nil", err)
			}
			result, err := ioutil.ReadFile(toPath)
			if err != nil {
				t.Fatalf(".ReadFile(tt.args.toPath) error = %v, want nil", err)
			}
			if !bytes.Equal(origin, result) {
				t.Fatalf("files are not equal")
			}
		})
	}
}

func TestCopy_ErrorCases(t *testing.T) {
	fromPath := "./testdata/input.txt"
	toPathFilename := "error.txt"
	tmpDir := os.TempDir()
	type args struct {
		fromPath string
		toPath   string
		offset   int64
		limit    int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "offset exceeds file size",
			args: args{
				fromPath: fromPath,
				toPath:   toPathFilename,
				offset:   1_000_000,
				limit:    10,
			},
			wantErr: ErrOffsetExceedsFileSize,
		},
		{
			name: "unsupported file: directory",
			args: args{
				fromPath: "./testdata",
				toPath:   toPathFilename,
				offset:   0,
				limit:    10,
			},
			wantErr: ErrUnsupportedFile,
		},
		{
			name: "unsupported file: character device",
			args: args{
				fromPath: "/dev/urandom",
				toPath:   toPathFilename,
				offset:   0,
				limit:    10,
			},
			wantErr: ErrUnsupportedFile,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			toPath := path.Join(tmpDir, tt.args.toPath)
			gotErr := Copy(tt.args.fromPath, toPath, tt.args.offset, tt.args.limit)
			if nil == gotErr || !errors.Is(gotErr, tt.wantErr) {
				t.Fatalf("Copy() error = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}
