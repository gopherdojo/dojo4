package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/gopherdojo/dojo4/kadai3-2/uobikiemukot/rget"
)

const (
	exitSuccess         = 0
	exitMissingArgument = 1
	exitFailure         = 2
)

func download(ctx context.Context, fp *os.File, url, rng string, size int64) <-chan error {
	ch := make(chan error)

	go func() {
		defer close(ch)
		var err error

		// ref: https://developer.mozilla.org/ja/docs/Web/HTTP/Headers/Accept-Ranges
		// Accept-Ranges: bytes => support GET with Range Header
		// Accept-Ranges: none  => not support GET with Range Header
		switch rng {
		case "bytes":
			err = rget.Parallel(ctx, fp, url, size)
		default:
			err = rget.Serial(ctx, fp, url, size)
		}

		ch <- err
	}()

	return ch
}

func main() {
	var err error

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: rget URL\n")
		os.Exit(exitMissingArgument)
	}
	url := os.Args[1]

	fp, err := ioutil.TempFile("", "rget-*")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(exitFailure)
	}
	defer os.Remove(fp.Name())

	size, rng, err := rget.Head(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(exitFailure)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	res := download(ctx, fp, url, rng, size)

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT)

loop:
	for {
		select {
		case err = <-res:
			fmt.Fprintln(os.Stderr, "break")
			break loop
		case <-sig:
			fmt.Fprintln(os.Stderr, "signal")
			cancel()
		}
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(exitFailure)
	}

	size, err = io.Copy(os.Stdout, fp)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(exitFailure)
	}
}
