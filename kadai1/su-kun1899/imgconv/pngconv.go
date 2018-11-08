package imgconv

import (
	"errors"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"os"
)

type ImageConv interface {
	Convert(src, dest string) error
}

type PngConv struct{}

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
		return errors.New(fmt.Sprintf("format is not jpeg. src = %v", src))
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
