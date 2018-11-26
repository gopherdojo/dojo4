package converter

import (
	"fmt"
	"github.com/keitarou/dojo4/kadai1/keitarou/lib/encoder"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
)

// ConvertWalkOption
type ConvertWalkOption struct {
	FilePathRoot string
	SrcExt       string
	DstExt       string
}

// By option, recursion
func ConvertWalk(option ConvertWalkOption) error {
	encoder, err := GetEncoder(option.DstExt)
	if err != nil {
		return err
	}
	err = filepath.Walk(option.FilePathRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != "."+option.SrcExt {
			return nil
		}
		return Convert(path, encoder)
	})
	return err
}

// Convert target file
func Convert(path string, encoder encoder.Encoder) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	image, _, err := image.Decode(file)
	if err != nil {
		return err
	}
	return encoder.Encode(&image, path)
}

// Get target encoder
func GetEncoder(ext string) (encoder.Encoder, error) {
	encoders := map[string]encoder.Encoder{
		"jpeg": encoder.Jpeg{},
		"png":  encoder.Png{},
		"gif":  encoder.Gif{},
	}
	if encoder, ok := encoders[ext]; ok {
		return encoder, nil
	}
	return nil, fmt.Errorf("not found encoder. ext is %s", ext)
}
