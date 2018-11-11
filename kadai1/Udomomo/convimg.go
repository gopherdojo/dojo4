package main

import (
	"image"
	_ "image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

//再帰探索する
//ユーザ定義型を作る: fromとtoをstructにまとめ、changeExtメソッドを関連づけてみる
//探索と変換をそれぞれ別パッケージにした方がよいかも

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

		convFileToPng(path, newPath)

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

func convFileToPng(path, newPath string) {
	println(path)
	file, err := os.Open(path)
	if err != nil {
		print("open failed")
		log.Fatal(err)
	}
	defer file.Close()

	decoded, _, err := image.Decode(file)
	if err != nil {
		print("decode failed")
		log.Fatal(err)
	}

	out, err := os.Create(newPath)
	if err != nil {
		print("create failed")
		log.Fatal(err)
	}
	defer out.Close()

	if err := png.Encode(out, decoded); err != nil {
		print("encode failed")
		log.Fatal(err)
	}

	os.Remove(path)
}
