package main

import (
	"errors"
	"fmt"
	"github.com/gopherdojo/dojo4/kadai2/su-kun1899/imgconv"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	os.Exit(runCmd(os.Args[1:]))
}

func runCmd(args []string) int {
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, errors.New("missing target directory"))
		return 1
	}

	targetDir := args[0]
	converter := imgconv.PngConv{}
	err := filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
		if !imgconv.Is(path, imgconv.JpegFormat) {
			return nil
		}

		src := path
		dest := replaceExt(src, "png")
		err = converter.Convert(src, dest)
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "%s => %s\n", src, dest)

		return nil
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	return 0
}

func replaceExt(fileName, newExt string) string {
	return fmt.Sprintf("%s.%s", strings.TrimSuffix(fileName, filepath.Ext(fileName)), newExt)
}
