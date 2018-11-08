// Package converter convert image extension to target extension
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

var supportedExts = map[string]bool{
	"png":  true,
	"jpeg": true,
	"jpg":  true,
	"gif":  true,
}

// Execute read file from path, convert file extension to afterExt and save it in outDir
func Execute(path string, outDir string, afterExt string) error {
	if _, ok := supportedExts[afterExt]; !ok {
		return fmt.Errorf("%s is not supported ext", afterExt)
	}

	in, err := os.Open(path)
	if err != nil {
		return err
	}
	defer in.Close()

	img, _, err := image.Decode(in)

	if err != nil {
		return err
	}

	dstDir := filepath.Join(outDir, filepath.Dir(path))
	err = createDir(dstDir)
	if err != nil {
		return err
	}

	filenameWithoutExt := filepath.Base(path[:len(path)-len(filepath.Ext(path))])
	dstFile := filepath.Join(dstDir, filenameWithoutExt) + "." + afterExt
	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer out.Close()

	switch afterExt {
	case "jpg", "jpeg":
		err = jpeg.Encode(out, img, nil)
	case "png":
		err = png.Encode(out, img)
	case "gif":
		err = gif.Encode(out, img, nil)
	}
	if err != nil {
		return err
	}
	return nil
}

// createDir creates directory from path if it does not exist
func createDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
