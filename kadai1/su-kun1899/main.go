package main

import (
	"errors"
	"fmt"
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
	println(targetDir)

	return 0
}


func replaceExt(fileName, newExt string) string {
	return fmt.Sprintf("%s.%s", strings.TrimSuffix(fileName, filepath.Ext(fileName)), newExt)
}
