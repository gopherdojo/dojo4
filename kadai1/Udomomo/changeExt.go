package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Conf struct {
	From string
	To   string
}

func SetConf(from, to string) Conf {
	return Conf{
		From: from,
		To:   to,
	}
}

//カレントディレクトリ: os.Getwd()
//名前を変える: os.Rename()
//再帰探索する
//ユーザ定義型を作る: fromとtoをstructにまとめ、changeExtメソッドを関連づけてみる

// rootDirにあるファイルの一覧を探索。
// ディレクトリがあれば再帰処理する。
func SearchFile(rootDir string, conf Conf) {
	files, err := ioutil.ReadDir(rootDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		path := rootDir + "/" + file.Name()
		if file.IsDir() {
			SearchFile(path, conf)
			continue
		}

		willChange, newPath := generateNewExt(path, conf)
		if willChange == false {
			continue
		}

		convFile(path, newPath)

	}
}

func generateNewExt(path string, conf Conf) (willConv bool, newPath string) {
	ext := filepath.Ext(path)

	//変換したい拡張子でなければ何もしない
	if ext != conf.From {
		return false, path
	}

	//変換する必要がなければ何もしない
	if conf.From == conf.To {
		return false, path
	}

	return true, path[:len(path)-len(ext)] + conf.To
}

func convFile(path, newPath string) {
	err := os.Rename(path, newPath)
	// 変換に失敗した場合も、エラーを出力するだけで次のファイルに移る
	if err != nil {
		println(err)
	}
}
