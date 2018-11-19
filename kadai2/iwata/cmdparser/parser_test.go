package cmdparser_test

import (
	"bytes"
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/gopherdojo/dojo4/kadai2/iwata/cmdparser"
	"github.com/gopherdojo/dojo4/kadai2/iwata/testutil"
)

func cmdArgs(cmd string) []string {
	return append([]string{"imgconv"}, strings.Split(cmd, " ")...)
}

func TestCmdConfig_SrcDir(t *testing.T) {
	dir := "/hoge"
	c := cmdparser.NewConfig(dir, "jpg", "png")
	if c.SrcDir() != dir {
		t.Errorf("CmdConfig.SrcDir() = %s, want %s", c.SrcDir(), dir)
	}
}

func TestCmdConfig_FromFormat(t *testing.T) {
	from := "jpg"
	c := cmdparser.NewConfig("./", from, "png")
	if c.FromFormat() != from {
		t.Errorf("CmdConfig.FromFormat() = %s, want %s", c.FromFormat(), from)
	}
}

func TestCmdConfig_ToFormat(t *testing.T) {
	to := "gif"
	c := cmdparser.NewConfig("./", "jpg", to)
	if c.ToFormat() != to {
		t.Errorf("CmdConfig.ToFormat() = %s, want %s", c.ToFormat(), to)
	}
}

func TestCmd_Parse(t *testing.T) {
	tests := []struct {
		name         string
		cmd          string
		wantStdError string
		wantConfig   *cmdparser.Config
		wantErr      error
	}{
		{"default options", "./", "", cmdparser.NewConfig("./", "jpg", "png"), nil},
		{"with valid options", "-from png -to gif ./", "", cmdparser.NewConfig("./", "png", "gif"), nil},
		{"need only one argument", "./ ./", "", nil, errors.New("only one arg")},
		{"same formats", "-from jpg -to jpg ./", "", nil, errors.New("Cannot set the same format")},
		{"show help message", "-h", "Usage of", nil, errors.New("Failed to paser args")},
		{"with invalid options", "-hoge -fuga", "flag provided but not defined", nil, errors.New("Failed to paser args")},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			errStream := new(bytes.Buffer)
			c := cmdparser.NewCmd(errStream)
			got, err := c.Parse(cmdArgs(tt.cmd))
			testutil.ContainsError(t, err, tt.wantErr, "Cmd.Parse() Error")
			if err != nil {
				if !strings.Contains(errStream.String(), tt.wantStdError) {
					t.Errorf("Output=%q, want %q", errStream.String(), tt.wantStdError)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.wantConfig) {
				t.Errorf("Cmd.Parse() = %v, want %v", got, tt.wantConfig)
			}
		})
	}
}
