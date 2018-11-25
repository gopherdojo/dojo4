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

	// TODO fromとtoをオプションで受け取りたい
	fromFormat := imgconv.JpegFormat
	toFormat := imgconv.PngFormat

	targetDir := args[0]
	err := filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
		if !imgconv.Is(path, fromFormat) {
			return nil
		}

		converted, err := imgconv.Convert(path, toFormat)
		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s => %s\n", path, converted)

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
