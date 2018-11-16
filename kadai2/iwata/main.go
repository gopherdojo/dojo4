package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gopherdojo/dojo4/kadai2/iwata/cmdparser"
	"github.com/gopherdojo/dojo4/kadai2/iwata/converter"
	"github.com/gopherdojo/dojo4/kadai2/iwata/filter"
)

func main() {
	cmd := cmdparser.NewCmd(os.Stderr)
	c, err := cmd.Parse(os.Args)
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
