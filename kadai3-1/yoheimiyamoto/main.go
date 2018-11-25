package main

import (
	"os"

	"github.com/YoheiMiyamoto/dojo4/kadai3-1/yoheimiyamoto/pingpong"
)

func main() {
	words := []string{"one", "two", "three", "four", "five"}
	pingpong.Play(os.Stdin, os.Stdout, words)

}
