package main_test

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func Test_format(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		want     string
	}{
		{"gif", filepath.Join("testdata", "syokuji_computer.gif"), "gif"},
		{"png", filepath.Join("testdata", "syokuji_computer.png"), "png"},
		{"jpeg", filepath.Join("testdata", "syokuji_computer.jpg"), "jpeg"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reader, err := os.Open(test.fileName)
			if err != nil {
				t.Error("unexpected error:", err)
				return
			}
			defer reader.Close()

			_, got, err := image.Decode(reader)
			if err != nil {
				t.Error("unexpected error:", err)
				return
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Format = %v, want %v", got, test.want)
				return
			}
		})
	}
}

func Test_convert(t *testing.T) {
	t.Skip()

	reader, err := os.Open(filepath.Join("testdata", "syokuji_computer.jpg"))
	if err != nil {
		t.Error("unexpected error:", err)
		return
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		t.Error("unexpected error:", err)
		return
	}

	writer, err := os.Create("testdata/writer.png")
	if err != nil {
		t.Error("unexpected error:", err)
		return
	}
	defer writer.Close()
	defer os.Remove("testdata/writer.png")

	err = png.Encode(writer, img)
	if err != nil {
		t.Error("unexpected error:", err)
		return
	}
}

func Test_convertExt(t *testing.T) {
	before := filepath.Join("testdata", "syokuji_computer.jpg")
	want := filepath.Join("testdata", "syokuji_computer.png")
	got := fmt.Sprintf("%s.%s", strings.TrimSuffix(before, filepath.Ext(before)), "png")

	if got != want {
		t.Errorf("got = %v, want %v", got, want)
		return
	}
}
