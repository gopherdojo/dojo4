package convimg

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

//SearchFile : searchFile関数が出力する変換後のパスを貯めておき、返り値として返す。
func SearchFile(rootDir, from, to string) []string {
	processingPaths := make([]string, 0)
	processingPaths = searchFile(rootDir, from, to, processingPaths)

	processedPaths := make([]string, 0)
	for _, p := range processingPaths {
		e := filepath.Ext(p)
		np := p[:len(p)-len(e)] + to
		processedPaths = append(processedPaths, convFile(p, np))
		os.Remove(p)
	}

	//os.Remove後の結果はユニットテストしづらいので、ここで簡易的に確認する
	if len(processingPaths) != len(processingPaths) {
		fmt.Println("len of conv results is wrong")
		os.Exit(1)
	}
	return processedPaths
}

//searchFile : rootDirにあるファイルの一覧を探索。ディレクトリがあれば再帰処理する。
func searchFile(rootDir, from, to string, processingPaths []string) []string {

	files, err := ioutil.ReadDir(rootDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		path := filepath.Join(rootDir, file.Name())
		if file.IsDir() {
			processingPaths = searchFile(path, from, to, processingPaths)
			continue
		}

		willConv := validateIfConvNeeded(path, from, to)
		if willConv == false {
			continue
		}

		processingPaths = append(processingPaths, path)
	}
	return processingPaths
}

//generateNewExt : 変換が必要な場合、変換後のパスを生成して返す
func validateIfConvNeeded(path, from, to string) bool {
	ext := filepath.Ext(path)

	//変換したい拡張子でなければ何もしない
	if ext != from {
		return false
	}

	//変換する必要がなければ何もしない
	if from == to {
		return false
	}

	return true //path[:len(path)-len(ext)] + to
}

//convFile : 変換を実行する
func convFile(path, newPath string) string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("open failed")
		log.Fatal(err)
	}
	defer file.Close()

	decoded, format, err := image.Decode(file)
	if err != nil {
		fmt.Println("decode failed")
		log.Fatal(err)
	}

	out, err := os.Create(newPath)
	if err != nil {
		fmt.Println("create failed")
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
		fmt.Println("encode failed")
		log.Fatal(err)
	}

	return newPath
}
