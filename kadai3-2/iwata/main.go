package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gopherdojo/dojo4/kadai3-2/iwata/cmdparser"
	"github.com/gopherdojo/dojo4/kadai3-2/iwata/downloader"
)

func main() {
	p := cmdparser.New(os.Stderr)
	c, err := p.Parse(os.Args)
	if err != nil {
		log.Fatalf("Failed to parse arguments: %+v", err)
	}

	name := filepath.Base(c.URL)
	temp, err := ioutil.TempDir(os.TempDir(), name)
	if err != nil {
		log.Fatalf("Failed to create temp dir: %+v", err)
	}

	d := downloader.New(c.Parallel, temp)
	cf, err := d.Do(c.URL, time.Duration(c.Timeout)*time.Second)
	if err != nil {
		log.Fatalf("Failed to download: %+v", err)
	}

	dist := filepath.Join(c.Output, name)
	if err := cf.Save(dist); err != nil {
		log.Fatalf("Failed to save downloaded file: %+v", err)
	}

	fmt.Printf("Download %s successfully and save as %s", c.URL, dist)
}
