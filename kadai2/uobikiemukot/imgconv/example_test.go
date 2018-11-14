package imgconv

import (
	"image/gif"
	"log"
)

func Example() {
	c := Converter{
		InputFormat:  "png",
		OutputFormat: "gif",
		GIFOptions:   gif.Options{NumColors: 256},
	}

	err := c.Run("test/subdir")
	if err != nil {
		log.Fatal(err)
	}
}
