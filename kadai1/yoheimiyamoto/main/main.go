package main

import (
	"flag"
	"fmt"
	"os"

	imageconvert "github.com/YoheiMiyamoto/dojo4/kadai1/yoheimiyamoto/image-convert"
)

func main() {
	dirPath := flag.String("dir", "", "dir path")
	srcFormat := flag.String("srcFormat", "jpg", "srcFormat")
	destFormat := flag.String("destFormat", "", "destFormat")
	flag.Parse()

	if *destFormat == "" {
		fmt.Println("destFormat is required")
		return
	}

	result, err := imageconvert.Converts(*dirPath, *srcFormat, *destFormat)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// os.Stdout.Write([]byte("hello"))
	// os.Stdout.Write([]byte(result))
	fmt.Fprintf(os.Stdout, "hello")
	fmt.Fprintf(os.Stdout, result)
}
