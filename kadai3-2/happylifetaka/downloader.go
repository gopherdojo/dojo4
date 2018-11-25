package main

import (
	"flag"
	"fmt"
	"github.com/happylifetaka/dojo4/kadai3-2/happylifetaka/downloader"
	"os"
)

type downloadByteInfo struct {
	start int64
	end int64
}

func main(){
	div := flag.Int64("div", 3, "file download division")
	usageMsg :="udage:downloader [-div] url saveFilePath"
	flag.Usage = func() {
		fmt.Println(usageMsg)
		flag.PrintDefaults()
		os.Exit(0)
	}
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("parameter error.")
		fmt.Println(usageMsg)
		flag.PrintDefaults()
		os.Exit(0)
	}
	url := args[0]
	saveFilePath := args[1]

	var d downloader.Downloader
	if err:=d.Download(url,saveFilePath,*div);err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}