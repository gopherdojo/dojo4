package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func Test_runCmd(t *testing.T) {
	// given
	targetDir := copyToTempDir(t, "testdata")
	expected := []string{
		filepath.Join(targetDir, "Jpeg.png"),
		filepath.Join(targetDir, "foo", "Jpeg.png"),
		filepath.Join(targetDir, "foo", "baz", "Jpeg.png"),
		filepath.Join(targetDir, "foo", "baz", "bar", "Jpeg.png"),
	}

	// when
	got := runCmd([]string{targetDir})

	// then
	if got != 0 {
		t.Errorf("runCmd() = %v, want %v", got, 0)
	}

	// and
	for _, created := range expected {
		if _, err := os.Stat(created); err != nil {
			t.Error("unexpected error:", err)
			return
		}
	}
}

func Test_runCmd_err(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "many args", args: args{[]string{"foo", "baz"}}, want: 1},
		{name: "no args", args: args{[]string{}}, want: 1},
		{name: "nil args", args: args{nil}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runCmd(tt.args.args); got != tt.want {
				t.Errorf("runCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func copyToTempDir(t *testing.T, src string) string {
	t.Helper()

	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	err = filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return os.MkdirAll(filepath.Join(tempDir, path), info.Mode())
		}

		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			t.Fatal("unexpected error:", err)
		}
		err = ioutil.WriteFile(filepath.Join(tempDir, path), bytes, info.Mode())
		if err != nil {
			t.Fatal("unexpected error:", err)
		}

		return nil
	})

	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	return filepath.Join(tempDir, src)
}
