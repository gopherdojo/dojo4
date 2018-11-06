// ディレクトリを指定する
// 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
// ディレクトリ以下は再帰的に処理する
// 変換前と変換後の画像形式を指定できる（オプション）

package imageconvert

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// Converts ...
// ディレクトリを探索して見つかったファイルをコンバート
func Converts(dirPath, srcFormat, distFormat string) error {
	return fileWalk(dirPath, srcFormat, func(path string) error {
		dest := changeExt(path, distFormat)
		img, err := decode(path, srcFormat)
		if err != nil {
			return err
		}
		err = convert(img, dest, distFormat)
		fmt.Printf("%s -> %s\n", path, dest)
		return err
	})
}

// format引数で指定したフォーマットのファイルを探索
// 取得したファイルのパスを引数として 引数のfuncを実行
func fileWalk(dirPath, format string, f func(string) error) error {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == fmt.Sprintf(".%s", format) {
			if f != nil {
				_err := f(path)
				if _err != nil {
					return _err
				}
			}
		}
		return nil
	})
	return err
}

// srcをimage.Imageにデコードする
func decode(src, format string) (image.Image, error) {
	in, err := os.Open(src)
	if err != nil {
		return nil, fmt.Errorf("decode: %s", err.Error())
	}
	defer in.Close()

	var img image.Image
	switch format {
	case "jpg":
		img, err = jpeg.Decode(in)
	case "png":
		img, err = png.Decode(in)
	case "gif":
		img, err = gif.Decode(in)
	}
	if err != nil {
		return nil, fmt.Errorf("decode: %s", err.Error())
	}

	return img, nil
}

// コンバート
func convert(img image.Image, dest, format string) error {
	out, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("convert: %s", err.Error())
	}
	defer out.Close()

	switch format {
	case "jpg":
	case "png":
		png.Encode(out, img)
	case "gif":
		gif.Encode(out, img, nil)
	}

	return nil
}

// ファイルの拡張子を変更
func changeExt(path, ext string) string {
	return string(path[:len(path)-len(ext)]) + ext
}
