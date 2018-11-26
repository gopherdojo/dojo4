package main

import (
	"flag"
	"github.com/keitarou/dojo4/kadai1/keitarou/lib"
	"log"
)

func main() {
	src := flag.String("s", "jpeg", "[jpeg/png/gif]")
	dst := flag.String("d", "png", "[jpeg/png/gif]")
	flag.Parse()
	filepath := flag.Arg(0)

	err := converter.ConvertWalk(converter.ConvertWalkOption{
		filepath,
		*src,
		*dst,
	})

	if err != nil {
		log.Fatalln(err)
	}
}
