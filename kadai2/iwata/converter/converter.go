package converter

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

type imgConverter interface {
	convert(dist io.Writer) error
}

type pngConverter struct {
	org image.Image
}

func (c pngConverter) convert(dist io.Writer) error {
	return png.Encode(dist, c.org)
}

type jpgConverter struct {
	org image.Image
}

func (c jpgConverter) convert(dist io.Writer) error {
	return jpeg.Encode(dist, c.org, nil)
}

type gifConverter struct {
	org image.Image
}

func (c gifConverter) convert(dist io.Writer) error {
	return gif.Encode(dist, c.org, nil)
}
