package cmd

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// Convert images's extension
func Convert(dir, from, to string) {
	err := convertFiles(dir, from, to)
	if err != nil {
		fmt.Println(err)
	}
}

func convertFiles(dir, from, to string) error {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == dir {
			return nil
		}
		if filepath.Ext(path) == "."+from {
			println(path)
			return convertFile(path, from, to)
		}
		return nil
	})

	return err
}

func convertFile(path, from, to string) error {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return err
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	newFile, err := os.Create(newFilePath(to, path))
	defer newFile.Close()
	if err != nil {
		return err
	}

	switch to {
	case "gif":
		return gif.Encode(newFile, img, nil)
	case "jpeg", "jpg":
		return jpeg.Encode(newFile, img, nil)
	case "png":
		return png.Encode(newFile, img)
	default:
		return fmt.Errorf("exension is invalid: %s", to)
	}
}

func newFilePath(to, filePath string) string {
	ext := filepath.Ext(filePath)
	return filePath[:len(filePath)-len(ext)] + "." + to
}
