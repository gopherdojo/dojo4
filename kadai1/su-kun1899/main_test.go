package main_test

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	_ "image/gif"
	"os"
	"reflect"
	"testing"
)

func Test_format(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		want     string
	}{
		{"gif", "testdata/syokuji_computer.gif", "gif"},
		{"png", "testdata/syokuji_computer.png", "png"},
		{"jpeg", "testdata/syokuji_computer.jpg", "jpeg"},
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
