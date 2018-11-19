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

		pp, err := convFile(p, np, to)
		if err != nil {
			log.Fatal(err)
		}
		processedPaths = append(processedPaths, pp)
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

	return true
}

//convFile : 変換を実行する
func convFile(path, newPath, toExt string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return newPath, fmt.Errorf("Original file open failed. Path: %s", path)
	}
	defer file.Close()

	decoded, _, err := image.Decode(file)
	if err != nil {
		return newPath, fmt.Errorf("Original file decode failed. Path: %s", path)
	}

	out, err := os.Create(newPath)
	if err != nil {
		return newPath, fmt.Errorf("New empty file creation failed. Path: %s", newPath)
	}
	defer out.Close()

	var c Converter
	switch toExt {
	case ".jpeg", ".jpg":
		{
			c = &jpgConverter{out, decoded}
		}
	case ".png":
		{
			c = &pngConverter{out, decoded}
		}
	case ".gif":
		{
			c = &gifConverter{out, decoded}
		}
	default:
		{
			return newPath, fmt.Errorf("Can't convert to illegal format: %s", toExt)
		}
	}

	if err := c.convimg(); err != nil {
		return newPath, fmt.Errorf("Encode failed from %s to %s", path, newPath)
	}

	return newPath, nil
}
