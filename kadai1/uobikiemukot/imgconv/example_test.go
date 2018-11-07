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

	imgs, err := c.SearchImages("test/subdir")
	if err != nil {
		log.Fatal(err)
	}

	for _, img := range imgs {
		log.Printf("converting %s...\n", img)
		err := c.ConvertImage(img)
		if err != nil {
			log.Println(err)
		}
	}
}
