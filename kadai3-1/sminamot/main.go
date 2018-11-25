package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

var qs = []string{
	"hoge",
	"fuga",
	"foo",
	"bar",
	"piyo",
}

func main() {
	tCh := limit(10 * time.Second)
	iCh := input(os.Stdin)
	rand.Seed(time.Now().UnixNano())
	aNum := 0
	qNum := 0
loop:
	for {
		q := getQuestion()
		fmt.Printf("[q: %s]\n", q)
		fmt.Print("> ")
		select {
		case m := <-iCh:
			qNum++
			if m == q {
				aNum++
				fmt.Print("matched!")
			} else {
				fmt.Print("unmatched!")
			}
			fmt.Printf(" (%d/%d)\n", aNum, qNum)
		case <-tCh:
			break loop
		}
	}
	fmt.Println()
	fmt.Printf("your score: %d\n", aNum)
}

func getQuestion() string {
	return qs[rand.Intn(len(qs))]
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

func limit(d time.Duration) <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		defer close(ch)
		<-time.After(d)
	}()
	return ch
}
