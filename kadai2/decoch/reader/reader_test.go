package reader

import (
	"os"
	"testing"
	"time"
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
}

func (f *fileInfoMock) Name() string {
	return f.name
}

func (f *fileInfoMock) Size() int64 {
	panic("invalid access")
	return 0
}
func (f *fileInfoMock) Mode() os.FileMode {
	panic("invalid access")
	return os.FileMode(0)
}
func (f *fileInfoMock) ModTime() time.Time {
	panic("invalid access")
	return time.Now()
}
func (f *fileInfoMock) IsDir() bool {
	panic("invalid access")
	return false
}
func (f *fileInfoMock) Sys() interface{} {
	panic("invalid access")
	return nil
}
