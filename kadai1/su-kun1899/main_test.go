package main_test

import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
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
