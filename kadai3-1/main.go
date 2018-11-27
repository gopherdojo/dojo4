package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
	"context"
	"time"
	"math/rand"
)

var gameString = []string{"apple","banana","peach","strawberry","cherry","watermelon","pineapple","grape"}
var totalGame = 0
var totalScore = 0

func main(){
	bc := context.Background()
	t := 30*time.Second
	ctx,cancel := context.WithTimeout(bc,t)
	defer cancel()

	ch := input(os.Stdin,ctx)
LOOP:
	for{
		fmt.Println(">")
		n := rand.Intn(len(gameString))
		randString := gameString[n]
		totalGame ++
		fmt.Println(randString)
		select {
		case <-ctx.Done():
			fmt.Println("time up")
			break LOOP
		default:
			typedString := <-ch
			if randString == typedString{
				fmt.Println("ok")
				totalScore ++
			}
		}
	}
	fmt.Printf("total score %d/%d",totalScore,totalGame)

}

func input(r io.Reader,ctx context.Context) <-chan string{
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan(){
			select{
			case <- ctx.Done():
				close(ch)
				return
			case ch <- s.Text():
			}
		}
	}()
	return ch
}