package main

import (
	"fmt"
	"log"

	"github.com/iwata/dojo4/kadai1/iwata/cmdparser"
	"github.com/iwata/dojo4/kadai1/iwata/converter"
	"github.com/iwata/dojo4/kadai1/iwata/filter"
)

func main() {
	c, err := cmdparser.Parse()
	if err != nil {
		log.Fatal(err)
	}

	files, err := filter.Files(c)
	if err != nil {
		log.Fatal(err)
	}

	if err := converter.ConvertFiles(files, c); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Converted %d %s files to %s in %s successfully\n", len(files), c.FromFormat(), c.ToFormat(), c.SrcDir())
}
