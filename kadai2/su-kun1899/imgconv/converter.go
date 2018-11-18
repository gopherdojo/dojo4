package imgconv

import (
	"image"
	"os"
)

const JpegFormat = "jpeg"

type ImageFile interface {
	ConvertTo(imageFormat string) bool
}

func Is(path, imageFormat string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

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

	return format == imageFormat
}

func ConvertTo(imageFormat string) bool {
	panic("not implemented")
}
