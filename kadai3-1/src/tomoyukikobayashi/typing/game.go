package typing

import (
	"bufio"
	"context"
	"io"
	"time"
)

// Game タイピングゲームを制御するロジックを返す
type Game interface {
	Run(int, io.Reader) (question <-chan string,
		answer <-chan string, result <-chan [2]int)
}

type constGame struct {
	Questioner
	// NOTE コンテキストに状態格納しようとしたがコピーして作り直すから、途中で書き換えてもルーチンをまたいでシェアされない模様
	// 結局共有変数に書いていることになるので、脆弱感がすごい
	clear       int
	miss        int
	currentWord string
}

// NewGame Gameのコンストラクタです
func NewGame() (Game, error) {
	q, err := NewQuestioner()
	if err != nil {
		return nil, err
	}
	return &constGame{Questioner: q}, nil
}

// Run ゲームを開始する
func (c *constGame) Run(timeout int, input io.Reader) (<-chan string,
	<-chan string, <-chan [2]int) {
	sc := bufio.NewScanner(input)

	qCh := make(chan string)
	aCh := make(chan string)
	routines := 2
	rCh := make(chan [2]int, routines)

	bc := context.Background()
	tc := time.Duration(timeout) * time.Second
	ctx, _ := context.WithTimeout(bc, tc)
	// NOTE defer cancel() defer cancelするとDone条件閉じてる

	// TODO 読みづらいので関数に切り出す
	go func() {
		for {
			word := c.GetNewWord(c.nextLevel())
			qCh <- word
			c.currentWord = word
			select {
			case <-ctx.Done():
				// TODO once.DO が使えるかも
				rCh <- [2]int{c.clear, c.miss}
				close(qCh)
				return
			default:
				//do nothing
			}
		}
	}()

	// 読み込みしてaChに送る
	go func() {
		for {
			if sc.Scan() {
				ans := sc.Text()
				aCh <- ans
				if c.isCorrect(ans) {
					c.clear = c.clear + 1
				} else {
					c.miss = c.miss + 1
				}
			}

			select {
			case <-ctx.Done():
				rCh <- [2]int{c.clear, c.miss}
				close(aCh)
				return
			default:
				//do nothing
			}
		}
	}()

	return qCh, aCh, rCh
}

// HACK 成功した回数に応じて、使う語彙のレベルを決める。ここは決め打ちで書いてる
func (c *constGame) nextLevel() int {
	if c.clear < 10 {
		return 1
	}
	if c.clear < 20 {
		return 2
	}
	return 3
}

func (c *constGame) isCorrect(word string) bool {
	return c.currentWord == word
}
