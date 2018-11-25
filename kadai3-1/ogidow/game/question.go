package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Questioner struct {
	words []string
}

func BuildQuestion() *Questioner {
	path := "./game/words.txt"
	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File %s could not read: %v\n", path, err)
		os.Exit(1)
	}
	defer f.Close()

	words := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		words = append(words, scanner.Text())
	}
	questioner := &Questioner{words: words}
	return questioner
}

func (q *Questioner) getWord() string {
	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(q.words))
	return q.words[index]
}
