package converter

import (
	"errors"
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
func New(in, out, directory string) *Converter {
	return &Converter{in: in, out: out, directory: directory}
}

// Run execute image convert function.
func (c Converter) Run() int {
	var err error
	fileSetSlice, err := c.dirWalk()
	if err != nil {
		fmt.Println(err)
		return 1
	}
	err = c.convert(fileSetSlice)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}

// dirWalk returns file set of input file name and output file in target directory.
func (c Converter) dirWalk() ([]FileSet, error) {
	fileSetSlice := make([]FileSet, 0, 50)
	err := filepath.Walk(c.directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(info.Name()) == ("." + c.in) {
			inputFileName := path
			outputFileName := c.outputFilePath(inputFileName)
			fileSet := FileSet{inputFileName: inputFileName, outputFileName: outputFileName}
			fmt.Println(inputFileName, outputFileName)
			fileSetSlice = append(fileSetSlice, fileSet)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return fileSetSlice, nil
}

// convert calls execute function every file set.
func (c Converter) convert(fileSetSlice []FileSet) error {
	var err error

	for _, fileSet := range fileSetSlice {
		err = c.execute(fileSet)
		if err != nil {
			break
		}
	}

	return err
}

// execute function open file and calls encode function.
func (c Converter) execute(fileset FileSet) error {
	var err error

	inputFile, err := os.Open(fileset.inputFileName)
	defer inputFile.Close()
	if err != nil {
		return err
	}

	outputFile, err := os.Create(fileset.outputFileName)
	defer outputFile.Close()
	if err != nil {
		return err
	}

	img, _, err := image.Decode(inputFile)
	if err != nil {
		return err
	}

	err = c.encode(outputFile, img)
	if err != nil {
		return err
	}

	return nil
}

// encode returns error by encode function.
// if output type is not support, Encode returns error.
func (c Converter) encode(file *os.File, image image.Image) error {
	switch c.out {
	case "jpg", "jpeg":
		err := jpeg.Encode(file, image, &jpeg.Options{Quality: 80})
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
