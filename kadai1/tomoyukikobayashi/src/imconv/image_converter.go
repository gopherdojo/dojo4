/*
Package imconv provides image convert mathods and
some validation logics, like if the package can handle
input format or not.
*/
package imconv

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
)

// Convert convert image from from fromat to to fromat
func Convert(path string, from string, to string) (string, error) {
	if !Supported(from) {
		return "", fmt.Errorf("can not handle %s", from)
	}

	if !Supported(to) {
		return "", fmt.Errorf("can not handle %s", to)
	}

	if isSameFormat(from, to) {
		return "", fmt.Errorf("%s and %s are the same format", from, to)
	}

	parts := strings.Split(path, ".")
	newPath := strings.Join(parts[:len(parts)-1], ".") + "." + to

	src, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return "", fmt.Errorf("cannot open file %s", path)
	}
	defer src.Close()

	img, _, err := image.Decode(src)
	if err != nil {
		log.Fatal(err)
		return "", fmt.Errorf("cannot decode file %s", path)
	}

	dst, err := os.Create(newPath)
	if err != nil {
		log.Fatal(err)
		return "", fmt.Errorf("cannot create file %s", newPath)
	}
	defer dst.Close()

	// HACK caseにベタがきしているが、formats定義を展開した方が良い
	// case条件にsliceを展開した物を使えないのだろうか。。。
	switch strings.ToLower(to) {
	case "jpg", "jpeg":
		if err := jpeg.Encode(dst, img, nil); err != nil {
			return "", err
		}
	case "png":
		if err := png.Encode(dst, img); err != nil {
			return "", err
		}
	default:
		log.Fatal("reached to invaid condition")
	}

	return newPath, nil
}
