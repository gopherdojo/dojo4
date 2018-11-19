package converter_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/gopherdojo/dojo4/kadai2/akuchii/converter"
)

func TestConvert_NewConverter(t *testing.T) {
	cases := []struct {
		afterExt        string
		expectedEncoder reflect.Type
	}{
		{"png", reflect.ValueOf(&converter.PngEncoder{}).Type()},
		{"jpeg", reflect.ValueOf(&converter.JpegEncoder{}).Type()},
		{"gif", reflect.ValueOf(&converter.GifEncoder{}).Type()},
	}
	for _, c := range cases {
		conv, err := converter.NewConverter("", "", c.afterExt)
		if err != nil {
			t.Fatalf("failed to crete new converter:%s", err)
		}
		if reflect.ValueOf(conv.ExportEncoder()).Type() != c.expectedEncoder {
			t.Fatalf("invalid encoder")
		}
	}
}

func TestConverter_Execute(t *testing.T) {
	cases := []struct {
		target   string
		outDir   string
		afterExt string
		hasError bool
	}{
		{"test.png", "out", "jpeg", false},
		{"test.jpeg", "out", "gif", false},
		{"test.gif", "out", "png", false},
		{"test.webp", "out", "png", true},
		{"test.png", "out", "webp", true},
		{"not_exist_file.png", "out", "jpeg", true},
	}

	for _, c := range cases {
		srcPath := filepath.Join("..", "testdata", c.target)
		filenameWithoutExt := testGetFileNameWithoutExt(t, c.target)
		dstPath := filepath.Join("..", "testdata", c.outDir, filenameWithoutExt) + "." + c.afterExt
		conv, err := converter.NewConverter(srcPath, c.outDir, c.afterExt)
		if err != nil {
			t.Fatalf("fail to create converter: %s", err)
		}
		err = conv.Execute()
		if err != nil && !c.hasError {
			t.Fatalf("fail to convert image: %s", err)
		}
		if _, err := os.Stat(dstPath); os.IsNotExist(err) && !c.hasError {
			t.Fatalf("fail to generate converted image: %s", err)
		}
		if !c.hasError {
			os.Remove(dstPath)
		}
	}
}

func testGetFileNameWithoutExt(t *testing.T, target string) string {
	t.Helper()

	return target[:len(target)-len(filepath.Ext(target))]
}
