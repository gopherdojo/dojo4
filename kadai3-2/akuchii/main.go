package main

import (
	"fmt"
	"os"

	"github.com/gopherdojo/dojo4/kadai3-2/akuchii/downloader"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("require url")
		os.Exit(1)
	}

	var splitNum uint = 4
	url := os.Args[1]
	d := downloader.NewDownloader(url, splitNum)

	if err := d.Prepare(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := d.Download(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := d.MergeFiles(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("download complete! : %v\n", d.GetFileName())
	os.Exit(0)
}
