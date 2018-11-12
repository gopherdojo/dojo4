package encoder

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// Encoder for .png
type Png struct {
}

// Encode for .png
func (Png) Encode(image *image.Image, srcPath string) error {
	newFile := fmt.Sprintf("%s.%s", strings.TrimSuffix(srcPath, filepath.Ext(srcPath)), "png")
	file, err := os.Create(newFile)
	if err != nil {
		return err
	}
	defer file.Close()
	err = png.Encode(file, *image)
	if err != nil {
		return err
	}
	return nil
}
