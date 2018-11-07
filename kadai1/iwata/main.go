package main

import (
	"fmt"
	"log"
	"os"

	"github.com/iwata/dojo4/kadai1/iwata/cmdparser"
	"github.com/iwata/dojo4/kadai1/iwata/converter"
)

func main() {
	c, err := cmdparser.Parse()
	if err != nil {
		log.Fatal(err)
	}

	src, err := os.Open(c.SrcDir())
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	cv, err := converter.ResolveConverter(src, c)
	if err != nil {
		log.Fatal(err)
	}

	dist := fmt.Sprintf("%s.%s", c.SrcDir(), c.ToFormat())
	df, err := os.Create(dist)
	if err != nil {
		log.Fatal(err)
	}
	defer df.Close()

	if err := cv.Convert(df); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("convert %s to %s", c.SrcDir(), dist)
}
