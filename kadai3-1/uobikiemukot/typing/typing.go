package typing

import (
	"bufio"
	"context"
	"io"
	"math/rand"
	"time"
)

// GenWord generate words and return string channel
func GenWord(ctx context.Context) <-chan string {
	ch := make(chan string, 0)
	l := len(words)

	rand.Seed(time.Now().Unix())

	go func() {
		defer close(ch)
		for {
			select {
			case ch <- words[rand.Intn(l)]:
			case <-ctx.Done():
				return
			}
		}
	}()

	return ch
}

// ReadInput read io.Reader and return string channel
func ReadInput(ctx context.Context, in io.Reader) <-chan string {
	ch := make(chan string, 0)

	go func() {
		defer close(ch)
		sc := bufio.NewScanner(in)
		for sc.Scan() {
			select {
			case ch <- sc.Text():
			case <-ctx.Done():
				return
			}
		}
	}()

	return ch
}
