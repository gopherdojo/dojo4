package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo2/kadai1/uobikiemukot/imgconv"
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

	c := imgconv.Config{
		InputFormat: *ifmt,
		OutputFormat: *ofmt,
	}

	root, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	files, err := c.SearchImages(root)
	if err != nil {
		log.Fatalf("SearchImages() failed: %s", err)
	}

	for _, f := range files {
		fmt.Fprintf(os.Stderr, "converting...: %s\n", f)
		err = c.ConvertImage(f)
		if err != nil {
			log.Printf("ConvertImage() failed: %s", err)
		}
	}
}
