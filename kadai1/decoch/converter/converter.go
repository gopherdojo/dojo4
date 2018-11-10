package converter

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func getName(fileName string) string {
	extension := filepath.Ext(fileName)
	return fileName[0 : len(fileName)-len(extension)]
}

func convertToJpg(src string) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	out, err := os.Create(getName(src) + ".jpg")
	if err != nil {
		return err
	}
	defer out.Close()

	jpeg.Encode(out, img, nil)

	os.Remove(src)
	return nil
}

func convertToPng(src string) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	out, err := os.Create(getName(src) + ".png")
	if err != nil {
		return err
	}
	defer out.Close()

	png.Encode(out, img)

	os.Remove(src)
	return nil
}

func convertToGif(src string) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	out, err := os.Create(getName(src) + ".gif")
	if err != nil {
		return err
	}
	defer out.Close()

	gif.Encode(out, img, nil)

	os.Remove(src)
	return nil
}
