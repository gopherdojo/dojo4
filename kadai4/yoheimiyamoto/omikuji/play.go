package omikuji

import (
	"math/rand"
	"time"
)

func Play() *Result {
	t := getType()
	return &Result{t}
}

func getType() string {
	// 三が日は大吉にする
	t := time.Now()
	if t.Month() == 1 {
		if t.Day() == 1 || t.Day() == 2 || t.Day() == 3 {
			return "大吉"
		}
	}

	// 三が日以外はランダム
	types := []string{
		"大吉",
		"中吉",
		"小吉",
		"凶",
		"大凶",
	}
	rand.Seed(time.Now().UnixNano())
	return types[rand.Intn(len(types))]
}
