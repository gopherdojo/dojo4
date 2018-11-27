package main

import (
	"flag"
	"fmt"
	"io"
	"runtime"
)

const (
	ExitCodeOK = iota
	ExitCodeError
	ExitCodeBadArgsError
	ExitCodeParseFlagsError
	ExitCodeInvalidFlagError
)

const name = "gget"

// CLI represents CLI interface for gget
type CLI struct {
	outStream, errStream io.Writer
}

// Run runs gget command
func (cli *CLI) Run(args []string) int {
	var parallel int

	flags := flag.NewFlagSet(name, flag.ContinueOnError)
	flags.Usage = func() {
		fmt.Fprint(cli.outStream, usage)
	}

	numCPU := runtime.NumCPU()
	flags.IntVar(&parallel, "parallel", numCPU, "")
	flags.IntVar(&parallel, "p", numCPU, "")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagsError
	}

	if parallel < 1 {
		fmt.Fprintf(cli.errStream, "Failed to set up gget: The number of parallels cannot be less than one\n")
		return ExitCodeInvalidFlagError
	}

	parsedArgs := flags.Args()
	if len(parsedArgs) != 1 {
		fmt.Fprintf(cli.errStream, "Invalid arguments: you need to set exactly one URL\n")
		return ExitCodeBadArgsError
	}

	request, err := NewRequest(parsedArgs[0], parallel)
	if err != nil {
		fmt.Fprintf(cli.errStream, "Error occurred while initializing a request: %s\n", err)
		return ExitCodeError
	}

	if err := request.Do(); err != nil {
		fmt.Fprintf(cli.errStream, "Error occurred while downloading the file: %s\n", err)
		return ExitCodeError
	}

	return ExitCodeOK
}

var usage = `Usage: gget [options...] URL

gget is a wget like command to download file, but downloads a file in parallel

OPTIONS:
  --parallel value, -p value  specifies the amount of parallelism (default: the number of CPU)
  --help, -h                  prints help

`
