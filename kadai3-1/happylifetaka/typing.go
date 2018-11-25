package main

import (
	"github.com/happylifetaka/dojo4/kadai3-1/happylifetaka/typinggame"
	"os"
)

func main() {
	var t typinggame.TypingGame
	t.Start(os.Stdin)
}