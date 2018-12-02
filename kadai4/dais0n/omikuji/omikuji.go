package omikuji

import (
	"errors"
	"math/rand"
	"time"
)

const (
	SuperLucky    = "大吉"
	LittleLucky   = "吉"
	Normal        = "中吉"
	LittleUnLucky = "小吉"
	UnLucky       = "凶"
)

type Omikuji struct {
	Clock  Clock
	Random Random
}

func (o *Omikuji) Do() (string, error) {
	if o.checkLuckyDay() {
		return SuperLucky, nil
	}
	luckyNum := o.mikuji()
	switch luckyNum {
	case 0:
		return SuperLucky, nil
	case 1:
		return LittleLucky, nil
	case 2:
		return Normal, nil
	case 3:
		return LittleUnLucky, nil
	case 4:
		return UnLucky, nil
	default:
		return "", errors.New("unexpected error")
	}
}

func (o *Omikuji) mikuji() int {
	if o.Random == nil {
		rand.Seed(time.Now().UnixNano())
		return rand.Intn(4)
	}
	return o.Random.Intn()
}

// CheckLuckyDay is
func (o *Omikuji) checkLuckyDay() bool {
	now := o.now()
	if now.Month() == 1 {
		if now.Day() == 1 || now.Day() == 2 || now.Day() == 3 {
			return true
		}
	}
	return false
}

func (o *Omikuji) now() time.Time {
	if o.Clock == nil {
		return time.Now()
	}
	return o.Clock.Now()
}
