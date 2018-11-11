package main

import (
	"flag"
	"fmt"
	"github.com/dais0n/dojo4/kadai1/dais0n/convimg"
	"io"
)

const (
	// ExitCodeOK is successful completion
	ExitCodeOK = iota
	// ExitCodeError is fail completion
	ExitCodeError
)

// CLI has stream
type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) Run(args []string) int {
	// Command Options
	var from string
	var to string
	var path string
	flags := flag.NewFlagSet("c", flag.ContinueOnError)
	flags.SetOutput(c.errStream)

	flags.StringVar(&from,"f", "jpg", "from extension")
	flags.StringVar(&to, "t", "png", "to extension")
	flags.StringVar(&path, "p", "./images", "path")
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}
	// Convert
	convert := convimg.NewConvert(from, to, path)
	if err := convert.Convert(); err != nil {
		fmt.Fprint(c.outStream, fmt.Sprintf("Error: %s", err))
		return ExitCodeError
	}
	return ExitCodeOK
}