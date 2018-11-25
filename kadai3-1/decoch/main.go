package main

import (
	"fmt"

	"github.com/decoch/dojo4/kadai3-1/decoch/score"
	"github.com/decoch/dojo4/kadai3-1/decoch/timer"
	"github.com/decoch/dojo4/kadai3-1/decoch/typing"
)

func main() {
	fmt.Println("### start typing game. ###")

	typeCh := typing.Words()
	timerCh := timer.NewCh(10)

	counter := new(score.Counter)

	for {
		select {
		case v1 := <-typeCh:
			if v1 {
				counter.Add(1)
			}
		case _ = <-timerCh:
			fmt.Printf("\n### Score: %d ###\n", counter.Value())
			fmt.Println("### end typing game. ###")
			return
		}
	}
}
