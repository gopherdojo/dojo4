package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("typing game start.")

	stopCh := make(chan struct{})
	doneCh := make(chan struct{})

	go startGame(stopCh, doneCh)

	time.Sleep(10 * time.Second)
	close(stopCh)

	<-doneCh // wait until game end.
	fmt.Println("end typing game.")
}

func startGame(stopCh, doneCh chan struct{}) {
	defer func() {
		close(doneCh) // notify game end.
	}()

	result := 0
	words, err := getWords()
	if err != nil {
		log.Fatalln("cannot get words.", err)
	}

	stdin := bufio.NewScanner(os.Stdin)
	for {
		rand.Seed(time.Now().UnixNano())
		position := rand.Intn(len(words))

		question := words[position]
		fmt.Println(question)

		stdin.Scan()
		input := stdin.Text()
		if question == input {
			fmt.Printf("### GOT IT!!!! ###\n\n")
			result++
		} else {
			fmt.Printf("### WRONG ###\n\n")
		}

		select {
		case <-stopCh:
			fmt.Printf("score: %d\n", result)
			return
		default:
			// loop
		}
	}
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
