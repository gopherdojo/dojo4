/*
Overview
このコードはGopher道場1回目の課題のソースコードです。
画像を変換するプログラムです。
Usage
go buildをし、バイナリを実行してください。
./yasu some_directory でsome_directoryに指定されたdirectory内の
全ての画像ファイルをpngファイルに変換します。
Options
バイナリを実行する際にオプションを指定する事が可能です。
to オプションを指定すると変換先の画像フォーマットを変更できます。
次のように指定するとgifフォーマットへと変換します。
./yasu -to gif some_directory
*/

package main

import (
	"dojo4/kadai1/yasu/converter"
	"flag"
	"io/ioutil"
	"os"
	"sync"
)

// オプション保持用変数
var (
	//before_format string
	afterFormat string // 変換先format
)

var wg = sync.WaitGroup{}

func main() {
	bindFlags()
	directory := os.Args[len(os.Args)-1]
	if err = analyzeFiles(directory); err != nil {
		log.Println(err)
	}
}

func bindFlags() {
	//flag.StringVar(&before_format, "from", "jpeg", "File format")
	flag.StringVar(&afterFormat, "to", "png", "Converted file format")
	flag.Parse()
}

func analyzeFiles(directory string) error {
	files, err := ioutil.ReadDir(directory) // Get all file information in directory
	if err != nil {
		return err
	}
	for _, fileInfo := range files {
		file, err := os.Open(directory + "/" + fileInfo.Name()) // Read file
		if err != nil {
			return err
		}
		wg.Add(1)
		go func(file *os.File) {
			if err := converter.ConvertImg(file, directory, afterFormat); err != nil {
				return err
			}
			wg.Done()
		}(file)
	}
	wg.Wait()
}
