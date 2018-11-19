package imageconvert

import (
	"fmt"
	"os"
	"testing"
)

const dirPath = "test-img"
const fileName = "img.jpg"

var filePath = fmt.Sprintf("%s/%s", dirPath, fileName)

func TestChangeExt(t *testing.T) {
	cases := []struct {
		ext string // 変換先のフォーマット
		in  string
		out string
	}{
		{"gif", "/path/hello.jpg", "/path/hello.gif"},
		{"png", "/path/hello.jpg", "/path/hello.png"},
		{"png", "/path/hello.test.jpg", "/path/hello.test.png"},
	}
	for _, c := range cases {
		actual := changeExt(c.in, c.ext)
		if actual != c.out {
			t.Errorf(`changeExt(%s, %s) => "%s", want "%s"`, c.in, c.ext, c.out, actual)
		}
	}
}

func TestFileWalk(t *testing.T) {
	expect := filePath
	var actual string

	err := fileWalk(dirPath, "jpg", func(path string) error {
		actual = path
		return nil
	})
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if actual != expect {
		t.Errorf(`expect="%s", actual="%s"`, expect, actual)
	}
}

func TestDecode(t *testing.T) {
	i, err := decode(filePath, "jpg")
	if err != nil {
		t.Error(err.Error())
	}
	if i == nil {
		t.Error("imgが見つかりません")
	}
}

func TestConvert(t *testing.T) {
	dest := changeExt(filePath, "gif")
	img, err := decode(filePath, "jpg")
	if err != nil {
		t.Error(err)
	}
	err = convert(img, dest, "git")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(dest)
}

// jpgをpng, gifに変換してみる
func TestConverts(t *testing.T) {
	destFormats := []string{"png", "gif"}
	for _, f := range destFormats {
		err := Converts(dirPath, "jpg", f)
		if err != nil {
			t.Errorf(err.Error())
			return
		}
		_, err = os.Open(changeExt(filePath, f)) // ファイルが存在していることを確認
		if err != nil {
			t.Errorf(err.Error())
			return
		}
		os.Remove(changeExt(filePath, f))
	}
}

// テスト用のjpegファイルを作成
// func createJPEG(filePath string) error {
// 	err := os.Mkdir(dirPath, 0777)
// 	if err != nil {
// 		return err
// 	}
// 	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
// 	f, err := os.Create(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
// 	return jpeg.Encode(f, img, nil)
// }
