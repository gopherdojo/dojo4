package converter

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// Converter Convert jpg to png or png to jpg.
// In represent a input file type.
// Out represent a output file type.
// Directory represent a target directory.
type Converter struct {
	in        string
	out       string
	directory string
}

// FileSet is Set of input file name and output file name.
type FileSet struct {
	inputFileName  string
	outputFileName string
}

// New return a new Converter.
func New() *Converter {
	var inputFile = flag.String("i", "jpg", "input image type")
	var outputFile = flag.String("o", "png", "output image type")
	flag.Parse()
	return &Converter{in: *inputFile, out: *outputFile, directory: os.Args[3]}
}

// Run execute image convert function.
func (c Converter) Run() {
	fileSetSlice := c.dirWalk()
	c.convert(fileSetSlice)
}

// dirWalk returns file set of input file name and output file in target directory.
func (c Converter) dirWalk() []FileSet {
	fileSetSlice := make([]FileSet, 0, 50)
	filepath.Walk(c.directory, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(info.Name()) == ("." + c.in) {
			inputFileName := path
			outputFileName := c.outputFilePath(inputFileName)
			fileSet := FileSet{inputFileName: inputFileName, outputFileName: outputFileName}
			fmt.Println(inputFileName, outputFileName)
			fileSetSlice = append(fileSetSlice, fileSet)
		}
		return nil
	})

	return fileSetSlice
}

// convert calls Encode function every file set.
func (c Converter) convert(fileSetSlice []FileSet) {
	for _, fileSet := range fileSetSlice {
		func(fileSet FileSet) {
			inputFile, _ := os.Open(fileSet.inputFileName)
			defer inputFile.Close()
			outputFile, _ := os.Create(fileSet.outputFileName)
			defer outputFile.Close()
			image, _, _ := image.Decode(inputFile)
			c.encode(outputFile, image)
		}(fileSet)
	}
}

// encode returns error by encode function.
// if output type is not support, Encode returns error.
func (c Converter) encode(file *os.File, image image.Image) error {
	switch c.out {
	case "jpg", "jpeg":
		err := jpeg.Encode(file, image, &jpeg.Options{Quality: 100})
		return err
	case "png":
		err := png.Encode(file, image)
		return err
	default:
		return errors.New("invalid output type")
	}
}

// outputFilePath returns output file path to correspond input file path.
func (c Converter) outputFilePath(inputFileName string) string {
	stringSlice := strings.Split(inputFileName, ".")
	outputFileName := stringSlice[0] + "." + c.out
	return outputFileName
}
