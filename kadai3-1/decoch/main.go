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

	gameCh := make(chan struct{})
	timerCh := make(chan struct{})

	go startGame(gameCh, timerCh)

	time.Sleep(10 * time.Second)
	close(timerCh)

	<-gameCh // wait until game end.
	fmt.Println("end typing game.")
}

func startGame(gameCh, timerCh chan struct{}) {
	result := 0

	go func() {
		<-timerCh
		fmt.Printf("score: %d\n", result)
		close(gameCh)
	}()

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
			result++
		} else {
			fmt.Printf("### WRONG ###\n\n")
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
