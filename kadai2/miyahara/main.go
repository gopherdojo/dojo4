package main

import (
	"flag"
	"fmt"
	"github.com/gopherdojo/dojo4/kadai2/miyahara/converter"
	"os"
)

var (
	inputFile  = flag.String("i", "jpg", "input image type")
	outputFile = flag.String("o", "png", "output image type")
)

func usage() {
	fmt.Fprint(os.Stderr, "usage: conv -i=[input file type] -o=[output file type] [target directory]\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprint(os.Stderr, "error: specify target directory\n")
		os.Exit(2)
	}

	converter := converter.New(*inputFile, *outputFile, flag.Arg(0))
	os.Exit(converter.Run())
}
