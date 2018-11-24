package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/dais0n/dojo4/kadai3-1/dais0n/game"
)

const WORDS_FILE = "words.txt"

func main() {
	words, err := readWords()
	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}
	ch := input(os.Stdin)
	g := game.NewGame(10)
	timer := g.GetTimer()
	for {
		g.SetWord(words)
		g.Question(os.Stdout)
		select {
		case v, ok := <-ch:
			if !ok {
				g.End(os.Stderr)
				os.Exit(0)
			}
			if g.Decision(v) {
				g.Correct()
			}
		case <-timer:
			g.End(os.Stdout)
			os.Exit(0)
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

func readWords() ([]string, error) {
	var words []string
	f, err := os.Open(WORDS_FILE)

	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	// get file size
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, nil
}
