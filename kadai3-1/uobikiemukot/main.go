package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gopherdojo/dojo4/kadai3-1/uobikiemukot/typing"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	input := typing.ReadInput(ctx, os.Stdin)
	word := typing.GenWord(ctx)

	var correct, sum int

loop:
	for {
		// get first word
		want := <-word
		fmt.Println("current word:" + want)

		// wait user's input
		got := <-input
		fmt.Println("you typed:" + got)

		if want == got {
			correct++
		}
		sum++

		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break loop
		}
	}

	fmt.Printf("accuracy: %.2f %% (%d/%d)\n", float32(correct)/float32(sum)*100, correct, sum)
}
