package main

import (
	"fmt"
	"time"
)

type Greeting struct {
	Clock Clock
}

type Clock interface {
	Now() time.Time
}

type clock struct {
	time time.Time
}

func (clock) Now() time.Time {
	return time.Now()
}

// モック用のストラクとをわざわざ作っている。
type mockClock struct{}

func (mockClock) Now() time.Time {
	return time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)
}

func main() {
	// g := Greeting{clock{}}
	g := Greeting{mockClock{}}
	fmt.Println(g.Clock.Now())
}
