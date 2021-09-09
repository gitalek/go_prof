package main

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestRunCmd(t *testing.T) {
	type param struct {
		name  string
		value string
	}
	tests := []struct {
		name           string
		cmdname        string
		param          param
		wantReturnCode int
	}{
		{
			name:    "should touch a file",
			cmdname: "touch",
			param: param{
				name:  "MY_NAME",
				value: "alex",
			},
			wantReturnCode: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir, err := ioutil.TempDir("", "temp")
			if err != nil {
				t.Fatalf("ioutil.TempDir() error = %v, want nil", err)
			}
			fullpath := path.Join(dir, tt.param.value)
			env := Environment{
				tt.param.name: EnvValue{Value: fullpath},
			}
			if gotReturnCode := RunCmd([]string{tt.cmdname, fullpath}, env); gotReturnCode != tt.wantReturnCode {
				t.Fatalf("RunCmd() = %v, want %v", gotReturnCode, tt.wantReturnCode)
			}
			_, err = os.Stat(fullpath)
			if os.IsNotExist(err) {
				t.Fatalf("file not exists err: %v", err)
			}
		})
	}
}
