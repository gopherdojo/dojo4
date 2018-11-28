package gameplayer

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"time"

	"github.com/gopherdojo/dojo4/kadai3-1/iwata/questions"
	"github.com/pkg/errors"
)

type GamePlayer struct {
	w  io.Writer
	r  io.Reader
	ql questions.List
}

func NewGame(w io.Writer, r io.Reader, ql questions.List) *GamePlayer {
	return &GamePlayer{w: w, r: r, ql: ql}
}

func (p *GamePlayer) Play(ctx context.Context, timeout time.Duration) (*Score, error) {
	p.display("Start Typing Game!!")
	p.display(fmt.Sprintf("The time limit is %d seconds\n", timeout))

	s := &Score{}
	n := 1
	ctxWT, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	ch, cherr := p.readAnswer(ctxWT)
	defer close(ch)
	defer close(cherr)

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
		case err := <-cherr:
			return nil, err
		case <-ctxWT.Done():
			p.display("\nThis challenge has been time up!!!")
			break GAMEEND
		}
	}

	return s, nil
}

func (p *GamePlayer) display(msg string) {
	fmt.Fprintln(p.w, msg)
}

func (p *GamePlayer) readAnswer(ctx context.Context) (chan string, chan error) {
	ch := make(chan string)
	cherr := make(chan error)
	go func() {
		s := bufio.NewScanner(p.r)
		for s.Scan() {
			select {
			case <-ctx.Done():
				break
			default:
				ch <- s.Text()
			}
		}
		if err := s.Err(); err != nil {
			cherr <- errors.Wrap(err, "Failed to read from standard input")
		}
	}()

	return ch, cherr
}

type Score struct {
	CorrectNum   int
	InCorrectNum int
}

func (s *Score) correct() {
	s.CorrectNum++
}

func (s *Score) inCorrect() {
	s.InCorrectNum++
}
