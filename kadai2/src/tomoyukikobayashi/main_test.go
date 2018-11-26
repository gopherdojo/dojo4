package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
)

// fixture
var (
	dataRoot = "testdata"
	emptyDir = "testdata/empty"
)

type ExitCode int

func (e ExitCode) Error() string {
	return fmt.Sprintf("exit code: %d", int(e))
}

func init() {
	// mainロジックの exit を上書き
	exit = func(n int) {
		panic(ExitCode(n))
	}
}

// 前提ファイルがあるかチェック
func checkPrecondition(t *testing.T) func() {
	// testdataの存在確認
	if _, err := os.Stat(dataRoot); err != nil {
		t.Fatal(err)
	}

	// 空のフォルダの作成
	if err := os.Mkdir(emptyDir, 0777); err != nil {
		t.Fatal(err)
	}

	return func() {
		if err := os.RemoveAll(emptyDir); err != nil {
			t.Error(err)
		}
	}
}

// exitに差し込まれたpanic(ExitCode)の内容が意図した通りか検証
func testExit(code int, f func()) (err error) {
	defer func() {
		// exitでわざとpanicさせてrecoverからエラーを取る
		e := recover()
		switch t := e.(type) {
		case ExitCode:
			if int(t) == code {
				err = nil
			} else {
				err = fmt.Errorf("expected exit %v got %v", code, e)
			}
		default:
			err = fmt.Errorf("expected exit %v got %v", code, e)
		}
	}()

	f()

	return errors.New("not exit")
}
func Test_main(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want int
	}{
		{
			name: "invalidArgs",
			args: []string{"main", "-i=invalid", "dir"},
			want: ExitInvalidArgs,
		},
		{
			name: "validArgs",
			args: []string{"main", "-f=jpg", "dir"},
			want: ExitError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args
			err := testExit(tt.want, func() {
				main()
			})
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func Test_Run(t *testing.T) {
	tearDown := checkPrecondition(t)

	tests := []struct {
		name     string
		args     []string
		wantCode int
		wantTxt  string
	}{
		{
			name:     "invalidArgs",
			args:     []string{"main", "-i=invalid", "dir"},
			wantCode: ExitInvalidArgs,
			wantTxt:  "Usage",
		},
		{
			name:     "dirNotSpecified",
			args:     []string{"main"},
			wantCode: ExitInvalidArgs,
			wantTxt:  "Usage",
		},
		{
			name:     "unsupported-f",
			args:     []string{"main", "-f=unsp", "dir"},
			wantCode: ExitInvalidArgs,
			wantTxt:  "supported",
		},
		{
			name:     "unsupported-t",
			args:     []string{"main", "-t=unsp", "dir"},
			wantCode: ExitInvalidArgs,
			wantTxt:  "supported",
		},
		{
			name:     "same-f-t",
			args:     []string{"main", "-f=jpg", "-t=jpg", "dir"},
			wantCode: ExitInvalidArgs,
			wantTxt:  "same",
		},
		{
			name:     "dirNotFount",
			args:     []string{"main", "notfound"},
			wantCode: ExitError,
			wantTxt:  "open",
		},
		{
			name:     "fileNotFount",
			args:     []string{"main", "testdata/empty"},
			wantCode: ExitSuccess,
			wantTxt:  "no file",
		},
		{
			name:     "valid",
			args:     []string{"main", "testdata"},
			wantCode: ExitSuccess,
			wantTxt:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// CLIの書き込みストリームを渡す
			outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
			cli := &CLI{outStream: outStream, errStream: errStream}

			got := cli.Run(tt.args)
			if got != tt.wantCode {
				t.Errorf("Exitcode got=%d, want %d", got, tt.wantCode)
			}

			if len(tt.wantTxt) > 1 {
				if !strings.Contains(errStream.String(), tt.wantTxt) &&
					!strings.Contains(outStream.String(), tt.wantTxt) {
					t.Errorf("Text got=%q, want %q", errStream.String(), tt.wantTxt)
				}
			}
		})
	}

	tearDown()
}
