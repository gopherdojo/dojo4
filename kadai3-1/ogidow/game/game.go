package game

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type GameConfig struct {
	timeout      time.Duration
	correctCount int
	q            *Questioner
	inputCh      <-chan string
	wg           sync.WaitGroup
}

func BuildGame() *GameConfig {
	return &GameConfig{
		timeout:      10 * time.Second,
		correctCount: 0,
		q:            BuildQuestion(),
		inputCh:      input(os.Stdin),
	}
}

func (gc *GameConfig) Run() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, gc.timeout)
	defer cancel()

	gc.wg.Add(1)
	go gc.play(ctx)
	gc.wg.Wait()
}

func (gc *GameConfig) play(ctx context.Context) {

	for {
		word := gc.q.getWord()
		fmt.Println(word)
		fmt.Print(">")
		ans := <-gc.inputCh
		if ans == word {
			fmt.Println("collect!")
			gc.correctCount++
		} else {
			fmt.Println("uncollect!")
		}

		select {
		case <-ctx.Done():
			fmt.Printf("%d times collect\n", gc.correctCount)
			gc.wg.Done()
			return
		default:
		}
	}
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
