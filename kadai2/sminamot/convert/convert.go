package convert

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrWrongInputSrcExtention = errors.New("src extension is wrong")
	ErrWrongInputDstExtention = errors.New("dst extension is wrong")
)

// A Conversion represents an conversion object that includes target file paths to convert and input/output formats
type Conversion struct {
	Files []string
	Src   string
	Dst   string
}

// New returns a new Conversion that includes target file paths to convert and input/output formats
func New(dir, src, dst string) (*Conversion, error) {
	if !validateExtension(src) {
		return nil, ErrWrongInputSrcExtention
	}
	if !validateExtension(dst) {
		return nil, ErrWrongInputDstExtention
	}

	files := []string{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == "."+src {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &Conversion{
		Files: files,
		Src:   src,
		Dst:   dst,
	}, nil
}

func validateExtension(e string) bool {
	return (e == "jpg" || e == "png" || e == "gif")
}

func filename(f, src, dst string) string {
	return strings.TrimSuffix(f, src) + dst
}

// Convert executes image conversion
func (c *Conversion) Convert() {
	for _, v := range c.Files {
		sf, err := os.Open(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		defer sf.Close()
		img, _, err := image.Decode(sf)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		df, err := os.Create(filename(v, c.Src, c.Dst))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		defer df.Close()

		switch c.Dst {
		case "jpg":
			jpeg.Encode(df, img, nil)
		case "gif":
			gif.Encode(df, img, nil)
		case "png":
			png.Encode(df, img)
		}
	}
}
