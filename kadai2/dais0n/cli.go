package main

import (
	"fmt"
	"github.com/dais0n/dojo4/kadai2/dais0n/convimg"
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
	convert *convimg.ConvImg
}

// Run command
func (c *CLI) Run(args []string) int {
	if err := c.convert.Convert(); err != nil {
		fmt.Fprint(c.outStream, fmt.Sprintf("Error: %s", err))
		return ExitCodeError
	}
	return ExitCodeOK
}
