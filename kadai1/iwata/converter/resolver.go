package converter

import (
	"fmt"
	"image"
	"io"
	"strings"
)

// ConvertOption is an interface for converter
type ConvertOption interface {
	ToFormat() string
}

func resolveConverter(src io.Reader, opt ConvertOption) (Converter, error) {
	img, _, err := image.Decode(src)
	if err != nil {
		return nil, fmt.Errorf("Failed to decode: %s", err)
	}

	to := strings.ToLower(opt.ToFormat())
	var c Converter
	switch to {
	case "jpeg", "jpg":
		c = &jpgConverter{img}
	case "png":
		c = &pngConverter{img}
	case "gif":
		c = &gifConverter{img}
	default:
		return nil, fmt.Errorf("%s is not supported format", opt.ToFormat())
	}

	return c, nil
}
