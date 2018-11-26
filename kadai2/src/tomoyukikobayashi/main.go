/*
Pacakge main is the entry point of this project.
This mainly provides interaction logics and parameters
used in CLI intrerfaces.
*/
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"tomoyukikobayashi/file"
	"tomoyukikobayashi/imconv"
)

// CLIのExitコード
const (
	ExitSuccess     = 0
	ExitInvalidArgs = iota
	ExitError       = iota
)

// Exitしてもテスト落ちない操作するようにエイリアスにしている
var exit = os.Exit

// CLI テストしやすいようにCLIの出力先を差し替えられるようにしている
type CLI struct {
	outStream io.Writer
	errStream io.Writer
}

// CLIツールのエントリーポイント
func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	exit(cli.Run(os.Args))
}

// Run テスト用に実行ロジックを切り出した内容
func (c *CLI) Run(args []string) int {
	flags := flag.NewFlagSet("imconv", flag.ContinueOnError)
	flags.SetOutput(c.errStream)

	tmpf, tmpt := "", ""
	f, t := &tmpf, &tmpt
	flags.StringVar(t, "f", "jpg", "input image file format")
	flags.StringVar(f, "t", "png", "output image file format")

	flags.Usage = func() {
		fmt.Fprintf(c.errStream, "Usage: %s [OPTIONS] dirname \n", os.Args[0])
		flags.PrintDefaults()
		fmt.Fprintf(c.errStream, "Supported formats are %v \n", imconv.SupportedExtensions())
	}

	if err := flags.Parse(args[1:]); err != nil {
		return ExitInvalidArgs
	}

	dir := flags.Arg(0)

	// 引数のバリデーション
	if len(dir) < 1 {
		// stream設定しているおかげで、こいつも指定したstreamに書いているよう
		flags.Usage()
		return ExitInvalidArgs
	}

	if !imconv.Supported(*f) {
		fmt.Fprintf(c.errStream, "supported formats are %v \n", imconv.SupportedExtensions())
		return ExitInvalidArgs
	}

	if !imconv.Supported(*t) {
		fmt.Fprintf(c.errStream, "supported formats are %v \n", imconv.SupportedExtensions())
		return ExitInvalidArgs
	}

	if *f == *t {
		fmt.Fprintf(c.errStream, "input format and output format are the same \n")
		return ExitInvalidArgs
	}

	// 条件にマッチするファイルパスを検索
	paths, err := file.Find(dir, imconv.GetFormatThesaurus(*f))
	if err != nil {
		fmt.Fprintf(c.errStream, "%v", err)
		return ExitError
	}

	if len(paths) < 1 {
		fmt.Fprintf(c.outStream, "no files matched")
		return ExitSuccess
	}

	// 画像を変換する
	// TOOD 画像が大量になる場合はgoroutineで並列処理しても良さそう
	for _, path := range paths {
		fmt.Fprintf(c.outStream, "src:"+path+"\n")
		newPath, err := imconv.Convert(path, *f, *t)
		if err != nil {
			fmt.Fprintf(c.errStream, "%v", err)
			return ExitError
		}
		fmt.Fprintf(c.outStream, "dst:"+newPath+"\n")
	}

	return ExitSuccess
}
