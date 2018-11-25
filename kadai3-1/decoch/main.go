package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	fmt.Println("### typing game start. ###")

	score := 0

	ch1 := input()
	ch2 := timer()

	var m sync.Mutex
	select {
	case v1 := <-ch1:
		m.Lock()
		if v1 {
			score++
		}
		m.Unlock()
	case _ = <-ch2:
		m.Lock()
		fmt.Printf("\n### Score: %d ###\n", score)
		fmt.Println("### typing game end. ###")
		m.Unlock()
	}
}

func input() <-chan bool {
	inputCh := make(chan bool)
	go func() {
		words, err := getWords()
		if err != nil {
			log.Fatalln("cannot get words.", err)
		}

		stdin := bufio.NewScanner(os.Stdin)
		rand.Seed(time.Now().UnixNano())
		for {
			position := rand.Intn(len(words))
			word := words[position]
			fmt.Println(word)

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
	return inputCh
}

func timer() <-chan struct{} {
	timerCh := make(chan struct{})
	go func() {
		time.Sleep(2 * time.Second)
		timerCh <- struct{}{}
		close(timerCh)
	}()
	return timerCh
}

func getWords() ([]string, error) {
	data, err := readFile()
	if err != nil {
		return nil, err
	}
	samples := strings.Split(strings.TrimSpace(string(data)), "\n")
	return samples, nil
}

func readFile() ([]byte, error) {
	data, err := ioutil.ReadFile("./data/word.txt")
	return data, err
}
