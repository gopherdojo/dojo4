package main

import "./game"

func main() {
	game := game.BuildGame()
	game.Run()
}
