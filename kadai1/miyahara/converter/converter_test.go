package converter

import (
	"image"
	"os"
	"reflect"
	"testing"
)

func TestConverter_Encode(t *testing.T) {
	cases := map[string]struct {
		converter      Converter
		inputFileName  string
		outputFileName string
	}{
		"jpg to png": {
			converter:      Converter{out: "png"},
			inputFileName:  "TestData/jpg/testData.jpg",
			outputFileName: "TestData/jpg/testOutput.png",
		},
		"png to jpeg": {
			converter:      Converter{out: "jpeg"},
			inputFileName:  "TestData/png/testData.png",
			outputFileName: "TestData/png/testOutput.jpg",
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			var err error
			outputFile, err := os.Create(tc.outputFileName)
			if err != nil {
				t.Error(err)
				if outputFile != nil {
					outputFile.Close()
				}
			}

			inputFile, err := os.Open(tc.inputFileName)
			if err != nil {
				t.Error(err)
				if inputFile != nil {
					inputFile.Close()
				}
			}

			convertImage, _, err := image.Decode(inputFile)
			if err != nil {
				t.Error(err)
			}

			err = tc.converter.encode(outputFile, convertImage)
			if err != nil {
				t.Error(err)
			}

			file, err := os.Open(tc.outputFileName)
			if err != nil {
				t.Error(err)
				if file != nil {
					file.Close()
				}
			}
			_, actual, err := image.Decode(file)
			if err != nil {
				t.Error(err)
			}

			if tc.converter.out != actual {
				t.Errorf("want %s, but actual %s\n", tc.converter.out, actual)
			}

			outputFile.Close()
			inputFile.Close()
			file.Close()
		})
	}
}

func TestConverter_OutputFilePath(t *testing.T) {
	cases := map[string]struct {
		converter      Converter
		inputFileName  string
		outputFileName string
	}{
		"jpg to png": {
			converter:      Converter{out: "png"},
			inputFileName:  "/test/test.jpg",
			outputFileName: "/test/test.png",
		},
		"png to jpg": {
			converter:      Converter{out: "jpg"},
			inputFileName:  "test/test.png",
			outputFileName: "test/test.jpg",
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			actual := tc.converter.outputFilePath(tc.inputFileName)
			if tc.outputFileName != actual {
				t.Errorf("want %s, but actual %s\n", tc.outputFileName, actual)
			}
		})
	}

}

func TestConverter_DirWalk(t *testing.T) {
	cases := map[string]struct {
		converter    Converter
		fileSetSlice []FileSet
	}{
		"testData Dir": {
			converter: Converter{in: "jpg", out: "png", directory: "./TestData/jpg"},
			fileSetSlice: []FileSet{
				FileSet{
					inputFileName:  "TestData/jpg/testData.jpg",
					outputFileName: "TestData/jpg/testData.png",
				},
			},
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			actual := tc.converter.dirWalk()
			if !reflect.DeepEqual(tc.fileSetSlice[0], actual[0]) {
				t.Errorf("want %v ,but actual %v", tc.fileSetSlice[0], actual[0])
			}
		})
	}
}
