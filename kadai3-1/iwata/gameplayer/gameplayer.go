package gameplayer

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"time"

	"github.com/gopherdojo/dojo4/kadai3-1/iwata/questions"
)

type GamePlayer struct {
	w  io.Writer
	r  io.Reader
	ql questions.List
}

func NewGame(w io.Writer, r io.Reader, ql questions.List) *GamePlayer {
	return &GamePlayer{w: w, r: r, ql: ql}
}

func (p *GamePlayer) Play(timeout int) *Score {
	p.display("Start Typing Game!!")
	p.display(fmt.Sprintf("The time limit is %d seconds", timeout))

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	s := &Score{w: p.w}
	n := 1

	ch := p.readAnswer()
	defer close(ch)

GAMEEND:
	for {
		q := p.ql.Give()
		p.display(fmt.Sprintf("Q%d: %s", n, q.Quiz))
		n++

		select {
		case answer := <-ch:
			if q.Match(answer) {
				s.correct()
			} else {
				s.inCorrect()
			}
		case <-ctx.Done():
			p.display("This challenge has been time up!!!")
			break GAMEEND
		}
	}

	return s
}

func (p *GamePlayer) display(msg string) {
	fmt.Fprintln(p.w, msg)
}

func (p *GamePlayer) readAnswer() chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(p.r)
		for s.Scan() {
			ch <- s.Text()
		}
	}()
	return ch
}

type Score struct {
	w            io.Writer
	correctNum   int
	inCorrectNum int
}

func (s *Score) correct() {
	s.correctNum++
}

func (s *Score) inCorrect() {
	s.inCorrectNum++
}

func (s *Score) Display() {
	fmt.Fprintf(s.w, "Correct:\t%d\nIn correct:\t%d\n", s.correctNum, s.inCorrectNum)
}
