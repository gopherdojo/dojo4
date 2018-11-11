package main

import (
	"flag"
	"log"

	"convimg"
)

func main() {

	var suffix = []string{"jpg", "jpeg", "png", "gif"}

	var (
		f string
		t string
	)

	flag.StringVar(&f, "str", "jpg", "format before conversion (default: jpg)")
	flag.StringVar(&t, "str", "jpg", "format after conversion (default: png)")

	flag.Parse()

	path := flag.Arg(0)

	if contains(suffix, f) != false || contains(suffix, t) != false {
		log.Fatal("Invalid suffix: %s, %s", f, t)
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
