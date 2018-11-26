// Package cmdparser provides ...
package cmdparser

import (
	"flag"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

// Config is a configuration
type Config struct {
	Timeout int
	TxtPath string
}

// Parser is a struct to parse arguments
type Parser struct {
	stderr io.Writer
}

func New(stderr io.Writer) *Parser {
	return &Parser{stderr: stderr}
}

func (p *Parser) Parse(args []string) (*Config, error) {
	config := &Config{}
	flags := flag.NewFlagSet("cmdparser", flag.ContinueOnError)
	flags.SetOutput(p.stderr)
	flags.IntVar(&config.Timeout, "timeout", 15, "Timeout limitation for waiting your typings")
	flags.Usage = func() {
		fmt.Fprintf(p.stderr, `
Usage of %s:
   %s [OPTIONS] TXT
TXT
   Path to a text file for typing games that needs to be written with line breaks separated.
OPTIONS
`, args[0], args[0])
		flags.PrintDefaults()
	}

	if err := flags.Parse(args[1:]); err != nil {
		return nil, errors.Wrap(err, "Failed to parse arguments")
	}
	if flags.NArg() != 1 {
		return nil, errors.Errorf("Not support %d args, only one arg", flags.NArg())
	}
	config.TxtPath = flags.Arg(0)

	return config, nil
}
