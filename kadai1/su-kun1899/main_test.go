package main

import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test_convertExt(t *testing.T) {
	before := filepath.Join("testdata", "syokuji_computer.jpg")
	want := filepath.Join("testdata", "syokuji_computer.png")
	got := fmt.Sprintf("%s.%s", strings.TrimSuffix(before, filepath.Ext(before)), "png")

	if got != want {
		t.Errorf("got = %v, want %v", got, want)
		return
	}
}

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
			if got := ReplaceExt(tt.args.fileName, tt.args.newExt); got != tt.want {
				t.Errorf("ReplaceExt() = %v, want %v", got, tt.want)
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
