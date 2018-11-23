package main

import (
	"flag"
	"fmt"
	"os"

	imageconvert "github.com/YoheiMiyamoto/dojo4/kadai2/yoheimiyamoto/image-convert"
)

func main() {
	dirPath := flag.String("dir", "", "dir path")
	srcFormat := flag.String("srcFormat", "jpg", "srcFormat")
	destFormat := flag.String("destFormat", "", "destFormat")
	flag.Parse()

	if *destFormat == "" {
		fmt.Fprintln(os.Stdout, "destFormat is required")
		os.Exit(1)
		return
	}

	result, err := imageconvert.Converts(*dirPath, *srcFormat, *destFormat)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, result)
}
