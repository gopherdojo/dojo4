package pingpong

import (
	"bufio"
	"fmt"
	"io"
	"time"
)

// Play ...
// words -> 出題するワード一覧
func Play(r io.Reader, w io.Writer, words []string) {
	inputCh := make(chan string)
	scoreCh := make(chan score)

	input(r, inputCh)

	fmt.Fprintln(w, "ゲームスタート！制限時間は5秒！")
	go game(w, words, inputCh, scoreCh)

	s := <-scoreCh
	fmt.Fprintln(w, "ゲーム終了")
	fmt.Fprintln(w, s.Result())
}

func game(w io.Writer, words []string, inputCh <-chan string, scoreCh chan<- score) {
	var s score

	go func() {
		select {
		case <-time.After(5 * time.Second):
			fmt.Fprintln(w, "タイムアウト！")
			scoreCh <- s
		}
	}()

	for {
		// wordsをすべて回答したらゲームを終了させる
		if s.count() == len(words) {
			scoreCh <- s
			break
		}
		word := words[s.count()]
		fmt.Fprintln(w, word)
		i := <-inputCh
		if i == word {
			fmt.Fprintln(w, "正解！")
			s.addCorrect()
		} else {
			fmt.Fprintln(w, "不正解！")
			s.addIncorrect()
		}
	}
}

// 標準入力から受け取ったテキストをchチャネルに送信。
func input(r io.Reader, ch chan<- string) {
	scanner := bufio.NewScanner(r)
	go func() {
		for scanner.Scan() {
			t := scanner.Text()
			ch <- t
		}
	}()
}
