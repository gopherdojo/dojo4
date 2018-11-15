// Package converter converts image extension to target extension
package converter

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
)

var supportedExts = map[string]bool{
	"png":  true,
	"jpeg": true,
	"jpg":  true,
	"gif":  true,
}

// Encoder encodes image
type Encoder interface {
	Encode(io.Writer, image.Image) error
}

// Converter converts image to other format of afterExt
type Converter struct {
	path     string
	outDir   string
	afterExt string
	encoder  Encoder
}

// JpegEncoder encodes image to jpeg format
type JpegEncoder struct {
}

// PngEncoder encodes image to png format
type PngEncoder struct {
}

// GifEncoder encodes image to gif format
type GifEncoder struct {
}

// Encode encodes to jpeg
func (c *JpegEncoder) Encode(w io.Writer, img image.Image) error {
	return jpeg.Encode(w, img, nil)
}

// Encode encodes to png
func (c *PngEncoder) Encode(w io.Writer, img image.Image) error {
	return png.Encode(w, img)
}

// Encode encodes to gif
func (c *GifEncoder) Encode(w io.Writer, img image.Image) error {
	return gif.Encode(w, img, nil)
}

// Encode encodes to specific format
func (c *Converter) Encode(w io.Writer, img image.Image) error {
	return c.encoder.Encode(w, img)
}

// Decode decodes input image
func (c *Converter) Decode(r io.Reader) (image.Image, error) {
	img, _, err := image.Decode(r)
	return img, err
}

// NewConverter creates new converter instance
func NewConverter(path, outDir, afterExt string) (*Converter, error) {
	var encoder Encoder
	switch afterExt {
	case "jpg", "jpeg":
		encoder = &JpegEncoder{}
	case "png":
		encoder = &PngEncoder{}
	case "gif":
		encoder = &GifEncoder{}
	}
	return &Converter{path, outDir, afterExt, encoder}, nil
}

// Execute reads file from path, convert file extension to afterExt and save it in outDir
func (c *Converter) Execute() error {
	if _, ok := supportedExts[c.afterExt]; !ok {
		return fmt.Errorf("%s is not supported ext", c.afterExt)
	}

	in, err := os.Open(c.path)
	if err != nil {
		return err
	}

	defer in.Close()

	img, err := c.Decode(in)
	if err != nil {
		return err
	}

	err = c.createDstDir()
	if err != nil {
		return err
	}

	dstPath, err := c.getDstPath()
	if err != nil {
		return err
	}

	out, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer out.Close()

	err = c.Encode(out, img)
	if err != nil {
		return err
	}
	return nil
}

func (c *Converter) getDstDir() (string, error) {
	srcAbs, err := filepath.Abs(c.path)
	if err != nil {
		return "", err
	}
	return filepath.Join(filepath.Dir(srcAbs), c.outDir), nil
}

func (c *Converter) createDstDir() error {
	dstDir, err := c.getDstDir()
	if err != nil {
		return err
	}

	if _, err := os.Stat(dstDir); os.IsNotExist(err) {
		err = os.MkdirAll(dstDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Converter) getDstPath() (string, error) {
	filenameWithoutExt := filepath.Base(c.path[:len(c.path)-len(filepath.Ext(c.path))])
	dstDir, err := c.getDstDir()
	if err != nil {
		return "", err
	}
	dstPath := filepath.Join(dstDir, filenameWithoutExt) + "." + c.afterExt
	return dstPath, nil
}
