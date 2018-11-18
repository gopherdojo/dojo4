package imgconv

import (
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const JpegFormat = "jpeg"
const PngFormat = "png"
const GifFormat = "gif"

type imgFile interface {
	format() string
	ext() string
	encode(w io.Writer, m image.Image) error
}

type pngFile struct{}

func (*pngFile) format() string {
	return PngFormat
}

func (*pngFile) ext() string {
	return "png"
}

func (*pngFile) encode(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}

type gifFile struct{}

func (*gifFile) format() string {
	return GifFormat
}

func (*gifFile) ext() string {
	return "gif"
}

func (*gifFile) encode(w io.Writer, m image.Image) error {
	return gif.Encode(w, m, nil)
}

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
	var f imgFile
	switch imageFormat {
	case PngFormat:
		f = &pngFile{}
	case GifFormat:
		f = &gifFile{}
	default:
		return fmt.Errorf("%s is unsupported format", imageFormat)
	}

	reader, err := os.Open(target)
	if err != nil {
		return err
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		return err
	}

	dest := replaceExt(target, f.ext())
	writer, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer writer.Close()

	err = f.encode(writer, img)
	if err != nil {
		return err
	}

	return os.Remove(target)
}

func replaceExt(fileName, newExt string) string {
	return fmt.Sprintf("%s.%s", strings.TrimSuffix(fileName, filepath.Ext(fileName)), newExt)
}
