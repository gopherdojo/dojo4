package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/happylifetaka/dojo4/kadai2/happylifetaka/imgconv"
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

	if len(args) != 1 {
		fmt.Println("parameter error.")
		fmt.Println(usageMsg)
		flag.PrintDefaults()
		os.Exit(1)
	}

	var ic imgconv.ImgConverter

	e := ic.SetConvertFormat(*fromFormat, *toFormat)

	if e != nil {
		fmt.Println(e)
		fmt.Println(usageMsg)
		flag.PrintDefaults()
		os.Exit(1)
	}

	rs, err := ic.Convert(args[0])
	if err != nil {
		fmt.Println(err)
	}

	for _, d := range rs {
		fmt.Println(d.Msg)
		if d.Err != nil {
			fmt.Println(d.Err)
		}
	}
}
