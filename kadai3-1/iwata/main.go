package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	ch := make(chan os.Signal, 1)
	defer signal.Stop(ch)
	signal.Notify(ch,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	ctx, cancel := sigHandledContext(ch)
	defer cancel()

	g := gameplayer.NewGame(os.Stdout, os.Stdin, ql)
	s, err := g.Play(ctx, time.Second*time.Duration(c.Timeout))
	if err != nil {
		log.Fatalf("Failed to play typing: %+v", err)
	}

	fmt.Printf("\n\nSCORE\nCorrect:\t%d\nIn correct:\t%d\n", s.CorrectNum, s.InCorrectNum)
}

func sigHandledContext(ch <-chan os.Signal) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
	LOOP:
		for {
			select {
			case s := <-ch:
				switch s {
				case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
					cancel()
				}
			case <-ctx.Done():
				break LOOP
			}
		}
	}()

	return ctx, cancel
}
