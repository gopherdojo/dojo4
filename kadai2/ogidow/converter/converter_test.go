package converter

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindFiles(t *testing.T) {
	cases := []struct {
		dir        string
		extensions []string
		files      []string
	}{
		{
			"../fixtures/images",
			[]string{".jpg"},
			[]string{"../fixtures/images/jpeg/logo.jpg"},
		},
		{
			"../fixtures/images",
			[]string{".png"},
			[]string{"../fixtures/images/png/logo.png"},
		},
		{
			"../fixtures/images",
			[]string{".gif"},
			[]string{"../fixtures/images/gif/logo.gif"},
		},
	}

	for _, c := range cases {
		result, err := FindFiles(c.dir, c.extensions)
		if err != nil {
			t.Errorf("FindFiles failed: %s\n", err)
		}

		if diff := cmp.Diff(result, c.files); diff != "" {
			t.Errorf("FindFiles differs: (-got +want)\n%s", diff)
		}
	}
}

func TestConvertSuccess(t *testing.T) {
	cases := []struct {
		srcPath       string
		destExtention string
		destFilePath  string
	}{
		{
			"../fixtures/images/jpeg/logo.jpg",
			".gif",
			"../fixtures/images/jpeg/logo.gif",
		},
		{
			"../fixtures/images/jpeg/logo.jpg",
			".png",
			"../fixtures/images/jpeg/logo.gif",
		},
		{
			"../fixtures/images/png/logo.png",
			".jpg",
			"../fixtures/images/png/logo.jpg",
		},
		{
			"../fixtures/images/png/logo.png",
			".gif",
			"../fixtures/images/png/logo.gif",
		},
		{
			"../fixtures/images/gif/logo.gif",
			".jpg",
			"../fixtures/images/gif/logo.jpg",
		},
		{
			"../fixtures/images/gif/logo.gif",
			".png",
			"../fixtures/images/gif/logo.png",
		},
	}

	for _, c := range cases {
		err := Convert(c.srcPath, c.destExtention)
		if err != nil {
			t.Errorf("Convert failed: %s\n", err)
		}

		filename := DestFilePath(c.srcPath, c.destExtention)
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			fmt.Errorf("file does not exist")
		} else {
			os.Remove(filename)
		}

	}

}
