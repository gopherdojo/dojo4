package main

import (
	"fmt"
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
	convert ConvImg
}
type ConvImg interface {
	Convert() error
}

// Run command
func (c *CLI) Run() int {
	if err := c.convert.Convert(); err != nil {
		fmt.Fprint(c.outStream, fmt.Sprintf("Error: %s", err))
		return ExitCodeError
	}
	return ExitCodeOK
}
