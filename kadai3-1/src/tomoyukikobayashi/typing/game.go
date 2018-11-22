package typing

import (
	"bufio"
	"context"
	"io"
	"time"
)

// Game タイピングゲームを制御するロジックを返す
type Game interface {
	Run() (question <-chan string, answer <-chan string, result <-chan [2]int)
}

type constGame struct {
	Questioner
	input   io.Reader
	timeout int
	// NOTE コンテキストにゲーム状態格納しようとしたがコピーして作り直すから、途中で書き換えてもルーチンをまたいでシェアされない(親子関係で)
	// 結局共有変数に書いていることになるので、脆弱感がすごい
	clear       int
	miss        int
	currentWord string
}

// NewGame Gameのコンストラクタです
func NewGame(timeout int, input io.Reader) (Game, error) {
	q, err := NewQuestioner()
	if err != nil {
		return nil, err
	}
	return &constGame{Questioner: q, timeout: timeout, input: input}, nil
}

// Run ゲームを開始する
func (c *constGame) Run() (<-chan string, <-chan string, <-chan [2]int) {
	tc := time.Duration(c.timeout) * time.Second
	// NOTE defer cancel() cancelするとDone条件閉じてゲーム終わっちゃう
	ctx, _ := context.WithTimeout(context.Background(), tc)

	// TODO routine数にあわせてサイズ調整は死ねるのでonce.DO が使えるかも
	routines := 2
	rCh := make(chan [2]int, routines)

	qCh := c.question(ctx, rCh)
	aCh := c.answer(ctx, rCh)

	return qCh, aCh, rCh
}

// 問題をqChに送る
func (c *constGame) question(ctx context.Context, rCh chan<- [2]int) <-chan string {
	qCh := make(chan string)
	go func() {
		for {
			word := c.GetNewWord(c.nextLevel())
			qCh <- word
			c.currentWord = word
			select {
			case <-ctx.Done():
				rCh <- [2]int{c.clear, c.miss}
				close(qCh)
				return
			default:
				//do nothing
			}
		}
	}()
	return qCh
}

// 回答をストリームから読み込みしてaChに送る
func (c *constGame) answer(ctx context.Context, rCh chan<- [2]int) <-chan string {
	sc := bufio.NewScanner(c.input)
	aCh := make(chan string)
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
			// contextがtimeoutしたら結果を返却
			case <-ctx.Done():
				rCh <- [2]int{c.clear, c.miss}
				close(aCh)
				return
			default:
				//do nothing
			}
		}
	}()
	return aCh
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

// ワードの比較
func (c *constGame) isCorrect(word string) bool {
	return c.currentWord == word
}
