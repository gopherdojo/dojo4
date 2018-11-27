package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan os.Signal)
	defer close(ch)
	signal.Notify(ch,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	handleSignal(ctx, cancel, ch)

	g := gameplayer.NewGame(os.Stdout, os.Stdin, ql)
	s, err := g.Play(ctx, c.Timeout)
	if err != nil {
		log.Fatalf("Failed to play typing: %+v", err)
	}

	s.Display()
}

func handleSignal(ctx context.Context, cancel context.CancelFunc, ch <-chan os.Signal) {
	go func() {
		for {
			s := <-ch
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				cancel()
				log.Println("Canceled")
			}
		}
	}()
}
