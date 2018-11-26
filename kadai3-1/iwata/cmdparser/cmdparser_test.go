package cmdparser_test

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/gopherdojo/dojo4/kadai3-1/iwata/cmdparser"
)

func cmdArgs(cmd string) []string {
	return append([]string{"imgconv"}, strings.Split(cmd, " ")...)
}

func TestCmd_Parse(t *testing.T) {
	tests := []struct {
		name    string
		cmd     string
		want    *cmdparser.Config
		wantErr bool
	}{
		{"default option", "./q.txt", &cmdparser.Config{Timeout: 15, TxtPath: "./q.txt"}, false},
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cmd.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
