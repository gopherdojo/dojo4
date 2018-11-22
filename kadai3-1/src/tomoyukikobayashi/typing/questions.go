// Package typing このパッケージはタイピングゲームに関するロジックとデータを格納します
package typing

import (
	"math/rand"
	"time"
)

// Questioner 質問で使うワードを提供します
type Questioner interface {
	// GetNewWord 引数で与えた難易度に対して一つランダムにワードを返します
	GetNewWord(level int) string
}

type constQuestioner struct {
	qs map[int][]string
}

// NewQuestioner Questionerのコンストラクタ
func NewQuestioner(data QuizData) Questioner {
	qs := map[int][]string{}
	for i := 1; i <= data.MaxLevel(); i++ {
		qs[i] = data.WordsByLevel(i)
	}
	q := &constQuestioner{
		qs: qs,
	}
	return q
}

func (q *constQuestioner) GetNewWord(level int) string {
	rand.Seed(time.Now().UnixNano())
	rand := rand.Intn(len(q.qs[level]))
	// HACK ほんとはmap okを見た方がいいけど、省略
	return q.qs[level][rand]
}
