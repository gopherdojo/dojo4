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
		name   string
		config imgconv.Converter
		root   string
		expect []string
	}{
		{
			name:   "search png",
			config: imgconv.Converter{InputFormat: "png"},
			root:   "testdata/subdir/a",
			expect: []string{"testdata/subdir/a/1.png", "testdata/subdir/a/c/2.png"},
		},
		{
			name:   "search jpg",
			config: imgconv.Converter{InputFormat: "jpg"},
			root:   "testdata/subdir/a",
			expect: []string{"testdata/subdir/a/1.jpg", "testdata/subdir/a/c/2.jpg"},
		},
		{
			name:   "search gif",
			config: imgconv.Converter{InputFormat: "gif"},
			root:   "testdata/subdir/b",
			expect: []string{"testdata/subdir/b/3.gif"},
		},
	}

	for _, d := range data {
		d := d
		t.Run(d.name, func(t *testing.T) {
			result, err := imgconv.ExportSearch(d.root, d.config.InputFormat)
			if err != nil {
				t.Errorf("search() failed: %s\n", err)
			}
			if !reflect.DeepEqual(result, d.expect) {
				t.Errorf("result unmatch: want(%v) got(%v)\n", d.expect, result)
			}
		})
	}
}

func TestDecode_Success(t *testing.T) {
	data := []struct {
		name string
		path string
	}{
		{
			name: "decode png",
			path: "testdata/subdir/a/1.png",
		},
		{
			name: "decode jpg",
			path: "testdata/subdir/a/c/2.jpg",
		},
		{
			name: "decode gif",
			path: "testdata/subdir/b/3.gif",
		},
	}

	for _, d := range data {
		d := d
		t.Run(d.name, func(t *testing.T) {
			_, err := imgconv.ExportDecode(d.path)
			if err != nil {
				t.Errorf("decode(%s) failed: %s\n", d.path, err)
			}
		})
	}
}

func TestDecode_Failure(t *testing.T) {
	data := []struct {
		name string
		path string
	}{
		{
			name: "file not exist",
			path: "file/not/found.png",
		},
		{
			name: "non image file",
			path: "testdata/subdir",
		},
	}

	for _, d := range data {
		d := d
		t.Run(d.name, func(t *testing.T) {
			_, err := imgconv.ExportDecode(d.path)
			if err == nil {
				t.Errorf("decode(%s) succeeded (must fail)\n", d.path)
			}
		})
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
	data := []struct {
		name      string
		converter imgconv.Converter
	}{
		{
			name:      "convert from png to jpg",
			converter: imgconv.Converter{InputFormat: "png", OutputFormat: "jpg"},
		},
		{
			name:      "convert from jpg to gif",
			converter: imgconv.Converter{InputFormat: "jpg", OutputFormat: "gif"},
		},
		{
			name:      "convert from gif to png",
			converter: imgconv.Converter{InputFormat: "gif", OutputFormat: "png"},
		},
	}

	output := "testdata/converted_image.jpg"
	img := testCreateImage(t, "testdata/subdir/a/1.png")

	for _, d := range data {
		d := d
		t.Run(d.name, func(t *testing.T) {
			err := imgconv.ExportConverterEncode(&d.converter, output, img)
			if err != nil {
				t.Errorf("encode() failed: %s\n", err)
			}
		})
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
