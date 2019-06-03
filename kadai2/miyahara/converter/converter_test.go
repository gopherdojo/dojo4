package converter

import (
	"image"
	"os"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	cases := map[string]struct {
		in        string
		out       string
		directory string
	}{
		"success": {
			in:        "inputFileName",
			out:       "outputFileName",
			directory: "targetDirectory",
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			converter := New(tc.in, tc.out, tc.directory)
			if converter.in != tc.in {
				t.Errorf("want %s,but actual %s\n", tc.in, converter.in)
			}

			if converter.out != tc.out {
				t.Errorf("want %s,but actual %s\n", tc.in, converter.out)
			}

			if converter.directory != tc.directory {
				t.Errorf("want %s, but actual %s\n", tc.directory, converter.directory)
			}
		})
	}
}

func TestConverter_Run(t *testing.T) {
	cases := map[string]struct {
		converter Converter
		expectInt int
	}{
		"success": {
			converter: Converter{in: "jpg", out: "png", directory: "TestData/jpg"},
			expectInt: 0,
		},
		"failed": {
			converter: Converter{in: "jpg", out: "png", directory: "test"},
			expectInt: 1,
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			actualInt := tc.converter.Run()
			if actualInt != tc.expectInt {
				t.Errorf("want %d, but actual %d", tc.expectInt, actualInt)
			}
		})
	}
}

func TestConverter_Convert(t *testing.T) {
	cases := map[string]struct {
		converter Converter
		fileSet   []FileSet
		errBool   bool
	}{
		"success": {
			converter: Converter{in: "jpg", out: "png"},
			fileSet: []FileSet{
				{
					inputFileName:  "TestData/jpg/testData.jpg",
					outputFileName: "TestData/jpg/testOutput.png",
				},
			},
			errBool: false,
		},
		"error": {
			converter: Converter{in: "jpg", out: "png"},
			fileSet: []FileSet{
				{
					inputFileName:  "test",
					outputFileName: "test",
				},
			},
			errBool: true,
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			err := tc.converter.convert(tc.fileSet)
			actualErrBool := err != nil
			if actualErrBool != tc.errBool {
				t.Errorf("want %t,but actual %t\n", tc.errBool, actualErrBool)
			}
		})
	}
}

func TestConverter_Execute(t *testing.T) {
	cases := map[string]struct {
		converter Converter
		fileSet   FileSet
		errBool   bool
	}{
		"success": {
			converter: Converter{in: "jpg", out: "png"},
			fileSet: FileSet{
				inputFileName:  "TestData/jpg/testData.jpg",
				outputFileName: "TestData/jpg/testOutput.png",
			},
			errBool: false,
		},
		"failed to open input file": {
			converter: Converter{in: "jpg", out: "png"},
			fileSet: FileSet{
				inputFileName:  "test",
				outputFileName: "testData/jpg/testOutput.png",
			},
			errBool: true,
		},
		"failed to open output file": {
			converter: Converter{in: "jpg", out: "png"},
			fileSet: FileSet{
				inputFileName:  "test",
				outputFileName: "testData/jpg/testOutput.png",
			},
			errBool: true,
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			err := tc.converter.execute(tc.fileSet)
			actualErrBool := err != nil
			if actualErrBool != tc.errBool {
				t.Errorf("want %t,but actual %t\n", tc.errBool, actualErrBool)
			}
		})
	}
}

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

			outputFile := testCreateFile(t, tc.outputFileName)
			defer outputFile.Close()

			inputFile := testOpenFile(t, tc.inputFileName)
			defer inputFile.Close()

			convertImage, _, err := image.Decode(inputFile)
			if err != nil {
				t.Error(err)
			}

			err = tc.converter.encode(outputFile, convertImage)
			if err != nil {
				t.Error(err)
			}

			file := testOpenFile(t, tc.outputFileName)
			defer file.Close()

			_, actual, err := image.Decode(file)
			if err != nil {
				t.Error(err)
			}

			if tc.converter.out != actual {
				t.Errorf("want %s, but actual %s\n", tc.converter.out, actual)
			}
		})
	}
}

func TestConverter_EncodeErr(t *testing.T) {
	cases := map[string]struct {
		converter Converter
		expectErr string
	}{
		"error": {
			converter: Converter{out: "error"},
			expectErr: "invalid output type",
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			err := tc.converter.encode(nil, nil)
			if err.Error() != tc.expectErr {
				t.Errorf("want %s,but actual %s\n", tc.expectErr, err.Error())
			}
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
		errBool      bool
	}{
		"testData Dir": {
			converter: Converter{in: "jpg", out: "png", directory: "./TestData/jpg"},
			fileSetSlice: []FileSet{
				{
					inputFileName:  "TestData/jpg/testData.jpg",
					outputFileName: "TestData/jpg/testOutput.png",
				},
			},
			errBool: false,
		},
		"not exist directory": {
			converter:    Converter{in: "jpg", out: "png", directory: "test"},
			fileSetSlice: nil,
			errBool:      true,
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			actual, err := tc.converter.dirWalk()
			if tc.fileSetSlice != nil {
				if !reflect.DeepEqual(tc.fileSetSlice[0], actual[0]) {
					t.Errorf("want %v ,but actual %v\n", tc.fileSetSlice[0], actual[0])
				}
			}
			actualErrBool := err != nil
			if tc.errBool != actualErrBool {
				t.Errorf("want %t, but actual %t\n", tc.errBool, actualErrBool)
			}
		})
	}
}

func testOpenFile(t *testing.T, fileName string) *os.File {
	t.Helper()
	file, err := os.Open(fileName)
	if err != nil {
		t.Fatal(err)
	}
	return file
}

func testCreateFile(t *testing.T, fileName string) *os.File {
	t.Helper()
	file, err := os.Create(fileName)
	if err != nil {
		t.Fatal(err)
	}
	return file
}
