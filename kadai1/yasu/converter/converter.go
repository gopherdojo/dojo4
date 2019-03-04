/*

Converter

Converterパッケージは画像の変換プログラムを提供します。
*/

package converter

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

// ConvertImg は画像ファイルを指定したformatに変換します。
func ConvertImg(file *os.File, directory string, afterFormat string) {
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	switch afterFormat {
	case "png":
		out, err := os.Create(file.Name() + ".png")
		if err != nil {
			panic(err)
		}
		convertToPng(&img, out)
	case "jpeg":
		out, err := os.Create(file.Name() + ".jpg")
		if err != nil {
			panic(err)
		}
		convertToJpeg(&img, out)
	case "gif":
		out, err := os.Create(file.Name() + ".gif")
		if err != nil {
			panic(err)
		}
		convertToGif(&img, out)
	}
}

func convertToPng(img *image.Image, out *os.File) {
	err := png.Encode(out, *img)
	if err != nil {
		panic(err)
	}
}

func convertToGif(img *image.Image, out *os.File) {
	opts := gif.Options{NumColors: 256}
	err := gif.Encode(out, *img, &opts)
	if err != nil {
		panic(err)
	}
}

func convertToJpeg(img *image.Image, out *os.File) {
	opts := jpeg.Options{Quality: 100}
	err := jpeg.Encode(out, *img, &opts)
	if err != nil {
		panic(err)
	}
}
