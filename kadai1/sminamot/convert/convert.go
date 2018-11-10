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

// A Convert represents an conversion object that includes target file paths to convert and input/output formats
type Convert struct {
	Files []string
	Src   string
	Dst   string
}

// New returns a new Convert that includes target file paths to convert and input/output formats
func New(dir, src, dst string) (*Convert, error) {
	if !validateExtension(src) {
		return nil, errors.New("src extension is wrong")
	}
	if !validateExtension(dst) {
		return nil, errors.New("dst extension is wrong")
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

	return &Convert{
		Files: files,
		Src:   src,
		Dst:   dst,
	}, nil
}

func validateExtension(e string) bool {
	return (e == "jpg" || e == "png" || e == "gif")
}

// Convert executes image conversion
func (c *Convert) Convert() {
	for _, v := range c.Files {
		sf, err := os.Open(v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer sf.Close()
		img, _, err := image.Decode(sf)
		if err != nil {
			fmt.Println(err)
			continue
		}

		df, err := os.Create(strings.TrimSuffix(v, c.Src) + c.Dst)
		if err != nil {
			fmt.Println(err)
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
