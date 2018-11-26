package encoder

import (
	"image"
)

// Basic implementation for encoder
type Encoder interface {
	// Basic implementation for encode
	Encode(image *image.Image, srcPath string) error
}
