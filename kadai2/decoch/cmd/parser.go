package cmd

import (
	"errors"
	"flag"

	"github.com/decoch/dojo4/kadai2/decoch/converter"
)

// Command represents command line args.
type Command struct {
	Dir        string
	InputType  converter.ImageType
	OutputType converter.ImageType
}

// Parse parse args.
func Parse() (*Command, error) {
	var cmd Command
	inputTypeFlag := flag.String("i", "jpg", "Input image file type.")
	outputTypeFlag := flag.String("o", "png", "Output image file type.")
	flag.Parse()
	inputType := converter.NewImageType(*inputTypeFlag)
	outputType := converter.NewImageType(*outputTypeFlag)

	args := flag.Args()
	if len(args) != 1 {
		return nil, errors.New("Invalid args")
	}
	dirName := args[0]
	cmd.Dir = dirName

	if inputType == nil || len(string(*inputType)) == 0 {
		return nil, errors.New("Invalid input image type.")
	}
	cmd.InputType = *inputType

	if outputType == nil || len(string(*outputType)) == 0 {
		return nil, errors.New("Invalid output image type.")
	}
	cmd.OutputType = *outputType
	return &cmd, nil
}
