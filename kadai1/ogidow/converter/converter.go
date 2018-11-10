package converter

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// FindFiles return image files
func FindFiles(dir string, extentions []string) ([]string, error) {
	files := []string{}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if !info.IsDir() {
				for _, ext := range extentions {
					if ext == filepath.Ext(path) {
						files = append(files, path)
					}
				}
			}
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to find files: %s", err)
	}
	return files, nil
}

// Convert is convert image from input format to output format
func Convert(srcPath string, destExtention string) error {
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()

	img, _, err := image.Decode(src)

	if err != nil {
		return fmt.Errorf("failed to convert: %s", err)
	}
	destPath := DestFilePath(srcPath, destExtention)
	dest, err := os.Create(destPath)

	if err != nil {
		return fmt.Errorf("failed to convert: %s", err)
	}

	defer dest.Close()

	switch destExtention {
	case ".png":
		err = png.Encode(dest, img)
	}

	if err != nil {
		return fmt.Errorf("failed to convert: %s", err)
	}
	fmt.Printf("convert %s to %s\n", srcPath, destPath)
	return nil
}

// DestFilePath return output file path
func DestFilePath(src string, destExtention string) string {
	destFilePath := src[:len(src)-len(filepath.Ext(src))]
	return destFilePath + destExtention
}
