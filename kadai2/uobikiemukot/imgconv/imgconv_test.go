package imgconv

import (
	"image"
	"image/jpeg"
	"os"
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	data := []struct {
		config Converter
		root   string
		expect []string
	}{
		{
			config: Converter{InputFormat: "png"},
			root:   "testdata/subdir/a",
			expect: []string{"testdata/subdir/a/1.png", "testdata/subdir/a/c/2.png"},
		},
		{
			config: Converter{InputFormat: "jpg"},
			root:   "testdata/subdir/a",
			expect: []string{"testdata/subdir/a/1.jpg", "testdata/subdir/a/c/2.jpg"},
		},
		{
			config: Converter{InputFormat: "gif"},
			root:   "testdata/subdir/b",
			expect: []string{"testdata/subdir/b/3.gif"},
		},
	}

	for _, d := range data {
		result, err := search(d.root, d.config.InputFormat)
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
		_, err := decode(input)
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
		_, err := decode(input)
		if err == nil {
			t.Errorf("decode(%s) succeeded (must fail)\n", input)
		}
	}
}

func testCreateImage(t *testing.T, path string) image.Image {
	t.Helper()

	img, err := decode(path)
	if err != nil {
		t.Fatal(err)
	}

	return img
}

func TestEncode_Success(t *testing.T) {
	data := []Converter{
		{InputFormat: "png", OutputFormat: "jpg"},
		{InputFormat: "jpg", OutputFormat: "gif"},
		{InputFormat: "gif", OutputFormat: "png"},
	}

	output := "testdata/converted_image.jpg"
	img := testCreateImage(t, "testdata/subdir/a/1.png")

	for _, config := range data {
		err := config.encode(output, img)
		if err != nil {
			t.Errorf("encode() failed: %s\n", err)
		}
	}
	os.Remove(output)
}

func TestRun(t *testing.T) {
	c := Converter{
		InputFormat:  "png",
		OutputFormat: "jpg",
		JPEGOptions:  jpeg.Options{Quality: 1},
	}

	err := c.Run("testdata/subdir")
	if err != nil {
		t.Fatalf("Converter() failed: %s\n", err)
	}
}
