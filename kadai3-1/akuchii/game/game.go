package game

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

// Game generates typing game
type Game struct {
	writer  io.Writer
	idx     int
	words   []string
	timeout int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewGame generates new Game instance
func NewGame(w io.Writer, words []string, timeout int) *Game {
	return &Game{writer: w, words: words, timeout: timeout}
}

// Start starts typing game
func (g Game) Start() int {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(g.timeout)*time.Second)
	defer cancel()

	ch := input(os.Stdin)
	cnt := 0
LOOP:
	for {
		g.setNewIdx()
		fmt.Fprintln(g.writer, "> "+g.word())
		select {
		case v := <-ch:
			if v == g.word() {
				cnt++
				fmt.Fprintln(g.writer, "correct!")
			} else {
				fmt.Fprintln(g.writer, "incorrect!")
			}
		case <-ctx.Done():
			break LOOP
		}
	}
	return cnt
}

func (g *Game) setNewIdx() {
	g.idx = rand.Intn(len(g.words))
}

func (g Game) word() string {
	return g.words[g.idx]
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}
