package main

import (
	"fmt"
	"log"

	"github.com/gopherdojo/dojo4/kadai2/iwata/cmdparser"
	"github.com/gopherdojo/dojo4/kadai2/iwata/converter"
	"github.com/gopherdojo/dojo4/kadai2/iwata/filter"
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
