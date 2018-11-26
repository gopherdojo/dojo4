package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gopherdojo/dojo4/kadai3-1/akuchii/game"
)

func main() {
	timeout := 5 * time.Second
	game := game.NewGame(os.Stdin, os.Stdout, getWords(), timeout)
	cnt := game.Start()
	fmt.Printf("\nfinish!\ncorrect count is %d\n", cnt)
	os.Exit(0)
}

func getWords() []string {
	return []string{"csharp", "python", "perl", "cplusplus", "ruby", "golang", "scala"}
}
