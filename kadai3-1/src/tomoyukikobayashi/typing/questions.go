// Package typing このパッケージはタイピングゲームに関するロジックとデータを格納します
package typing

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	yamler "gopkg.in/yaml.v2"
)

const (
	wordFile = "words.yaml"
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
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func (q *constQuestioner) GetNewWord(level int) string {
	rand := random(1, len(q.qs[level]))
	// HACK ほんとはmap okを見た方がいいけど、省略
	return q.qs[level][rand]
}

// Yaml HACK 使用しているパッケージの使用でしょうがなくpublicにしている
type Yaml struct {
	Level1 []string `yaml:"Level1,flow"`
	Level2 []string `yaml:"Level2,flow"`
	Level3 []string `yaml:"Level3,flow"`
}

// TOOD NewGameでinterfaceとio.Reader渡してやる方が、Questionerが汎用になる
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

	// TODO レベルは今の所少ないのでとりあえずベタがき
	// yamlの構成工夫 or interfaceとしてレベル別に取る関数生やして動的に撮れる方が良い
	q.qs[1] = y.Level1
	q.qs[2] = y.Level2
	q.qs[3] = y.Level3

	return nil
}
