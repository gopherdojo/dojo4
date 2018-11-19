package converter

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"regexp"
)

// Converter converts images inside a directory from input type to output type
type Converter struct {
	inputType  string
	outputType string
}

var regexpPath = regexp.MustCompile("\\.(jpg|jpeg|png|gif)$")
var (
	jpgExt  = ".jpg"
	jpegExt = ".jpeg"
	gifExt  = ".gif"
	pngExt  = ".png"
)

// New Converter from input type and output type
func New(inputType string, outputType string) Converter {
	return Converter{inputType: inputType, outputType: outputType}
}

// Convert image
func (c *Converter) ConvertImage(path string) (err error) {
	img, err := c.readImage(path)
	if err != nil {
		return err
	}
	if img == nil {
		return nil
	}
	outputPath := regexpPath.ReplaceAllString(path, "."+c.outputType)
	err = c.writeImage(outputPath, img)
	if err != nil {
		return err
	}
	return nil
}

func (c *Converter) readImage(path string) (image.Image, error) {
	var image image.Image
	fmt.Println("Read file: " + path)
	ext := regexpPath.FindString(path)
	if ext != "."+c.inputType {
		return nil, nil
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	inputExt := "." + c.inputType
	switch inputExt {
	case jpgExt, jpegExt:
		image, err = jpeg.Decode(file)
	case pngExt:
		image, err = png.Decode(file)
	case gifExt:
		image, err = gif.Decode(file)
	default:
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return image, nil
}

func (c *Converter) writeImage(path string, image image.Image) error {
	fmt.Println("Write file: " + path)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	outputExt := "." + c.outputType
	switch outputExt {
	case jpgExt, jpegExt:
		err = jpeg.Encode(file, image, nil)
	case pngExt:
		err = png.Encode(file, image)
	case gifExt:
		err = gif.Encode(file, image, nil)
	}
	if err != nil {
		return err
	}
	return nil
}
