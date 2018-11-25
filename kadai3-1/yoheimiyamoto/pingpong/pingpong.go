package pingpong

import (
	"bufio"
	"fmt"
	"os"
)

// Play ...
// words -> 出題するワード一覧
func Play(stdin, stdout *os.File, words []string) {
	fmt.Println("ゲームスタート！\nゲームを途中で終了する場合はexitを入力してください。")
	inputCh := make(chan string)
	done := make(chan struct{})
	var r result
	input(os.Stdin, inputCh, done)
	var count int
	for {
		word := words[count]
		fmt.Println(word)

		select {
		case a := <-inputCh:
			if a == word {
				fmt.Fprintln(stdout, "正解！")
				r.addCorrect()
			} else {
				fmt.Fprintln(stdout, "不正解！")
				r.addIncorrect()
			}
			if count == (len(words) - 1) {
				fmt.Println(r)
				os.Exit(0)
			}
			count++
		case <-done:
			fmt.Fprintln(stdout, "終了します。")
			fmt.Println(r)
			os.Exit(0)
		}
	}
}

// 標準入力から受け取ったテキストをchチャネルに送信。
// 標準入力から「exit」というワードを受け取った場合は、doneチャネルをcloseする。
func input(stdin *os.File, ch chan<- string, done chan<- struct{}) {
	scanner := bufio.NewScanner(stdin)
	go func() {
		for scanner.Scan() {
			t := scanner.Text()
			if t == "exit" {
				close(done)
			}
			ch <- t
		}
	}()
}
