package converter

import (
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
)

func ConvertImg(file *os.File, directory string) {
	format := "png"
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	switch format {
	case "png":
		out, err := os.Create(file.Name() + ".png")
		if err != nil {
			panic(err)
		}
		convertToPng(&img, out)
	}
}

func convertToPng(img *image.Image, out *os.File) {
	err := png.Encode(out, *img)
	if err != nil {
		panic(err)
	}
}
