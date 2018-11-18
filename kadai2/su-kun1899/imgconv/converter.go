package imgconv

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

const JpegFormat = "jpeg"
const PngFormat = "png"

// Is checks whether path's format is imageFormat or not.
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

func ConvertTo(target, imageFormat string) error {
	reader, err := os.Open(target)
	if err != nil {
		return err
	}
	defer reader.Close()

	img, format, err := image.Decode(reader)
	if err != nil {
		return err
	} else if format != "jpeg" {
		// TODO jpegやpng以外も対応する
		return fmt.Errorf("format is not jpeg. src = %v", target)
	}

	// TODO jpegやpng以外も対応する
	dest := replaceExt(target, "png")
	writer, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer writer.Close()

	// TODO jpegやpng以外も対応する
	err = png.Encode(writer, img)
	if err != nil {
		return err
	}

	return nil
}

func replaceExt(fileName, newExt string) string {
	return fmt.Sprintf("%s.%s", strings.TrimSuffix(fileName, filepath.Ext(fileName)), newExt)
}
