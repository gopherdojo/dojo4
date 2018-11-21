// Package typing このパッケージはタイピングゲームに関するロジックとデータを格納します
package typing

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"

	yamler "gopkg.in/yaml.v2"
)

const (
	wordFile = "words.yaml"
)

// Questioner 質問で使うワードを提供します
type Questioner interface {
	// GetNewWord 引数で与えた難易度に対して一つランダムにワードを返します
	GetNewWord(level int) (string, error)
}

type constQuestioner struct {
	qs map[int][]string
}

// NewQuestioner Questionerのコンストラクタです
func NewQuestioner() (Questioner, error) {
	q := &constQuestioner{
		qs: map[int][]string{},
	}
	if err := q.loadWords(); err != nil {
		return nil, err
	}
	return q, nil
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func (q *constQuestioner) GetNewWord(level int) string {
	rand := random(0, len(q.qs[level]))
	// HACK ほんとはmap okを見た方がいいけど、省略
	return "", q.qs[level][rand]
}

// Yaml HACK 使用しているパッケージの使用でしょうがなくpublicにしている
type Yaml struct {
	Level1 []string `yaml:"Level1,flow"`
	Level2 []string `yaml:"Level2,flow"`
	Level3 []string `yaml:"Level3,flow"`
}

// TOOD ファイル直呼びするんじゃなく、NewQuestionerでio.Reader渡すようにしておくとテストしやすいかな
func (q *constQuestioner) loadWords() error {
	cur, _ := os.Getwd()
	// yamlファイルから語彙リストを読み出す
	data, err := ioutil.ReadFile(filepath.Join(cur, wordFile))
	if err != nil {
		return err
	}

	var y Yaml
	err = yamler.Unmarshal([]byte(data), &y)
	if err != nil {
		return err
	}

	// HACK レベルは今の所少ないのでとりあえずベタがき。
	// yamlの構成工夫して動的に取れた方がいい
	q.qs[1] = y.Level1
	q.qs[2] = y.Level2
	q.qs[3] = y.Level3

	return nil
}
