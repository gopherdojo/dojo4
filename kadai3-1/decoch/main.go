package main

import (
	"fmt"
	"log"
	"time"

	"github.com/decoch/dojo4/kadai3-1/decoch/score"
	"github.com/decoch/dojo4/kadai3-1/decoch/typing"
)

func main() {
	fmt.Println("### start typing game. ###")

	typeCh, err := typing.Words()
	if err != nil {
		log.Fatal(err)
	}
	timerCh := time.After(10 * time.Second)

	counter := new(score.Counter)

	for {
		select {
		case v1 := <-typeCh:
			if v1 {
				counter.Add(1)
			}
		case <-timerCh:
			fmt.Printf("\n### Score: %d ###\n", counter.Value())
			fmt.Println("### end typing game. ###")
			return
		}
	}
}
