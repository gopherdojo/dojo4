package game

import (
	"fmt"
	"io"
	"math/rand"
	"time"
)

type Game struct {
	correct int
	time    int64
	word    string
}

func NewGame(time int64) *Game {
	return &Game{time: time}
}

func (g *Game) Decision(answer string) bool {
	return g.word == answer
}

func (g *Game) SetWord(words []string) {
	rand.Seed(time.Now().UnixNano())
	wordNum := rand.Intn(len(words))
	g.word = words[wordNum]
}

func (g *Game) Correct() {
	g.correct++
}

func (g *Game) End(writer io.Writer) {
	fmt.Fprintf(writer, "\nnumeber of correct answers is %d\n", g.correct)
}

func (g *Game) Question(writer io.Writer) {
	fmt.Fprintf(writer, "> %s\n", g.word)
}

func (g *Game) GetTimer() <-chan time.Time {
	return time.After(time.Duration(g.time) * time.Second)
}
