package converter

import (
	"fmt"
	"image"
	"io"
	"strings"
)

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
