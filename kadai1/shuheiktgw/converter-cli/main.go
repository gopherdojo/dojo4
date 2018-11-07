// Package main provides CLI interface for converter package
//
// Usage: converter-cli [options...] PATH
//
// OPTIONS:
//   --from value, -f value  specifies a image extension converted from (default: .jpg)
//   --to value, -t value    specifies a image extension converted to (default: .png)
//   --help, -h              prints out help
package main

import (
	"os"
)

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
