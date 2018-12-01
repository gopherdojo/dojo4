package typing

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Words() (<-chan bool, error) {
	inputCh := make(chan bool)

	words, err := getWords()
	if err != nil {
		return nil, err
	}

	go func() {
		stdin := bufio.NewScanner(os.Stdin)
		rand.Seed(time.Now().UnixNano())

		for {
			position := rand.Intn(len(words))
			word := words[position]

			fmt.Println(word) // show expected word.

			stdin.Scan()
			input := stdin.Text()
			if word == input {
				fmt.Printf("### GOT IT!!!! ###\n\n")
				inputCh <- true
			} else {
				fmt.Printf("### WRONG ###\n\n")
				inputCh <- false
			}
		}
		close(inputCh)
	}()

	return inputCh, nil
}

func getWords() ([]string, error) {
	data, err := ioutil.ReadFile("./data/word.txt")
	if err != nil {
		return nil, err
	}
	samples := strings.Split(strings.TrimSpace(string(data)), "\n")
	return samples, nil
}
