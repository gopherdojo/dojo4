package convimg

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

type ConvImg struct {
	From string
	To string
	Path string
}

func NewConvert(from string, to string, path string) *ConvImg {
	return &ConvImg{From: from, To: to, Path: path}
}

func (c *ConvImg) Convert() error {
	files, err := findImages(c.Path, c.From, log.New(os.Stdout, "convimg",1))
	if err != nil {
		return err
	}
	switch c.To {
	case "png":
		fmt.Printf("%v", files)
		err = convertToPngs(files)
	case "jpg":
		err = convertToJpgs(files)
	case "gif":
		err = convertToGifs(files)
	default:
		return errors.New("invalid extension")
	}
	if err != nil {
		fmt.Printf("%v", err)
	}
	return  nil
}

func findImages(path string, from string, log *log.Logger) ([]string, error) {
	var files []string
	if path == "" {
		return nil, errors.New("invalid file path")
	}
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if ext := filepath.Ext(info.Name()); ext == "." + from {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func convertToPngs(paths []string) error {
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

func convertToJpgs(paths []string) error {
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

func convertToGifs(paths []string) error {
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
