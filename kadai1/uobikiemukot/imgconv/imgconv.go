package imgconv

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Config image conversion configuration.
type Config struct {
	InputFormat  string // input image format: "png" or "gif" or "jpg"
	OutputFormat string // output image format: "png" or "gif" or "jpg" 
	JPEGOptions  jpeg.Options // Options for jpeg.Encode()
	GIFOptions   gif.Options // Options for gif.Encode()
}

// search pass a root directory as an argument.
// return all image files included in root directory and sub-directory.
func search(root, suffix string) ([]string, error) {
	var list []string

	f := func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() && !info.IsDir() && strings.HasSuffix(path, suffix) {
			list = append(list, path)
		}
		return nil
	}

	err := filepath.Walk(root, f)
	if err != nil {
		return nil, fmt.Errorf("%s (path: %s)", err, root)
	}

	return list, nil
}

// encode encode image and write to specified path.
func (c *Config) encode(path string, img image.Image) error {
	output, err := os.Create(path)
	if err != nil {
		return err
	}
	defer output.Close()

	switch c.OutputFormat {
	case "png":
		err = png.Encode(output, img)
	case "jpg":
		err = jpeg.Encode(output, img, &c.JPEGOptions)
	case "gif":
		err = gif.Encode(output, img, &c.GIFOptions)
	}
	if err != nil {
		return err
	}

	return nil
}

// decode decode input image by specified path.
func decode(path string) (image.Image, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer input.Close()

	img, _, err := image.Decode(input)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// Converter search images from root directory and convert to specified format.
func (c *Config) Converter(root string) error {
	list, err := search(root, c.InputFormat)
	if err != nil {
		return fmt.Errorf("search(%s): %s", root, err)
	}

	for _, in := range list {
		log.Printf("converting... %s\n", in)
		img, err := decode(in)
		if err != nil {
			return fmt.Errorf("decode(%s): %s", in, err)
		}
		out := strings.TrimSuffix(in, c.InputFormat) + c.OutputFormat
		err = c.encode(out, img)
		if err != nil {
			return fmt.Errorf("encode(%s): %s", in, err)
		}
	}

	return nil
}
