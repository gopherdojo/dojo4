package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/happylifetaka/dojo4/kadai1/imgconv"
)

func main() {
	fromFormat := flag.String("f", "jpg", "from image type.[jpg,gif,png]")
	toFormat := flag.String("t", "png", "to image type.[jpg,gif,png]")
	usageMsg := "usage:imgconv [option -f] [option -t] targetFilePath"
	flag.Usage = func() {
		fmt.Println(usageMsg)
		flag.PrintDefaults()
		os.Exit(0)
	}
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 || !imgconv.FormatCheck(*fromFormat, *toFormat) {
		fmt.Println("parameter error.")
		fmt.Println(usageMsg)
		flag.PrintDefaults()
		os.Exit(1)
	}
	_, err := os.Stat(args[0])
	if err != nil {
		fmt.Println("target file path is not exist.")
		os.Exit(1)
	}
	errCnv := imgconv.Convert(args[0], *fromFormat, *toFormat)
	if errCnv != nil {
		fmt.Println(errCnv)
	}
}
