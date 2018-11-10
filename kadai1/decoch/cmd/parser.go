package cmd

import (
	"errors"
	"flag"

	"github.com/decoch/dojo4/kadai1/decoch/converter"
)

type Command struct {
	Dir         string
	ConvertType converter.Type
}

func Parse() (*Command, error) {
	var cmd Command
	reverseFlag := flag.Bool("r", false, "Reverse image convert")
	flag.Parse()
	reverse := *reverseFlag

	args := flag.Args()
	if len(args) != 1 {
		return nil, errors.New("Invalid args")
	}
	dirName := args[0]
	cmd.Dir = dirName

	var convertType converter.Type
	if reverse {
		convertType = converter.PngToJpg
	} else {
		convertType = converter.JpgToPng
	}
	cmd.ConvertType = convertType

	return &cmd, nil
}
