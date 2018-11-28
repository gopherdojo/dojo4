package main

import (
	"os"

	"github.com/happylifetaka/dojo4/kadai3-1/happylifetaka/typinggame"
)

func main() {
	var t typinggame.TypingGame
	t.Start(os.Stdin)
}
