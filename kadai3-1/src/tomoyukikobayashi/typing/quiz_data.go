package typing

import (
	"io"

	yamler "gopkg.in/yaml.v2"
)

// QuizData クイズデータを読み出す
type QuizData interface {
	MaxLevel() int
	WordsByLevel(int) []string
}

// QuizSource HACK 使用しているパッケージの使用でしょうがなくpublicにしている
type QuizSource struct {
	Level1 []string `yaml:"Level1,flow"`
	Level2 []string `yaml:"Level2,flow"`
	Level3 []string `yaml:"Level3,flow"`
}

// NewQuizData クイズデータを生成する
func NewQuizData(source io.Reader) (QuizData, error) {
	var s QuizSource
	if err := yamler.NewDecoder(source).Decode(&s); err != nil {
		return nil, err
	}

	return &s, nil
}

// MaxLevel クイズの最高難易度を返す
func (q *QuizSource) MaxLevel() int {
	// HACK 決め打ち
	return 3
}

// WordsByLevel 指定されたレベルの語彙を返す
func (q *QuizSource) WordsByLevel(level int) []string {
	// HACK
	switch level {
	case 1:
		return q.Level1
	case 2:
		return q.Level2
	case 3:
		return q.Level3
	}
	return nil
}
