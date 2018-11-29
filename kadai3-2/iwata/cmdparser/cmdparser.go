// Package cmdparser provides ...
package cmdparser

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/pkg/errors"
)

// Config is a configuration
type Config struct {
	Parallel int
	Timeout  int
	Output   string
	URL      string
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
	flags.IntVar(&config.Parallel, "n", runtime.NumCPU(), "Parallel number")
	flags.StringVar(&config.Output, "o", "./", "Output directory")
	flags.IntVar(&config.Timeout, "timeout", 15, "Timeout for HTTP conntection")
	flags.Usage = func() {
		fmt.Fprintf(p.stderr, `
Usage of %s:
   %s [OPTIONS] URL
URL
   Download URL
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
	info, err := os.Stat(config.Output)
	if err != nil {
		return nil, errors.Wrap(err, "output dir is invalid")
	}
	if !info.IsDir() {
		return nil, errors.Errorf("%s is not directory", config.Output)
	}

	config.URL = flags.Arg(0)

	return config, nil
}
