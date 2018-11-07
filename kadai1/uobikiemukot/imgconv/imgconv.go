package imgconv

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
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

// SearchImages pass a root directory as an argument.
// return all image files included in root directory and sub-directory
func (c *Config) SearchImages(root string) ([]string, error) {
	var list []string

	fp, err := os.Open(root)
	if err != nil {
		return nil, fmt.Errorf("%s (path: %s)", err, root)
	}
	defer fp.Close()

	info, err := fp.Stat()
	if err != nil {
		return nil, fmt.Errorf("%s (path: %s)", err, root)
	}

	if !info.IsDir() {
		return nil, fmt.Errorf("%s is not directory", root)
	}

	infos, err := fp.Readdir(0)
	if err != nil {
		return nil, fmt.Errorf("%s (path: %s)", err, root)
	}

	for _, info := range infos {
		path := filepath.Join(root, info.Name())
		if info.IsDir() {
			newList, err := c.SearchImages(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "SearchImages() failed: %s", err)
			} else {
				list = append(list, newList...)
			}
		} else if info.Mode().IsRegular() {
			if strings.HasSuffix(path, c.InputFormat) {
				list = append(list, path)
			}
		}
	}

	return list, nil
}

func (c *Config) encode(w io.Writer, m image.Image) error {
	var err error

	switch c.OutputFormat {
	case "png":
		err = png.Encode(w, m)
	case "jpg":
		err = jpeg.Encode(w, m, &c.JPEGOptions)
	case "gif":
		err = gif.Encode(w, m, &c.GIFOptions)
	}

	return err
}

// ConvertImage convert image file (from Config.InputFormat to Config.OutputFormat).
func (c *Config) ConvertImage(ipath string) error {
	in, err := os.Open(ipath)
	if err != nil {
		return fmt.Errorf("%s (path: %s)", err, ipath)
	}
	defer in.Close()

	img, _, err := image.Decode(in)
	if err != nil {
		return fmt.Errorf("%s (path: %s)", err, ipath)
	}

	opath := strings.TrimSuffix(ipath, c.InputFormat) + c.OutputFormat
	out, err := os.Create(opath)
	if err != nil {
		return fmt.Errorf("%s (path: %s)", err, opath)
	}
	defer out.Close()

	err = c.encode(out, img)
	if err != nil {
		return fmt.Errorf("%s (path: %s)", err, opath)
	}

	return nil
}
