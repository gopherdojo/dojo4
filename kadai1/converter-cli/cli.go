package main

import (
	"flag"
	"fmt"
	"io"

	"github.com/shuheiktgw/dojo4/kadai1/converter-cli/converter"
)

const (
	ExitCodeOK = iota
	ExitCodeExpectedError
	ExitCodeUnexpectedError
	ExitCodeBadArgs
	ExitCodeParseFlagsError
	ExitCodeInvalidFlagError
)

const Name = "converter-cli"

var extensions = [4]string{".gif", ".jpeg", ".jpg", ".png"}

// CLI represents CLI interface for converter
type CLI struct {
	outStream, errStream io.Writer
}

// Run executes `converter-cli` command and converts images's extension
func (cli *CLI) Run(args []string) int {
	var (
		from string
		to   string
	)

	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.Usage = func() {
		fmt.Fprint(cli.outStream, usage)
	}

	flags.StringVar(&from, "from", ".jpg", "")
	flags.StringVar(&from, "f", ".jpg", "")

	flags.StringVar(&to, "to", ".png", "")
	flags.StringVar(&to, "t", ".png", "")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagsError
	}

	if !validateExtensions(from) {
		fmt.Fprintf(cli.errStream, "Failed to set up converter-cli: invalid extension `%s` is given for --from flag\n"+
			"Please choose an extension from one of those: %v\n\n", from, extensions)
		return ExitCodeInvalidFlagError
	}

	if !validateExtensions(to) {
		fmt.Fprintf(cli.errStream, "Failed to set up converter-cli: invalid extension `%s` is given for --to flag\n"+
			"Please choose an extension from one of those: %v\n\n", to, extensions)
		return ExitCodeInvalidFlagError
	}

	path := flags.Args()
	if len(path) != 1 {
		fmt.Fprintf(cli.errStream, "Failed to set up converter-cli: invalid argument\n"+
			"Please specify the exact one path to a directly or a file\n\n")
		return ExitCodeBadArgs
	}

	files, err := converter.Convert(from, to, path[0])
	if err != nil {
		if _, ok := err.(converter.Handled); ok {
			fmt.Fprintf(cli.errStream, "Failed to execute converter-cli\n"+
				"%s\n\n", err)
			return ExitCodeExpectedError
		}

		fmt.Fprintf(cli.errStream, `converter-cli failed because of the following error.

%s

You might encounter a bug with converter-cli, so please report it to https://github.com/xxx/xxxx

`, err)
		return ExitCodeUnexpectedError
	}

	fmt.Fprintf(cli.outStream, "converter-cli successfully converted following files to `%s`.\n", to)
	fmt.Fprintf(cli.outStream, "%s\n\n", files)
	return ExitCodeOK
}

func validateExtensions(ext string) bool {
	for _, e := range extensions {
		if ext == e {
			return true
		}
	}

	return false
}

var usage = `Usage: converter-cli [options...] PATH

converter-cli is a command line tool to convert image extension

OPTIONS:
  --from value, -f value  specifies a image extension converted from (default: .jpg)
  --to value, -t value    specifies a image extension converted to (default: .png)
  --help, -h              prints out help

`
