package imgconv

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

type imgConvertToError struct {
	Code int
	Msg  string
}

// 引数「fromFormat」「toFormat」が変換可能なフォーマットかチェックします。
// Check the format of the parameter.
func FormatCheck(fromFormat string, toFormat string) bool {
	if fromFormat == toFormat {
		return false
	}
	if fromFormat != "jpg" && fromFormat != "png" && fromFormat != "gif" {
		return false
	}
	if toFormat != "jpg" && toFormat != "png" && toFormat != "gif" {
		return false
	}
	return true
}

// 「dir」で指定したディレクトリ配下全ての「fromFormat」の画像形式に一致する画像ファイルを「toFormat」の画像形式に変換します。
// Convert all images under the specified directory.
// Target the file with the argument "fromFormat".
// Convert to "toFormat" image format.
func Convert(dir string, fromFormat string, toFormat string) error {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fileExt := filepath.Ext(info.Name())
			newFileName := strings.TrimSuffix(path, fileExt)

			if isTargetFile(fromFormat, fileExt) {
				fmt.Println(path + " -> " + newFileName + "." + toFormat)
				errConv := convertTo(path, newFileName+"."+toFormat, toFormat)
				if errConv.Code != 0 {
					fmt.Println(errConv.Msg)
				}
			}
		}
		return nil
	})
	return err
}

// 「fileExt」が指定した画像形式「fromFormat」に一致する拡張子か調べます。
// Determines whether "fileExt" matches the specified image format "fromFormat" extension.
func isTargetFile(fromFormat string, fileExt string) bool {
	if fromFormat == "jpg" && (fileExt == ".jpg" || fileExt == ".jpeg") {
		return true
	}
	if fromFormat == "gif" && fileExt == ".gif" {
		return true
	}
	if fromFormat == "png" && fileExt == ".png" {
		return true
	}
	return false
}

//　「fromFilePath」のファイルを「toFormat」の画像形式に変換し、「toFilePath」のファイルとして保存します。
// Convert "fromFilePath" file to "toFormat" image format and save it as "toFilePath" file.
func convertTo(fromFilePath string, toFilePath string, toFormat string) imgConvertToError {
	file, err := os.Open(fromFilePath)
	if err != nil {
		return imgConvertToError{100, "input file open error."}
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return imgConvertToError{101, "input file decode error."}
	}

	out, err := os.Create(toFilePath)
	if err != nil {
		return imgConvertToError{102, "output file create error."}
	}
	defer out.Close()

	var errEnc error
	if toFormat == "jpg" {
		opts := &jpeg.Options{Quality: 80}
		errEnc = jpeg.Encode(out, img, opts)
	} else if toFormat == "png" {
		errEnc = png.Encode(out, img)
	} else if toFormat == "gif" {
		opts := &gif.Options{NumColors: 256}
		errEnc = gif.Encode(out, img, opts)
	}
	if errEnc != nil {
		return imgConvertToError{103, "input file decode error."}
	}
	return imgConvertToError{0, "encode complete."}
}
