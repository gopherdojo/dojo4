package main

import (
	"context"
	"fmt"
	"os"

	"github.com/YoheiMiyamoto/dojo4/kadai3-2/yoheimiyamoto/downloader"
)

func main() {
	c := downloader.NewClient("https://images.pexels.com/photos/248304/pexels-photo-248304.jpeg")
	err := c.Download(context.Background())
	if err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
	}
}
