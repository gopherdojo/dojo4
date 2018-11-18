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

	args := flag.Args()
	dirName, err := parseCommandLineArgs(args)
	if err != nil {
		return nil, err
	}
	cmd.Dir = dirName

	inputType, err := parseImageType(*inputTypeFlag)
	if err != nil {
		return nil, err
	}
	cmd.InputType = *inputType

	outputType, err := parseImageType(*outputTypeFlag)
	if err != nil {
		return nil, err
	}
	cmd.OutputType = *outputType

	return &cmd, nil
}

func parseCommandLineArgs(args []string) (string, error) {
	if len(args) != 1 {
		return "", errors.New("Invalid args")
	}
	return args[0], nil
}

func parseImageType(str string) (*converter.ImageType, error) {
	inputType := converter.NewImageType(str)
	if inputType == nil || len(string(*inputType)) == 0 {
		return nil, errors.New("Invalid image type.")
	}
	return inputType, nil
}
