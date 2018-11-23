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
func Converts(dirPath, srcFormat, distFormat string) (string, error) {
	var result string
	err := fileWalk(dirPath, srcFormat, func(path string) error {
		img, err := decode(path, srcFormat)
		if err != nil {
			return err
		}
		dest := changeExt(path, distFormat)
		err = convert(img, dest, distFormat)
		if err != nil {
			return err
		}
		result += fmt.Sprintf("%s -> %s\n", path, dest)
		return nil
	})
	if err != nil {
		return "", err
	}
	return result, nil
}

// format引数で指定したフォーマットのファイルを探索
// 取得したファイルのパスを引数として 引数のfuncを実行
func fileWalk(dirPath, format string, f func(string) error) error {
	if f == nil {
		panic("引数fは必須です")
	}
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) != fmt.Sprintf(".%s", format) {
			return nil
		}
		err = f(path)
		if err != nil {
			return err
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
	case "jpg", "jpeg":
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
	// fmt.Printf("dest: %s", dest)
	out, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("convert: %s", err.Error())
	}
	defer out.Close()

	switch format {
	case "jpg", "jpeg":
		err = jpeg.Encode(out, img, nil)
	case "png":
		err = png.Encode(out, img)
	case "gif":
		err = gif.Encode(out, img, nil)
	}

	if err != nil {
		return fmt.Errorf("convert: %s", err.Error())
	}

	return nil
}

// ファイルの拡張子を変更
func changeExt(path, ext string) string {
	return string(path[:len(path)-len(filepath.Ext(path))+1]) + ext
}
