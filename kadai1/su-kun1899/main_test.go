package main

import (
	"os"
	"path/filepath"
	"testing"
)

func Test_replaceExt(t *testing.T) {
	type args struct {
		fileName string
		newExt   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "txt to dat", args: args{fileName: "foo.txt", newExt: "dat"}, want: "foo.dat"},
		{name: "jpg to png", args: args{fileName: "foo.jpg", newExt: "png"}, want: "foo.png"},
		{
			name: "path included",
			args: args{
				fileName: filepath.Join("foo", "baz", "bar.jpg"),
				newExt:   "png",
			},
			want: filepath.Join("foo", "baz", "bar.png"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceExt(tt.args.fileName, tt.args.newExt); got != tt.want {
				t.Errorf("replaceExt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listFiles(t *testing.T) {
	err := filepath.Walk("testdata", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		println(path)
		return nil
	})

	if err != nil {
		t.Error("unexpected error:", err)
		return
	}
}

func Test_runCmd(t *testing.T) {
	// given
	targetDir := "testdata"
	expected := []string{
		filepath.Join("testdata", "Jpeg.png"),
		filepath.Join("testdata", "foo", "Jpeg.png"),
		filepath.Join("testdata", "foo", "baz", "Jpeg.png"),
		filepath.Join("testdata", "foo", "baz", "bar", "Jpeg.png"),
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

	// cleanup
	defer func() {
		for _, created := range expected {
			if err := os.Remove(created); err != nil {
				t.Error("unexpected error:", err)
				return
			}
		}
	}()
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
