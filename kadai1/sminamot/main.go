package main

import (
	"flag"
	"log"

	"github.com/sminamot/dojo4/kadai1/sminamot/convert"
)

func main() {
	var src, dst string

	flag.StringVar(&src, "src", "jpg", "input format")
	flag.StringVar(&dst, "dst", "png", "output format")
	flag.Parse()
	if len(flag.Args()) == 0 {
		log.Fatal("need to specify target dir")
	}
	dir := flag.Args()[0]

	c, err := convert.New(dir, src, dst)
	if err != nil {
		log.Fatal(err)
	}
	c.Convert()
}
