/*
Pacakge main is the entry point of this project.
This mainly provides interaction logics and parameters
used in CLI intrerfaces.
*/
package main

import (
	"flag"
	"fmt"
	"os"

	"file"
	"imconv"
)

// CLIのExitコード
const (
	ExitSuccess = 0
	ExitError   = iota
)

// CLIのオプションパラメタ
var (
	f = flag.String("f", "jpg", "input image file format")
	t = flag.String("t", "png", "output image file format")
)

func init() {
	// 引数がおかしいときはUsageを表示する
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] dirname \n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "Supported formats are %v \n", imconv.SupportedExtensions())
	}
}

// CLIツールのエントリーポイント
func main() {
	flag.Parse()
	dir := flag.Arg(0)

	// 引数のバリデーション
	if len(dir) < 1 {
		flag.Usage()
		os.Exit(ExitError)
	}

	if !imconv.Supported(*f) {
		fmt.Fprintf(os.Stderr, "Supported formats are %v \n", imconv.SupportedExtensions())
		os.Exit(ExitError)
	}

	if !imconv.Supported(*t) {
		fmt.Fprintf(os.Stderr, "Supported formats are %v \n", imconv.SupportedExtensions())
		os.Exit(ExitError)
	}

	if *f == *t {
		fmt.Fprintf(os.Stderr, "input format and output format are the same \n")
		os.Exit(ExitError)
	}

	// 条件にマッチするファイルパスを検索
	paths, err := file.Find(dir, imconv.GetFormatThesaurus(*f))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(ExitError)
	}

	if len(paths) < 1 {
		fmt.Println("no files matched")
		os.Exit(ExitSuccess)
	}

	// 画像を変換する
	// TOOD 画像が大量になる場合はgoroutineで並列処理しても良さそう
	for _, path := range paths {
		fmt.Println("src:", path)
		newPath, err := imconv.Convert(path, *f, *t)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(ExitError)
		}
		fmt.Println("dst:", newPath)
	}

	os.Exit(ExitSuccess)
}
