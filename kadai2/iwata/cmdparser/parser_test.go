package cmdparser_test

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/gopherdojo/dojo4/kadai2/iwata/cmdparser"
)

func cmdArgs(cmd string) []string {
	return strings.Split(fmt.Sprintf("%s %s", "imgconv", cmd), " ")
}

func TestCmd_Parse(t *testing.T) {
	tests := []struct {
		name         string
		cmd          string
		wantStdError string
		wantConfig   *cmdparser.CmdConfig
		wantErr      error
	}{
		{"default options", "./", "", cmdparser.NewCmdConfig("./", "jpg", "png"), nil},
		{"with valid options", "-from png -to gif ./", "", cmdparser.NewCmdConfig("./", "png", "gif"), nil},
		{"need only one argument", "./ ./", "", nil, errors.New("only one arg")},
		{"same formats", "-from jpg -to jpg ./", "", nil, errors.New("Cannot set the same format")},
		{"show help message", "-h", "Usage of", nil, errors.New("Failed to paser args")},
		{"with invalid options", "-hoge -fuga", "flag provided but not defined", nil, errors.New("Failed to paser args")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errStream := new(bytes.Buffer)
			c := cmdparser.NewCmd(errStream)
			got, err := c.Parse(cmdArgs(tt.cmd))
			testError(t, err, tt.wantErr, "Cmd.Parse() Error")
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

func testError(t *testing.T, gotErr, wantErr error, msg string) {
	t.Helper()

	if gotErr == nil && wantErr == nil {
		return
	} else if gotErr == nil {
		t.Fatalf("%s: want [%s] error, but got nil", msg, wantErr)
	} else if wantErr == nil {
		t.Fatalf("%s: got [%s] error, but want nil", msg, gotErr)
	}

	if strings.Contains(gotErr.Error(), wantErr.Error()) == false {
		t.Errorf("%s: [%s] not contains [%s]", msg, gotErr, wantErr)
	}
}
