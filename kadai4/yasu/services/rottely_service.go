package services

import (
	"errors"
	"math/rand"
	"time"
)

func Rottely() (string, error) {
	var randomInt int
	var fortune string
	var err error
	location, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(location)
	newYearsDay := time.Date(2018, 1, 1, 0, 0, 0, 0, location)
	endDay := time.Date(2018, 1, 4, 0, 0, 0, 0, location)
	if !now.Before(newYearsDay) && !now.After(endDay) {
		fortune = "大吉"
	} else {
		rand.Seed(time.Now().UnixNano())
		randomInt = rand.Intn(17)
		switch randomInt {
		case 0:
			fortune = "大大吉"
		case 1:
			fortune = "大吉"
		case 2:
			fortune = "向大吉"
		case 3:
			fortune = "末大吉"
		case 4:
			fortune = "吉凶末分末大吉"
		case 5:
			fortune = "吉"
		case 6:
			fortune = "中吉"
		case 7:
			fortune = "小吉"
		case 8:
			fortune = "後吉"
		case 9:
			fortune = "末吉"
		case 10:
			fortune = "吉凶不分末吉"
		case 11:
			fortune = "末凶相交末吉"
		case 12:
			fortune = "吉凶相半"
		case 13:
			fortune = "吉凶相央"
		case 14:
			fortune = "小吉後吉"
		case 15:
			fortune = "凶後吉"
		case 16:
			fortune = "凶後大吉"
		default:
			err = errors.New("Rottely error: Invalid integer")
		}
	}
	return fortune, err
}
