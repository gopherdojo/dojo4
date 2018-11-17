
package main

import (
	"flag"
	"github.com/dais0n/dojo4/kadai2/dais0n/convimg"
	"os"
)

func main() {
	// Command Options
	var from string
	var to string
	var path string
	flags := flag.NewFlagSet("conv", flag.ContinueOnError)
	flags.StringVar(&from,"f", "jpg", "from extension")
	flags.StringVar(&to, "t", "png", "to extension")
	flags.StringVar(&path, "p", "./images", "images dir path")

	if err := flags.Parse(os.Args[1:]); err != nil {
		os.Exit(1)
	}
	convert := convimg.NewConvert(from, to, path)
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr, convert: convert}

	os.Exit(cli.Run(os.Args))
}
