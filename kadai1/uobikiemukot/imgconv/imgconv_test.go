package imgconv

import (
	"os"
	"testing"
)

func TestSearch(t *testing.T) {
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
		result, err := search(d.root, d.config.InputFormat)
		if err != nil {
			t.Errorf("search() failed: %s\n", err)
		}
		if result[0] != d.expect {
			t.Errorf("result unmatch: want(%s) got(%s)\n", d.expect, result)
		}
	}
}

func TestDecode(t *testing.T) {
	data := []struct {
		config Config
		input  string
		output string
	}{
		{
			input:  "test/subdir/a/1.png",
		},
		{
			input:  "test/subdir/a/c/2.jpg",
		},
		{
			input:  "test/subdir/b/3.gif",
		},
	}

	for _, d := range data {
		_, err := decode(d.input)
		if err != nil {
			t.Errorf("ConvertImage() failed: %s\n", err)
		}
		defer os.Remove(d.output)
	}
}

func TestEncode(t *testing.T) {

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
		img, err := decode(d.input)
		if err != nil {
			t.Errorf("decode(%s): %s\n", d.input, err)
		}

		err = d.config.encode(d.output, img)
		if err != nil {
			t.Errorf("ConvertImage() failed: %s\n", err)
		}
		defer os.Remove(d.output)
	}
}
