package converter

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"strings"
)

type imgConverter interface {
	convert(dist io.Writer) error
}

func (c *Converter) resolveConverter(src io.Reader) (imgConverter, error) {
	img, _, err := image.Decode(src)
	if err != nil {
		return nil, fmt.Errorf("Failed to decode: %s", err)
	}

	to := strings.ToLower(c.opt.ToFormat())
	var ic imgConverter
	switch to {
	case "jpeg", "jpg":
		ic = &jpgConverter{img}
	case "png":
		ic = &pngConverter{img}
	case "gif":
		ic = &gifConverter{img}
	default:
		return nil, fmt.Errorf("%s is not supported format", c.opt.ToFormat())
	}

	return ic, nil
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
