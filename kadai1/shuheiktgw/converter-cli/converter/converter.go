// Package converter provides a functionality to convert image extension
package converter

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// Handled interface represents the error is properly handled and expected
type Handled interface {
	Handled() bool
}

// HandedError represents the error is properly handled and expected
type HandedError struct {
	Message string
}

// Error returns error message for HandedError
func (r *HandedError) Error() string {
	return r.Message
}

// Handled suggests that HandedError implemented Handled interface
func (r *HandedError) Handled() bool {
	return true
}

// Convert converts file extension in the specified path recursively
func Convert(from, to, path string) ([]string, error) {
	filePaths, err := findFilePaths(from, path)
	if err != nil {
		return nil, err
	}

	for _, filePath := range filePaths {
		if err := convert(to, filePath); err != nil {
			return nil, err
		}
	}

	return filePaths, nil
}

func findFilePaths(from, path string) ([]string, error) {
	var filePaths []string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if _, ok := err.(*os.PathError); ok {
			return &HandedError{Message: err.Error()}
		}

		if err != nil {
			return err
		}

		if filepath.Ext(path) == from {
			filePaths = append(filePaths, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	if len(filePaths) == 0 {
		return nil, &HandedError{Message: fmt.Sprintf("could not find files with the specified extension. path: %s, extension: %s", path, from)}
	}

	return filePaths, nil
}

func convert(to, filePath string) error {
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		return err
	}

	img, _, err := image.Decode(file)

	if err != nil {
		return err
	}

	out, err := os.Create(newFilePath(to, filePath))
	defer out.Close()

	if err != nil {
		return err
	}

	switch to {
	case ".gif":
		return gif.Encode(out, img, nil)
	case ".jpeg", ".jpg":
		return jpeg.Encode(out, img, nil)
	case ".png":
		return png.Encode(out, img)
	default:
		return fmt.Errorf("unsupported extension is specified: %s", to)
	}
}

func newFilePath(to, filePath string) string {
	ext := filepath.Ext(filePath)
	return filePath[:len(filePath)-len(ext)] + to
}
