package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gopherdojo/dojo4/kadai1/Udomomo/convimg"
)

func main() {

	SUFFIX := []string{"jpg", "jpeg", "png", "gif"}
	SEPARATOR := "."

	var (
		f string
		t string
	)

	flag.StringVar(&f, "f", "jpg", "format before conversion")
	flag.StringVar(&t, "t", "png", "format after conversion")

	flag.Parse()

	path := flag.Arg(0)

	if contains(suffix, f) == false || contains(suffix, t) == false {
		fmt.Printf("Invalid suffix: %s, %s", f, t)
		os.Exit(1)
	}

	convimg.SearchFile(path, "."+f, "."+t)
}

func contains(su []string, fl string) bool {
	for _, s := range su {
		if s == fl {
			return true
		}
	}
	return false
}
