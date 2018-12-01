package cmdparser_test

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/gopherdojo/dojo4/kadai3-2/iwata/cmdparser"
)

func cmdArgs(cmd string) []string {
	return append([]string{"cmdparser"}, strings.Split(cmd, " ")...)
}

func TestCmd_Parse(t *testing.T) {
	pwd, _ := os.Getwd()

	tests := []struct {
		name    string
		cmd     string
		want    *cmdparser.Config
		wantErr bool
	}{
		{
			"default option",
			"https://example.com",
			&cmdparser.Config{
				Parallel: uint(runtime.NumCPU()),
				Timeout:  15,
				Output:   "./",
				URL:      "https://example.com",
			},
			false,
		},
		{
			"set options",
			fmt.Sprintf("-n 6 -timeout 10 -o %s https://example.com", pwd),
			&cmdparser.Config{
				Parallel: 6,
				Timeout:  10,
				Output:   pwd,
				URL:      "https://example.com",
			},
			false,
		},
		{
			"with invalid options",
			"-p 6 https://example.com",
			nil,
			true,
		},
		{
			"wrong arguments number",
			"https://example.com https://yahoo.co.jp",
			nil,
			true,
		},
		{
			"output dir does not exist",
			"-o notexist https://example.com",
			nil,
			true,
		},
		{
			"output dir exists, but not directory",
			"-o cmdparser_test.go https://example.com",
			nil,
			true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			stderr := &bytes.Buffer{}
			c := cmdparser.New(stderr)
			got, err := c.Parse(cmdArgs(tt.cmd))
			if (err != nil) != tt.wantErr {
				t.Errorf("Cmd.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Cmd.Parse() = %v, want %v, differs: (-got +want):\n%s", got, tt.want, diff)
			}
		})
	}
}
