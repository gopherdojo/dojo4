package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	ExitCodeOK = iota
)

var questions = []string{"Golang", "Ruby", "Scala", "Java", "PHP", "JavaScript", "C", "C++"}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Run() int {
	score := 0
	done := time.After(5 * time.Second)

	for flag := true; flag; {
		result := make(chan bool)
		go run(result)

		select {
		case <-done:
			flag = false
		case r := <-result:
			if r {
				score++
			}
		}
	}

	fmt.Println("Time up!")
	fmt.Printf("Your score is %d\n", score)

	return ExitCodeOK
}

func run(channel chan<- bool) {
	defer close(channel)

	q := questions[rand.Intn(7)]
	fmt.Printf("Type it: %s\n", q)

	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	a := in.Text()

	isCorrect := q == a

	if isCorrect {
		fmt.Println("Correct")
	} else {
		fmt.Println("Wrong")
	}

	channel <- isCorrect
}
