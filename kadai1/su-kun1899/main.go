package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Hello, dojo")
}

func replaceExt(fileName, newExt string) string {
	return fmt.Sprintf(
		"%s.%s", strings.TrimSuffix(fileName, filepath.Ext(fileName)),
		newExt,
	)
}
