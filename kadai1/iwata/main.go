package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/iwata/dojo4/kadai1/iwata/cmdparser"
	"github.com/iwata/dojo4/kadai1/iwata/converter"
)

type filterConfig interface {
	SrcDir() string
	FromFormat() string
}

func main() {
	c, err := cmdparser.Parse()
	if err != nil {
		log.Fatal(err)
	}

	files, err := filterFiles(c)
	if err != nil {
		log.Fatal(err)
	}

	if err := converter.ConvertFiles(files, c); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Converted %d %s files to %s in %s successfully\n", len(files), c.FromFormat(), c.ToFormat(), c.SrcDir())
}

func filterFiles(c filterConfig) ([]string, error) {
	var files []string
	ext := fmt.Sprintf(".%s", c.FromFormat())

	err := filepath.Walk(c.SrcDir(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ext) {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}
