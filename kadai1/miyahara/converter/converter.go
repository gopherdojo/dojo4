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
	In        string
	Out       string
	Directory string
}

// FileSet is Set of input file name and output file name.
type FileSet struct {
	inputFileName  string
	outputFileName string
}

// NewConverter return a new Converter.
func NewConverter() *Converter {
	var inputFile = flag.String("i", "jpg", "input image type")
	var outputFile = flag.String("o", "png", "output image type")
	flag.Parse()
	return &Converter{In: *inputFile, Out: *outputFile, Directory: os.Args[3]}
}

// Run execute image convert function.
func (c Converter) Run() {
	fileSetSlice := c.DirWalk()
	c.Convert(fileSetSlice)
}

// DirWalk returns file set of input file name and output file in target directory.
func (c Converter) DirWalk() []FileSet {
	fileSetSlice := make([]FileSet, 0, 50)
	filepath.Walk(c.Directory, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(info.Name()) == ("." + c.In) {
			inputFileName := path
			outputFileName := c.OutputFilePath(inputFileName)
			fileSet := FileSet{inputFileName: inputFileName, outputFileName: outputFileName}
			fmt.Println(inputFileName, outputFileName)
			fileSetSlice = append(fileSetSlice, fileSet)
		}
		return nil
	})

	return fileSetSlice
}

// Convert calls Encode function every file set.
func (c Converter) Convert(fileSetSlice []FileSet) {
	for _, fileSet := range fileSetSlice {
		func(fileSet FileSet) {
			inputFile, _ := os.Open(fileSet.inputFileName)
			defer inputFile.Close()
			outputFile, _ := os.Create(fileSet.outputFileName)
			defer outputFile.Close()
			image, _, _ := image.Decode(inputFile)
			c.Encode(outputFile, image)
		}(fileSet)
	}
}

// Encode returns error by encode function.
// if output type is not support, Encode returns error.
func (c Converter) Encode(file *os.File, image image.Image) error {
	switch c.Out {
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

// OutputFilePath returns output file path to correspond input file path.
func (c Converter) OutputFilePath(inputFileName string) string {
	stringSlice := strings.Split(inputFileName, ".")
	outputFileName := stringSlice[0] + "." + c.Out
	return outputFileName
}
