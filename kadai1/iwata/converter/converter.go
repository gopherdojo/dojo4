package converter

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

// Converter is an interface to convert an image
type Converter interface {
	Convert(dist io.Writer) error
}

type pngConverter struct {
	org image.Image
}

func (c pngConverter) Convert(dist io.Writer) error {
	return png.Encode(dist, c.org)
}

type jpgConverter struct {
	org image.Image
}

func (c jpgConverter) Convert(dist io.Writer) error {
	return jpeg.Encode(dist, c.org, nil)
}

type gifConverter struct {
	org image.Image
}

func (c gifConverter) Convert(dist io.Writer) error {
	return gif.Encode(dist, c.org, nil)
}
