package handler

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Fortune おみくじ機能を提供する
func Fortune(w http.ResponseWriter, r *http.Request) {
	f := fourtunes.Omikuji()
	// TOOD こいつはレスポンス毎ではなく、middlewareなどで横断的にやらせたい
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	f.WriteJSON(w)
}

// Fourtune おみくじ用データとロジックを格納する
type Fourtune struct {
	Luck    string `json:"luck"`
	Message string `json:"message"`
}

// WriteJSON 与えられたライターに自信をjson化したものを書き込む
func (f *Fourtune) WriteJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	if err := enc.Encode(f); err != nil {
		return err
	}
	return nil
}

const (
	// DAIKICHI 大吉
	DAIKICHI = 0
	// CYUKICHI 中吉
	CYUKICHI = iota
	// KICHI 吉
	KICHI = iota
	// SYOKICHI 小吉
	SYOKICHI = iota
	// KYO 凶
	KYO = iota
	// DAIKYO 大凶
	DAIKYO = iota
)

// Fourtunes 全てのおみくじデータを保持して、おみくじを引くロジックを提供
type Fourtunes struct {
	data map[int]Fourtune
	// 時間をモックできるようにClockを生やしている
	Clock
}

// Clock 任意の時間を返却するインターフェイス
type Clock interface {
	Now() time.Time
}

// HACK 絶対書き換えられないようにしたいが、閉じ込めて毎回コピーするのも今回はやりすぎなのでしない
var fourtunes = Fourtunes{
	data: map[int]Fourtune{
		DAIKICHI: Fourtune{Luck: "大吉", Message: "最高やでー"},
		CYUKICHI: Fourtune{Luck: "中吉", Message: "ついてんなー"},
		KICHI:    Fourtune{Luck: "吉", Message: "まずまずやね"},
		SYOKICHI: Fourtune{Luck: "小吉", Message: "ぼちぼちやね"},
		KYO:      Fourtune{Luck: "凶", Message: "気を落とすなよ"},
		DAIKYO:   Fourtune{Luck: "大凶", Message: "ウケるwwww"},
	},
}

// Omikuji ランダムにおみくじ結果を返す
func (fs *Fourtunes) Omikuji() Fourtune {
	// 正月期間は必ず大吉
	if fs.shoudBeHappy() {
		return fourtunes.data[DAIKICHI]
	}
	rand := rand.Intn(len(fourtunes.data))
	// fourtunesの要素のポインタ渡したくないので値で返してる
	return fourtunes.data[rand]
}

func (fs *Fourtunes) shoudBeHappy() bool {
	now := fs.now()
	if now.Month() != time.Month(1) {
		return false
	}
	if now.Day() > 3 {
		return false
	}
	return true
}

func (fs *Fourtunes) now() time.Time {
	if fs.Clock == nil {
		// TOOD デバッグ用 消す
		//return time.Date(2018, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
		return time.Now()
	}
	return fs.Clock.Now()
}
