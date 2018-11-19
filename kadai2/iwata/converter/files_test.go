package converter_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/gopherdojo/dojo4/kadai2/iwata/converter"
	"github.com/gopherdojo/dojo4/kadai2/iwata/testutil"
)

type mockConvertOption struct {
	format string
}

func (o *mockConvertOption) ToFormat() string {
	return o.format
}

func newMockConvertOption(format string) converter.ConvertOption {
	return &mockConvertOption{format}
}

func filePathes(files []string) []string {
	var goLogoDir = "../test/fixtures/images/Go-Logo"
	pathes := make([]string, 0, len(files))
	for _, f := range files {
		pathes = append(pathes, fmt.Sprintf("%s/%s", goLogoDir, f))
	}
	return pathes
}

func TestConvertFiles(t *testing.T) {
	type args struct {
		files []string
		c     converter.ConvertOption
	}
	tests := []struct {
		name         string
		args         args
		wantErr      error
		createdFiles []string
	}{
		{"empty files", args{[]string{}, newMockConvertOption("jpg")}, nil, []string{}},
		{
			"not exits files",
			args{[]string{"hoge.jpg", "fuga.png"}, newMockConvertOption("gif")},
			errors.New("Failed to open"),
			[]string{},
		},
		{
			"not support format",
			args{[]string{"JPG/Go-Logo_Aqua.jpg"}, newMockConvertOption("svg")},
			errors.New("svg is not supported format"),
			[]string{},
		},
		{
			"cannot decode file",
			args{[]string{"EPS/Go-Logo_Versions.eps"}, newMockConvertOption("png")},
			errors.New("Failed to decode"),
			[]string{},
		},
		{"convert to png", args{[]string{
			"JPG/Go-Logo_Black.jpg",
			"JPG/Go-Logo_Yellow.jpg",
		}, newMockConvertOption("png")}, nil, []string{
			"JPG/Go-Logo_Black.png",
			"JPG/Go-Logo_Yellow.png",
		}},
		{"convert to jpg", args{[]string{
			"PNG/Go-Logo_Blue.png",
			"PNG/Go-Logo_LightBlue.png",
		}, newMockConvertOption("jpg")}, nil, []string{
			"PNG/Go-Logo_Blue.jpg",
			"PNG/Go-Logo_LightBlue.jpg",
		}},
		{"convert to gif", args{[]string{
			"JPG/Go-Logo_Aqua.jpg",
			"PNG/Go-Logo_Fuchsia.png",
			"PNG/Go-Logo_White.png",
		}, newMockConvertOption("gif")}, nil, []string{
			"JPG/Go-Logo_Aqua.gif",
			"PNG/Go-Logo_Fuchsia.gif",
			"PNG/Go-Logo_White.gif",
		}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := converter.ConvertFiles(filePathes(tt.args.files), tt.args.c)
			testutil.ContainsError(t, err, tt.wantErr, "converter.ConvertFiles() Error")
			for _, f := range filePathes(tt.createdFiles) {
				testFileExist(t, f)
				os.Remove(f)
			}
		})
	}
}

func testFileExist(t *testing.T, file string) {
	t.Helper()
	if _, err := os.Stat(file); os.IsNotExist(err) {
		t.Errorf("[%s] does not exist", file)
	}
}
