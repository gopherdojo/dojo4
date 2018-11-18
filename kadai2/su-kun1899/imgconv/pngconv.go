package imgconv

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"os"
)

// ImageConv is a interface for converter of image format.
type ImageConv interface {
	Convert(src, dest string) error
	IsConvertible(path string) bool
}

// PngConv is a struct for converter of png format.
type PngConv struct{}

// Convert converts src file to png format file.
// The new file outs in dest file.
func (PngConv) Convert(src, dest string) error {
	reader, err := os.Open(src)
	if err != nil {
		return err
	}
	defer reader.Close()

	img, format, err := image.Decode(reader)
	if err != nil {
		return err
	} else if format != "jpeg" {
		return fmt.Errorf("format is not jpeg. src = %v", src)
	}

	writer, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer writer.Close()

	err = png.Encode(writer, img)
	if err != nil {
		return err
	}

	return nil
}
