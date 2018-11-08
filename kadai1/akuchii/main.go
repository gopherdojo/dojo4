package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo4/kadai1/akuchii/converter"
)

// extension before convert
var beforeExt string

// extension after convert
var afterExt string

// source directory
var srcDir string

// destination directory
var dstDir string

// exit code
const (
	ExitCodeOK    int = iota // 0
	ExitCodeError            // 1
)

func init() {
	flag.StringVar(&beforeExt, "b", "jpeg", "extension before convert")
	flag.StringVar(&afterExt, "a", "png", "extension before convert")
	flag.StringVar(&srcDir, "s", "", "source directory to convert files")
	flag.StringVar(&dstDir, "d", "out", "destination directory to convert files")
}

func main() {
	flag.Parse()

	err := validateArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(ExitCodeError)
	}

	fmt.Printf("start to convert ext before: %s, after: %s\n", beforeExt, afterExt)
	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == "."+beforeExt {
			err = converter.Execute(path, dstDir, afterExt)
			if err != nil {
				fmt.Println(err)
				os.Exit(ExitCodeError)
			}
			fmt.Printf("%s is converted to %s\n", path, afterExt)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(ExitCodeError)
	}
	fmt.Println("convert is finished")
	os.Exit(ExitCodeOK)
}

// validateArgs validates input arguments
func validateArgs() error {
	if beforeExt == "" || afterExt == "" || srcDir == "" || dstDir == "" {
		flag.PrintDefaults()
		return errors.New("empty arg is not allowed")
	}

	if _, err := os.Stat(srcDir); os.IsNotExist(err) {
		return err
	}

	if beforeExt == afterExt {
		return fmt.Errorf("before and after ext is same: %s", beforeExt)
	}
	return nil
}
