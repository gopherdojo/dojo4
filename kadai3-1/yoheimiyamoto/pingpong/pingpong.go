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
	inputDone := make(chan struct{}, 1)
	input(r, inputCh, inputDone)

	var s score

	fmt.Fprintln(w, "ゲームスタート！制限時間は5秒！")
GAME:
	for {
		// wordsをすべて回答したらゲームを終了させる
		if s.count() == (len(words) - 1) {
			break GAME
		}

		word := words[s.count()]
		fmt.Fprintln(w, word)

		select {
		case i := <-inputCh:
			if i == word {
				fmt.Fprintln(w, "正解！")
				s.addCorrect()
			} else {
				fmt.Fprintln(w, "不正解！")
				s.addIncorrect()
			}
		case <-time.After(5 * time.Second):
			break GAME
		}
	}
	close(inputDone) // inputのチャネルを閉じる
	fmt.Fprintln(w, "タイムアウト。ゲーム終了")
	fmt.Println(s.Result())
}

// 標準入力から受け取ったテキストをchチャネルに送信。
func input(r io.Reader, ch chan<- string, done <-chan struct{}) {
	scanner := bufio.NewScanner(r)
	go func() {
		for scanner.Scan() {
			t := scanner.Text()
			ch <- t
		}
	}()

	// ゲームが終了した場合、inputも終了させる。
	go func() {
		select {
		case <-done:
			return
		}
	}()

	// 上記を以下のようにgoroutine使わずに記述するとゲームがフリーズしてしまう理由が理解できていないです。
	// select {
	// case <-done:
	// 	return
	// }
}
