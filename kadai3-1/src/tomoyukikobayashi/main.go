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

	"tomoyukikobayashi/typing"
)

// CLIのExitコード
const (
	ExitSuccess     = 0
	ExitError       = 1
	ExitInvalidArgs = 2
)

// Exitしてもテスト落ちない操作するようにエイリアスにしている
var exit = os.Exit

// CLI テストしやすいようにCLIの出力先を差し替えられるようにしている
type CLI struct {
	inStream  io.Reader
	outStream io.Writer
	errStream io.Writer
}

// CLIツールのエントリーポイント
func main() {
	cli := &CLI{inStream: os.Stdin, outStream: os.Stdout, errStream: os.Stderr}
	exit(cli.Run(os.Args))
}

// Run テスト用に実行ロジックを切り出した内容
func (c *CLI) Run(args []string) int {

	flags := flag.NewFlagSet("typing", flag.ContinueOnError)
	flags.SetOutput(c.errStream)

	var t int
	flags.IntVar(&t, "t", 30, "time to play (second) default=30s")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitInvalidArgs
	}

	game, err := typing.NewGame(t, c.inStream)
	if err != nil {
		fmt.Fprintf(c.outStream, "failed to initizalize game %v", err)
		return ExitError
	}

	// TOOD 長いことテスト固めたくないので外から与えるようにする
	fmt.Fprintf(c.outStream, "start game %d sec\n", t)
	qCh, aCh, rCh := game.Run()

	for {
		q, progress := <-qCh
		fmt.Fprintf(c.outStream, "%v\n", q)
		if !progress {
			break
		}

		fmt.Fprintf(c.outStream, ">")
		a, progress := <-aCh
		fmt.Fprintf(c.outStream, "%v\n", a)
		if !progress {
			break
		}
	}

	r := <-rCh
	fmt.Fprintf(c.outStream, "clear %v miss %v\n", r[0], r[1])

	return ExitSuccess
}
