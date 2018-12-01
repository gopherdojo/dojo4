package reader

import (
	"os"
	"testing"
)

func TestHasExtension(t *testing.T) {
	ex := "dummy"
	mock := fileInfoMock{name: ex}
	if !hasExtension(&mock, []string{ex}) {
		t.Fatal("expected true, actual false")
	}
}

func TestNotHasExtension(t *testing.T) {
	ex1 := "dummy1"
	ex2 := "dummy2"
	mock := fileInfoMock{name: ex1}
	if hasExtension(&mock, []string{ex2}) {
		t.Fatal("expected false, actual true")
	}
}

type fileInfoMock struct {
	name string
	os.FileInfo
}

func (f *fileInfoMock) Name() string {
	return f.name
}
