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

func (PngConv) IsConvertible(path string, info os.FileInfo) bool {
	if info.IsDir() {
		return false
	}

	file, err := os.Open(path)
	if err != nil {
		return false
	}
	defer file.Close()

	_, format, err := image.Decode(file)
	if err != nil {
		return false
	}

	return format == "jpeg"
}
