package convimg

import (
	"image"
	_ "image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// rootDirにあるファイルの一覧を探索。
// ディレクトリがあれば再帰処理する。

func SearchFile(rootDir, from, to string) {
	files, err := ioutil.ReadDir(rootDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		path := rootDir + "/" + file.Name()
		if file.IsDir() {
			SearchFile(path, from, to)
			continue
		}

		willChange, newPath := generateNewExt(path, from, to)
		if willChange == false {
			continue
		}

		println(convFile(path, newPath))

	}
}

func generateNewExt(path, from, to string) (willConv bool, newPath string) {
	ext := filepath.Ext(path)

	//変換したい拡張子でなければ何もしない
	if ext != from {
		return false, path
	}

	//変換する必要がなければ何もしない
	if from == to {
		return false, path
	}

	return true, path[:len(path)-len(ext)] + to
}

func convFile(path, newPath string) string {
	file, err := os.Open(path)
	if err != nil {
		print("open failed")
		log.Fatal(err)
	}
	defer file.Close()

	decoded, format, err := image.Decode(file)
	if err != nil {
	}

	out, err := os.Create(newPath)
	if err != nil {
		print("create failed")
		log.Fatal(err)
	}
	defer out.Close()

	var c Converter
	switch format {
	case "jpeg", "jpg":
		{
			c = &jpgConverter{out, decoded}
		}
	case "png":
		{
			c = &pngConverter{out, decoded}
		}
	case "gif":
		{
			c = &gifConverter{out, decoded}
		}
	default:
		{
			log.Fatal("Can't generate converter: illegal format")
		}
	}

	if err := c.convimg(); err != nil {
		print("encode failed")
		log.Fatal(err)
	}

	os.Remove(path)
	return newPath
}
