package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReplaceExt(t *testing.T) {
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
