package convimg

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

type ConvImg struct {
	// before image extension
	From string
	// after image extension
	To string
	// image dir path
	Path string
}

// NewConvert is constructor
func NewConvert(from string, to string, path string) *ConvImg {
	return &ConvImg{From: from, To: to, Path: path}
}

// Convert images from one extension to another
func (c *ConvImg) Convert() error {
	files, err := c.findImages()
	if err != nil {
		return err
	}
	fmt.Println(files)
	switch c.To {
	case "png":
		err = c.convertToPngs(files)
	case "jpg":
		err = c.convertToJpgs(files)
	case "gif":
		err = c.convertToGifs(files)
	default:
		return errors.New("invalid extension")
	}
	return  nil
}

func (c *ConvImg)findImages() ([]string, error) {
	var files []string
	if c.Path == "" {
		return nil, errors.New("invalid file path")
	}
	err := filepath.Walk(c.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if ext := filepath.Ext(info.Name()); ext == "." + c.From {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (c *ConvImg)convertToPngs(paths []string) error {
	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		decodeImg, _,err := image.Decode(file)
		if err != nil {
			return err
		}
		writeFile, err := os.Create(path[:len(path)-len(".png")] + ".png")
		if err != nil {
			return err
		}
		png.Encode(writeFile, decodeImg)
	}
	return nil
}

func (c *ConvImg) convertToJpgs(paths []string) error {
	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		decodeImg, _,err := image.Decode(file)
		if err != nil {
			return err
		}
		writeFile, err := os.Create(path[:len(path)-len(".jpg")] + ".jpg")
		if err != nil {
			return err
		}
		jpeg.Encode(writeFile, decodeImg, nil)
	}
	return nil
}

func (c *ConvImg) convertToGifs(paths []string) error {
	for _,path := range paths {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		decodeImg, _,err := image.Decode(file)
		if err != nil {
			return err
		}
		writeFile, err := os.Create(path[:len(path)-len(".gif")] + ".gif")
		if err != nil {
			return err
		}
		gif.Encode(writeFile, decodeImg, nil)
	}
	return nil
}
