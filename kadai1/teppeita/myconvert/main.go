// Package myconvert is a command to convert image extension
package main

import (
	"flag"

	"github.com/teppeita/dojo4/kadai1/teppeita/myconvert/cmd"
)

var (
	from string
	to   string
)

func init() {
	flag.StringVar(&from, "from", "jpeg", "")
	flag.StringVar(&to, "to", "png", "")
}

func main() {
	flag.Parse()
	cmd.Convert(flag.Arg(0), from, to)
}
