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
	ch := make(chan string)
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
	ch := make(chan string)

	go func() {
		defer close(ch)
		sc := bufio.NewScanner(in)
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			if sc.Scan() {
				ch <- sc.Text()
			}
		}
	}()

	return ch
}
