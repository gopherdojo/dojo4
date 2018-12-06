package imgconv_test

import (
	"image"
	"image/jpeg"
	"os"
	"reflect"
	"testing"

	"github.com/gopherdojo/dojo4/kadai2/uobikiemukot/imgconv"
)

func TestSearch(t *testing.T) {
	data := []struct {
		config imgconv.Converter
		root   string
		expect []string
	}{
		{
			config: imgconv.Converter{InputFormat: "png"},
			root:   "testdata/subdir/a",
			expect: []string{"testdata/subdir/a/1.png", "testdata/subdir/a/c/2.png"},
		},
		{
			config: imgconv.Converter{InputFormat: "jpg"},
			root:   "testdata/subdir/a",
			expect: []string{"testdata/subdir/a/1.jpg", "testdata/subdir/a/c/2.jpg"},
		},
		{
			config: imgconv.Converter{InputFormat: "gif"},
			root:   "testdata/subdir/b",
			expect: []string{"testdata/subdir/b/3.gif"},
		},
	}

	for _, d := range data {
		result, err := imgconv.ExportSearch(d.root, d.config.InputFormat)
		if err != nil {
			t.Errorf("search() failed: %s\n", err)
		}
		if !reflect.DeepEqual(result, d.expect) {
			t.Errorf("result unmatch: want(%v) got(%v)\n", d.expect, result)
		}
	}
}

func TestDecode_Success(t *testing.T) {
	data := []string{
		"testdata/subdir/a/1.png",
		"testdata/subdir/a/c/2.jpg",
		"testdata/subdir/b/3.gif",
	}

	for _, input := range data {
		_, err := imgconv.ExportDecode(input)
		if err != nil {
			t.Errorf("decode(%s) failed: %s\n", input, err)
		}
	}
}

func TestDecode_Failure(t *testing.T) {
	data := []string{
		"file/not/found.png",
		"testdata/subdir",
	}

	for _, input := range data {
		_, err := imgconv.ExportDecode(input)
		if err == nil {
			t.Errorf("decode(%s) succeeded (must fail)\n", input)
		}
	}
}

func testCreateImage(t *testing.T, path string) image.Image {
	t.Helper()

	img, err := imgconv.ExportDecode(path)
	if err != nil {
		t.Fatal(err)
	}

	return img
}

func TestEncode_Success(t *testing.T) {
	data := []imgconv.Converter{
		{InputFormat: "png", OutputFormat: "jpg"},
		{InputFormat: "jpg", OutputFormat: "gif"},
		{InputFormat: "gif", OutputFormat: "png"},
	}

	output := "testdata/converted_image.jpg"
	img := testCreateImage(t, "testdata/subdir/a/1.png")

	for _, c := range data {
		err := imgconv.ExportConverterEncode(&c, output, img)
		if err != nil {
			t.Errorf("encode() failed: %s\n", err)
		}
	}
	os.Remove(output)
}

func TestRun(t *testing.T) {
	c := imgconv.Converter{
		InputFormat:  "png",
		OutputFormat: "jpg",
		JPEGOptions:  jpeg.Options{Quality: 1},
	}

	err := c.Run("testdata/subdir")
	if err != nil {
		t.Fatalf("Converter() failed: %s\n", err)
	}
}
