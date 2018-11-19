package main

import (
	"flag"
	"fmt"
	"github.com/gopherdojo/dojo4/kadai2/phamvanhung2e123/converter"
	"os"
	"path/filepath"
)

var (
	jpgType  = "jpg"
	jpegType = "jpeg"
	gifType  = "gif"
	pngType  = "png"
)

var (
	inputType  string
	outputType string
)

var validatedTypes = map[string]bool{gifType: true, jpgType: true, jpegType: true, pngType: true}

func convert(path string, inputType string, outputType string) (err error) {
	converter := converter.New(inputType, outputType)
	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		switch {
		case err != nil:
			return err
		case info.IsDir():
			return nil
		default:
			return converter.ConvertImage(path)
		}
	})
	return err
}

func isValidType(imageType string) bool {
	_, ok := validatedTypes[imageType]
	return ok
}

func init() {
	flag.StringVar(&inputType, "i", jpgType, "Input file type")
	flag.StringVar(&outputType, "o", pngType, "Output file type")
}

func main() {
	flag.Parse()
	if !isValidType(inputType) || !isValidType(outputType) {
		fmt.Println("Please input valid type -i [png, jpeg, gif, png] -o [png, jpeg, gif, png] ")
	}
	if flag.NArg() == 0 {
		fmt.Println("Please input file path")
	}

	for i := 0; i < flag.NArg(); i++ {
		path := flag.Arg(i)
		err := convert(path, inputType, outputType)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
	}
}
