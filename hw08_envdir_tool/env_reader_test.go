package main

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"testing"
)

func TestReadDir(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		args    args
		want    Environment
		wantErr bool
	}{
		{
			name: "#1",
			args: args{dir: "./testdata/env"},
			want: Environment{
				"BAR":   EnvValue{Value: "bar"},
				"FOO":   EnvValue{Value: "   foo\nwith new line"},
				"HELLO": EnvValue{Value: `"hello"`},
				"EMPTY": EnvValue{},
				"UNSET": EnvValue{NeedRemove: true},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadDir(tt.args.dir)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ReadDir() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("ReadDir() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadDir_ErrorCases(t *testing.T) {
	type file struct {
		name    string
		content string
	}
	tests := []struct {
		name    string
		env     string
		file    file
		wantErr error
	}{
		{
			name: "Invalid variable name",
			env:  "invalid_variable_name",
			file: file{
				name:    "FOO=",
				content: "foo",
			},
			wantErr: ErrInvalidName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// BEGIN setup
			dir, err := ioutil.TempDir("", tt.env)
			if err != nil {
				t.Fatalf("ioutil.TempDir() error = %v, want nil", err)
			}
			f, err := os.Create(path.Join(dir, tt.file.name))
			if err != nil {
				t.Fatalf("os.Create() error = %v, want nil", err)
			}
			_, err = f.Write([]byte(tt.file.content))
			if err != nil {
				t.Fatalf("f.Write() error = %v, want nil", err)
			}
			err = f.Close()
			if err != nil {
				t.Fatalf("f.Close() error = %v, want nil", err)
			}
			// END setup

			if _, gotErr := ReadDir(dir); nil == gotErr || !errors.Is(gotErr, tt.wantErr) {
				t.Fatalf("ReadDir() error = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}
