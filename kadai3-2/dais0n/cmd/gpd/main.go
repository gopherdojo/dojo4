package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/dais0n/dojo4/kadai3-2/dais0n/client"
)

func main() {
	// get flags
	flags := flag.NewFlagSet("gpd", flag.ContinueOnError)
	if err := flags.Parse(os.Args[1:]); err != nil {
		os.Exit(1)
	}
	url, _ := url.Parse(flags.Arg(0))
	logger := log.New(os.Stdout, "gpdLog: ", 0)
	// create gpc client
	gpdClient, err := client.NewGpcClient(*url, logger)
	if err != nil {
		fmt.Println(err)
	}
	// check whether you can request content by range access
	if !gpdClient.RequestHeader() {
		fmt.Fprintf(os.Stdout, "unable to request by range access")
	}
	// request
	gpdClient.GetContent()
}
