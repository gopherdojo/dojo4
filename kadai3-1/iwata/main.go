package main

import (
	"log"
	"os"

	"github.com/gopherdojo/dojo4/kadai3-1/iwata/cmdparser"
	"github.com/gopherdojo/dojo4/kadai3-1/iwata/gameplayer"
	"github.com/gopherdojo/dojo4/kadai3-1/iwata/questions"
)

func main() {
	p := cmdparser.New(os.Stderr)
	c, err := p.Parse(os.Args)
	if err != nil {
		log.Fatalf("Failed to parse command options: %+v", err)
	}

	ql, err := questions.Parse(c.TxtPath)
	if err != nil {
		log.Fatalf("Failed to parse text file: %+v", err)
	}

	g := gameplayer.NewGame(os.Stdout, os.Stdin, ql)
	s, err := g.Play(c.Timeout)
	if err != nil {
		log.Fatalf("Failed to play typing: %+v", err)
	}

	s.Display()
}
