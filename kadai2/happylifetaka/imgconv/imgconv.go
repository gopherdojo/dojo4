package imgconv

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// ImgConverter 画像変換器。「SetConvertFormat」メソッドでフォーマットを指定して、「Convert」メソッドで変換処理を行います。
// Image converter. Specify the format with the "SetConvertFormat" method and convert it with the "Convert" method.
type ImgConverter struct {
	fromFormat string
	toFormat   string
	setFormat  bool
}

type singleImgConverter struct {
	fromFilePath string
	toFilePath   string
	toFormat     string
}

// Result 「Msg」には変換元ファイルパスと変換先ファイルパスを記すメッセージが格納されています。
// 「Err」には変換時エラーがあった場合、エラーメッセージが格納されています。
// "Msg" contains messages describing the conversion source file path and conversion destination file path.
// "Err" contains an error message if there is an error at conversion.
type Result struct {
	Msg string
	Err error
}

// SetConvertFormat 変換元画像フォーマットと変換後画像フォーマットを指定します。
// Specify the source image format and the converted image format.
func (ic *ImgConverter) SetConvertFormat(fromFormat string, toFormat string) error {
	if fromFormat == toFormat {
		return errors.New("The same value must not be specified for fromFormat and toFormat")
	}
	if fromFormat != "jpg" && fromFormat != "png" && fromFormat != "gif" {
		return errors.New("fromFormat value of A is incorrect.allow value jpg png gif")
	}
	if toFormat != "jpg" && toFormat != "png" && toFormat != "gif" {
		return errors.New("toFormat value of A is incorrect.allow value jpg png gif")
	}
	ic.fromFormat = fromFormat
	ic.toFormat = toFormat
	ic.setFormat = true
	return nil

}

// Convert 「dir」で指定したディレクトリ配下全ての「fromFormat」の画像形式に一致する画像ファイルを「toFormat」の画像形式に変換します。
// Convert all images under the specified directory.
// Target the file with the argument "fromFormat".
// Convert to "toFormat" image format.
func (ic *ImgConverter) Convert(dir string) ([]Result, error) {
	rs := []Result{}
	if ic.setFormat != true {
		return rs, errors.New("not set format")
	}
	if _, err := os.Stat(dir); err != nil {
		return rs, errors.New("target file path is not exist")
	}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fileExt := filepath.Ext(info.Name())
			newFileName := strings.TrimSuffix(path, fileExt)
			if isTargetFile(ic.fromFormat, fileExt) {
				var sic singleImgConverter
				sic.fromFilePath = path
				sic.toFilePath = newFileName + "." + ic.toFormat
				sic.toFormat = ic.toFormat
				err := sic.convertTo()
				if err != nil {
					rs = append(rs, Result{Msg: path + " -> " + newFileName + "." + ic.toFormat, Err: err})
				} else {
					rs = append(rs, Result{Msg: path + " -> " + newFileName + "." + ic.toFormat, Err: nil})
				}
			}
		}
		return nil
	})
	return rs, err
}

// isTargetFile 「fileExt」が指定した画像形式「fromFormat」に一致する拡張子か調べます。
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

//　convertTo 「fromFilePath」のファイルを「toFormat」の画像形式に変換し、「toFilePath」のファイルとして保存します。
// Convert "fromFilePath" file to "toFormat" image format and save it as "toFilePath" file.
func (sic *singleImgConverter) convertTo() error {
	file, err := os.Open(sic.fromFilePath)
	if err != nil {
		return errors.New("input file open error")
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return errors.New("input file decode error")
	}

	out, err := os.Create(sic.toFilePath)
	if err != nil {
		return errors.New("output file create error")
	}
	defer out.Close()

	var errEnc error
	if sic.toFormat == "jpg" {
		opts := &jpeg.Options{Quality: 80}
		errEnc = jpeg.Encode(out, img, opts)
	} else if sic.toFormat == "png" {
		errEnc = png.Encode(out, img)
	} else if sic.toFormat == "gif" {
		opts := &gif.Options{NumColors: 256}
		errEnc = gif.Encode(out, img, opts)
	}
	if errEnc != nil {
		return errors.New("input file decode error")
	}
	return nil
}
