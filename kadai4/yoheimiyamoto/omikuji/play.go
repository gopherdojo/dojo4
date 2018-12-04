package omikuji

import (
	"math/rand"
	"time"
)

func play() *Result {
	now := time.Now()
	t := getType(now)
	return &Result{t}
}

// 吉凶を取得（大吉,中吉,小吉,凶,大凶）
func getType(now time.Time) string {
	// 三が日は大吉にする
	if now.Month() == 1 {
		if now.Day() == 1 || now.Day() == 2 || now.Day() == 3 {
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
