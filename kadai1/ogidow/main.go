package main

import (
	"flag"
	"fmt"

	"github.com/ogidow/dojo4/kadai1/ogidow/converter"
)

func main() {
	argument := converter.Argument{}

	flag.StringVar(&argument.InputFormat, "i", "jpg", "format for input image")
	flag.StringVar(&argument.OutputFormat, "o", "png", "format for output image")
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("Argument is required")
		return
	}

	if !argument.IsValid() {
		fmt.Println("invalid Arguments")
		return
	}

	dir := flag.Arg(0)

	files, err := converter.FindFiles(dir, argument.InputExtensions())
	for _, file := range files {
		err = converter.Convert(file, argument.OutputExtension())
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if err != nil {
		fmt.Println(err)
		return
	}
}
