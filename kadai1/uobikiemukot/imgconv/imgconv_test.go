package imgconv

import (
	"os"
	"testing"
)

func TestSearchImages(t *testing.T) {
	data := []struct {
		config Config
		root   string
		expect string
	}{
		{
			config: Config{InputFormat: "png"},
			root:   "test/subdir/a",
			expect: "test/subdir/a/1.png",
		},
		{
			config: Config{InputFormat: "jpg"},
			root:   "test/subdir/a",
			expect: "test/subdir/a/c/2.jpg",
		},
		{
			config: Config{InputFormat: "gif"},
			root:   "test/subdir/b",
			expect: "test/subdir/b/3.gif",
		},
	}

	for _, d := range data {
		result, err := d.config.SearchImages("./test/subdir")
		if err != nil {
			t.Errorf("SearchImages() failed: %s\n", err)
		}
		if result[0] != d.expect {
			t.Errorf("result unmatch: want(%s) got(%s)\n", d.expect, result)
		}
	}
}

func TestConvertImage(t *testing.T) {
	data := []struct {
		config Config
		input  string
		output string
	}{
		{
			config: Config{InputFormat: "png", OutputFormat: "jpg"},
			input:  "test/subdir/a/1.png",
			output: "test/subdir/a/1.jpg",
		},
		{
			config: Config{InputFormat: "jpg", OutputFormat: "gif"},
			input:  "test/subdir/a/c/2.jpg",
			output: "test/subdir/a/c/2.gif",
		},
		{
			config: Config{InputFormat: "gif", OutputFormat: "png"},
			input:  "test/subdir/b/3.gif",
			output: "test/subdir/b/3.png",
		},
	}

	for _, d := range data {
		err := d.config.ConvertImage(d.input)
		if err != nil {
			t.Errorf("ConvertImage() failed: %s\n", err)
		}
		defer os.Remove(d.output)
	}
}
