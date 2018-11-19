package converter_test

import (
	"github.com/gopherdojo/dojo4/kadai2/phamvanhung2e123/converter"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"testing"
)

type testCase struct {
	description string
	path        string
	inputType   string
	outputType  string
	outputFile  string
}

func TestConvert(t *testing.T) {
	var testFixtures = []testCase{
		{
			description: "Test convert jpg to png",
			inputType:   "jpg",
			outputType:  "png",
			path:        "../fixtures/jpg/input1/gopher1.jpg",
			outputFile:  "../fixtures/jpg/input1/gopher1.png",
		},
		{
			description: "Test convert jpeg to png",
			inputType:   "jpeg",
			outputType:  "png",
			path:        "../fixtures/jpeg/input1/gopher1.jpeg",
			outputFile:  "../fixtures/jpeg/input1/gopher1.png",
		},
		{
			description: "Test convert png to jpg",
			inputType:   "png",
			outputType:  "jpg",
			path:        "../fixtures/png/input1/gopher1.png",
			outputFile:  "../fixtures/png/input1/gopher1.jpg",
		},
		{
			description: "Test convert gif to png",
			inputType:   "gif",
			outputType:  "png",
			path:        "../fixtures/gif/input1/gopher1.gif",
			outputFile:  "../fixtures/gif/input1/gopher1.png",
		},
		{
			description: "Test convert gif to jpg",
			inputType:   "gif",
			outputType:  "jpg",
			path:        "../fixtures/gif/input1/gopher1.gif",
			outputFile:  "../fixtures/gif/input1/gopher1.jpg",
		},
	}

	for _, testFixture := range testFixtures {
		c := converter.New(testFixture.inputType, testFixture.outputType)
		t.Run("Check convert", func(t *testing.T) {
			checkConvert(t, c, testFixture.path)
		})
		t.Run("Check format", func(t *testing.T) {
			checkFormat(t, testFixture.outputFile, testFixture.outputType)
		})
	}
}

func checkConvert(t *testing.T, c converter.Converter, path string)  {
	t.Helper()
	if err := c.ConvertImage(path); err != nil {
		t.Errorf("Error: %s", err)
	}
}

func checkFormat(t *testing.T, path string, fileType string) {
	t.Helper()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Errorf("Expected output file %s %s is not exist", path, err.Error())
	}
	file, err := os.Open(path)
	if err != nil {
		t.Errorf("Couldn't open file path: %s, fileType: %s, error: %v", path, fileType, err)
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
		t.Errorf("Couldn't decode path: %s, fileType: %s, error: %v", path, fileType, err)
	}
}
