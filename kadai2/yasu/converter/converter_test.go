package converter

import (
	"image"
	"os"
	"testing"
)

func TestConvertImg(t *testing.T) {
	file, err := os.Open("../sample_image/3d.jpg")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	directory := "sample_image"
	err = ConvertImg(file, directory, "png")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}

func TestconvertToGif(t *testing.T) {
	file, err := os.Open("../sample_image/3d.jpg")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	out, err := os.Create(file.Name() + ".gif")
	if err := convertToGif(&img, out); err != nil {
		t.Fatalf("failed test %#v", err)
	}
}

func TestconvertToPng(t *testing.T) {
	file, err := os.Open("../sample_image/3d.jpg")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	out, err := os.Create(file.Name() + ".gif")
	if err := convertToPng(&img, out); err != nil {
		t.Fatalf("failed test %#v", err)
	}
}

func TestconvertToJpg(t *testing.T) {
	file, err := os.Open("../sample_image/3d.jpg")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	out, err := os.Create(file.Name() + ".gif")
	if err := convertToJpeg(&img, out); err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
