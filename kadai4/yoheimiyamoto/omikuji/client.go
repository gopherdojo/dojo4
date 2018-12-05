package omikuji

import (
	"math/rand"
	"time"
)

type client struct {
	now now
}

type now func() time.Time

func NewClient() *client {
	return &client{
		now(func() time.Time {
			return time.Now()
		}),
	}
}

func (c *client) play() *Result {
	t := getType(c.now())
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
