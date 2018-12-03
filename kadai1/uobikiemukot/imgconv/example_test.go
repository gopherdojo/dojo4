package imgconv

import (
	"image/gif"
	"log"
)

func Example() {
	c := Config{
		InputFormat:  "png",
		OutputFormat: "gif",
		GIFOptions:   gif.Options{NumColors: 256},
	}

	err := c.Converter("test/subdir")
	if err != nil {
		log.Fatal(err)
	}
}
