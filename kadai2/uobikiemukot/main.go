package main

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/gopherdojo/dojo4/kadai2/uobikiemukot/imgconv"
)

func main() {
	ifmt := flag.String("i", "jpg", "input image format (default: jpg)")
	ofmt := flag.String("o", "png", "output image format (default: png)")
	flag.Parse()

	if len(flag.Args()) < 1 {
		log.Fatal("usage: imgconv -i INPUT_FORMAT -o OUTPUT_FORMAT DIR")
	}

	if *ifmt == *ofmt {
		log.Fatalf("input format(%s) == output format(%s): nothing to do\n", *ifmt, *ofmt)
	}

	c := imgconv.Converter{
		InputFormat:  *ifmt,
		OutputFormat: *ofmt,
	}

	root, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Run(root)
	if err != nil {
		log.Fatal(err)
	}
}
