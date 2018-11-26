package encoder

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"
)

// Encoder for .jpeg
type Jpeg struct {
}

// Encode for .jpeg
func (Jpeg) Encode(image *image.Image, srcPath string) error {
	newFile := fmt.Sprintf("%s.%s", strings.TrimSuffix(srcPath, filepath.Ext(srcPath)), "jpeg")
	file, err := os.Create(newFile)
	if err != nil {
		return err
	}
	defer file.Close()
	err = jpeg.Encode(file, *image, nil)
	if err != nil {
		return err
	}
	return nil
}
