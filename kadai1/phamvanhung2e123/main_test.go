package main

import (
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"testing"
)

func TestConvert(t *testing.T) {
	var testFixtures = []struct {
		description string
		path        string
		inputType   string
		outputType  string
		outputFiles []string
	}{
		{
			description: "Test convert jpg to png",
			inputType:   "jpg",
			outputType:  "png",
			path:        "fixtures/jpg",
			outputFiles: []string{"fixtures/jpg/input1/gopher1.png", "fixtures/jpg/input2/gopher2.png"},
		},
		{
			description: "Test convert jpeg to png",
			inputType:   "jpeg",
			outputType:  "png",
			path:        "fixtures/jpeg",
			outputFiles: []string{"fixtures/jpeg/input1/gopher1.png", "fixtures/jpeg/input2/gopher2.png"},
		},
		{
			description: "Test convert png to jpg",
			inputType:   "png",
			outputType:  "jpg",
			path:        "fixtures/png",
			outputFiles: []string{"fixtures/png/input1/gopher1.jpg", "fixtures/png/input2/gopher2.jpg"},
		},
		{
			description: "Test convert gif to png",
			inputType:   "gif",
			outputType:  "png",
			path:        "fixtures/gif",
			outputFiles: []string{"fixtures/gif/input1/gopher1.png", "fixtures/gif/input2/gopher2.png"},
		},
		{
			description: "Test convert gif to jpg",
			inputType:   "gif",
			outputType:  "jpg",
			path:        "fixtures/gif",
			outputFiles: []string{"fixtures/gif/input1/gopher1.jpg", "fixtures/gif/input2/gopher2.jpg"},
		},
	}
	for _, testFixture := range testFixtures {
		if err := convert(testFixture.path, testFixture.inputType, testFixture.outputType); err != nil {
			t.Errorf("Error: %s", err)
		}

		for _, file := range testFixture.outputFiles {
			if _, err := os.Stat(file); os.IsNotExist(err) {
				t.Errorf("Expected output file %s %s is not exist", file, err.Error())
			}
			if !isValidFormat(file, testFixture.outputType) {
				t.Errorf("Output file %s is in wrong format", file)
			}
		}

		teardown(testFixture.outputFiles)
	}
}

func teardown(paths []string) {
	for _, path := range paths {
		if err := os.Remove(path); err != nil {
			if os.IsNotExist(err) {
				continue
			}
			log.Fatal(err)
		}
	}
}

func isValidFormat(path string, fileType string) bool {
	file, err := os.Open(path)
	if err != nil {
		return false
	}
	defer file.Close()

	switch fileType {
	case "jpg", "jpeg":
		_, err = jpeg.Decode(file)
	case "png":
		_, err = png.Decode(file)
	case "gif":
		_, err = gif.Decode(file)
	}

	if err != nil {
		return false
	}
	return true
}
