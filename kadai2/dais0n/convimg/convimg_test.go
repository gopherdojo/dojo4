package convimg

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func Test_findImages(t *testing.T) {
	t.Helper()
	cases := [] struct{
		from string
		paths []string
	}{
		{
			from: "jpg",
			paths: []string{"../testdata/gopher.jpg", "../testdata/testdata2/gopher.jpg"},
		},
		{
			from: "png",
			paths: []string{"../testdata/gopher.png", "../testdata/testdata2/gopher.png"},
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("test from %s, data %v", c.from, c.paths), func(t *testing.T) {
			convimg := NewConvert(c.from, "test", "../testdata")
			paths, err := convimg.findImages()
			if err != nil {
				t.Error("unexpected error")
			}
			if !reflect.DeepEqual(paths, c.paths) {
				t.Errorf("unexpected error want %v, get %v", c.paths, paths)
			}
		})
	}
}

func TestConvImg_Convert_InvalidExtension(t *testing.T) {
	t.Helper()
	c := struct {
		from string
		err error
	}{
		from: "invalidExtension",
		err: errors.New("invalid extension"),
	}
	convimg := NewConvert(c.from, "test","../testdata")
	err := convimg.Convert()
	if err.Error() != c.err.Error() {
		t.Errorf("unexpected error want %v get %v", c.err, err)
	}
}

func Test_convertToPngs(t *testing.T) {
}
