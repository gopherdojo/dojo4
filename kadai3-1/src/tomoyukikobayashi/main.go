/*
Pacakge main is the entry point of this project.
This mainly provides interaction logics and parameters
used in CLI intrerfaces.
*/
package main

import (
	"fmt"
	"io"
	"os"

	"tomoyukikobayashi/typing"
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

	q, err := typing.NewQuestioner()
	fmt.Printf("%v %v", q, err)

	w, err := q.GetNewWord(1)
	fmt.Printf("%v %v", w, err)

	return ExitSuccess
}
