package encoder

import (
	"fmt"
	"image"
	"image/gif"
	"os"
	"path/filepath"
	"strings"
)

// Encoder for .gif
type Gif struct {
}

// Encode for .gif
func (Gif) Encode(image *image.Image, srcPath string) error {
	newFile := fmt.Sprintf("%s.%s", strings.TrimSuffix(srcPath, filepath.Ext(srcPath)), "gif")
	file, err := os.Create(newFile)
	if err != nil {
		return err
	}
	defer file.Close()
	err = gif.Encode(file, *image, nil)
	if err != nil {
		return err
	}
	return nil
}
