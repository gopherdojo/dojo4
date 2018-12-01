package imgconv

import (
	"errors"
	"fmt"
	"testing"
)

func TestSetConvertFormat(t *testing.T) {
	var ic ImgConverter

	testCases := []struct {
		fromFormat       string
		toFormat         string
		resutlErr        bool
		resultFromFormat string
		resultToFormat   string
		resultSetFormat  bool
	}{
		{fromFormat: "jpg", toFormat: "jpg", resutlErr: true, resultFromFormat: "", resultToFormat: "", resultSetFormat: false},
		{fromFormat: "jpg", toFormat: "png", resutlErr: false, resultFromFormat: "jpg", resultToFormat: "png", resultSetFormat: true},
		{fromFormat: "jpg", toFormat: "gif", resutlErr: false, resultFromFormat: "jpg", resultToFormat: "gif", resultSetFormat: true},

		{fromFormat: "png", toFormat: "png", resutlErr: true, resultFromFormat: "", resultToFormat: "", resultSetFormat: false},
		{fromFormat: "png", toFormat: "jpg", resutlErr: false, resultFromFormat: "png", resultToFormat: "jpg", resultSetFormat: true},
		{fromFormat: "png", toFormat: "gif", resutlErr: false, resultFromFormat: "png", resultToFormat: "gif", resultSetFormat: true},

		{fromFormat: "gif", toFormat: "gif", resutlErr: true, resultFromFormat: "", resultToFormat: "", resultSetFormat: false},
		{fromFormat: "gif", toFormat: "png", resutlErr: false, resultFromFormat: "gif", resultToFormat: "png", resultSetFormat: true},
		{fromFormat: "gif", toFormat: "jpg", resutlErr: false, resultFromFormat: "gif", resultToFormat: "jpg", resultSetFormat: true},

		{fromFormat: "hoge", toFormat: "jpg", resutlErr: true, resultFromFormat: "gif", resultToFormat: "jpg", resultSetFormat: false},
		{fromFormat: "gif", toFormat: "fuga", resutlErr: true, resultFromFormat: "gif", resultToFormat: "jpg", resultSetFormat: false},
	}
	for _, c := range testCases {
		err := ic.SetConvertFormat(c.fromFormat, c.toFormat)

		if err == nil {
			if c.resutlErr {
				t.Fatal("fromFormat:" + c.fromFormat + " toFormat:" + c.toFormat + " err")
			}
			if c.fromFormat != c.resultFromFormat {
				t.Fatal("fromFormat:" + c.fromFormat + " toFormat:" + c.toFormat + " resultFromFormat err")
			}
			if c.toFormat != c.resultToFormat {
				t.Fatal("fromFormat:" + c.fromFormat + " toFormat:" + c.toFormat + " resultToFormat err")
			}
			if c.resultSetFormat == false {
				t.Fatal("fromFormat:" + c.fromFormat + " toFormat:" + c.toFormat + " resultSetFormat err")
			}
		} else {
			if c.resutlErr == false {
				t.Fatal("fromFormat:" + c.fromFormat + " toFormat:" + c.toFormat + " err")
			}
			if c.resultSetFormat == true {
				t.Fatal("fromFormat:" + c.fromFormat + " toFormat:" + c.toFormat + " resultSetFormat err")
			}
		}

	}
}

func TestConvert(t *testing.T) {
	var ic ImgConverter

	testCases := []struct {
		fromFormat string
		toFormat   string
		path       string
		err        error
		result     []Result
	}{
		{fromFormat: "jpg", toFormat: "png", path: "../test_data/jpg/dir1", err: nil,
			result: []Result{{Msg: "../test_data/jpg/dir1/dir_test1.jpg -> ../test_data/jpg/dir1/dir_test1.png", Err: nil}}},
		{fromFormat: "jpg", toFormat: "gif", path: "../test_data/jpg/dir1", err: nil,
			result: []Result{{Msg: "../test_data/jpg/dir1/dir_test1.jpg -> ../test_data/jpg/dir1/dir_test1.gif", Err: nil}}},

		{fromFormat: "png", toFormat: "jpg", path: "../test_data/png/dir1", err: nil,
			result: []Result{{Msg: "../test_data/png/dir1/dir_test1.png -> ../test_data/png/dir1/dir_test1.jpg", Err: nil}}},
		{fromFormat: "png", toFormat: "gif", path: "../test_data/png/dir1", err: nil,
			result: []Result{{Msg: "../test_data/png/dir1/dir_test1.png -> ../test_data/png/dir1/dir_test1.gif", Err: nil}}},

		{fromFormat: "gif", toFormat: "png", path: "../test_data/gif/dir1", err: nil,
			result: []Result{{Msg: "../test_data/gif/dir1/dir_test1.gif -> ../test_data/gif/dir1/dir_test1.png", Err: nil}}},
		{fromFormat: "gif", toFormat: "jpg", path: "../test_data/gif/dir1", err: nil,
			result: []Result{{Msg: "../test_data/gif/dir1/dir_test1.gif -> ../test_data/gif/dir1/dir_test1.jpg", Err: nil}}},

		{fromFormat: "gif", toFormat: "gif", path: "../test_data/gif/dir1", err: errors.New("not set format"),
			result: []Result{{Msg: "../test_data/gif/dir1/dir_test1.gif -> ../test_data/gif/dir1/dir_test1.jpg", Err: nil}}},

		{fromFormat: "jpg", toFormat: "png", path: "hoge", err: errors.New("target file path is not exist"),
			result: []Result{{Msg: "../test_data/jpg/dir1/dir_test1.jpg -> ../test_data/jpg/dir1/dir_test1.png", Err: nil}}},

		{fromFormat: "jpg", toFormat: "png", path: "../test_data/jpg/dir3", err: nil,
			result: []Result{{Msg: "../test_data/jpg/dir3/dummy.jpg -> ../test_data/jpg/dir3/dummy.png", Err: errors.New("input file decode error")}}},
	}

	for _, c := range testCases {
		ic.SetConvertFormat(c.fromFormat, c.toFormat)
		rs, err := ic.Convert(c.path)

		if err != nil && fmt.Sprintf("%s", err) != fmt.Sprintf("%s", c.err) {
			fmt.Println(err)
			t.Fatal("fromFormat:" + c.fromFormat + " toFormat:" + c.toFormat + " path:" + c.path + "err")
		}

		i := 0
		for _, d := range rs {
			if c.result[i].Msg != d.Msg {
				t.Fatal("fromFormat:" + c.fromFormat + " toFormat:" + c.toFormat + " path:" + c.path + " result Msg err")
			}
			if d.Err != nil && fmt.Sprintf("%s", c.result[i].Err) != fmt.Sprintf("%s", d.Err) {
				t.Fatal("fromFormat:" + c.fromFormat + " toFormat:" + c.toFormat + " path:" + c.path + " result err")
			}
			i++
		}
	}

}
func TestIsTargetFile(t *testing.T) {
	testCases := []struct {
		fromFormat string
		fileExt    string
		result     bool
	}{
		{fromFormat: "jpg", fileExt: ".jpg", result: true},
		{fromFormat: "jpg", fileExt: ".jpeg", result: true},
		{fromFormat: "jpg", fileExt: ".png", result: false},

		{fromFormat: "gif", fileExt: ".gif", result: true},
		{fromFormat: "gif", fileExt: ".jpg", result: false},
		{fromFormat: "gif", fileExt: ".png", result: false},

		{fromFormat: "png", fileExt: ".png", result: true},
		{fromFormat: "png", fileExt: ".jpg", result: false},
		{fromFormat: "png", fileExt: ".gif", result: false},
	}

	for _, c := range testCases {
		if isTargetFile(c.fromFormat, c.fileExt) != c.result {
			t.Fatal("fromFormat:" + c.fromFormat + " fileExt:" + c.fileExt + " failed test")
		}
	}
}

func TestConvertTo(t *testing.T) {
	var sic singleImgConverter
	testCases := []struct {
		fromFilePath string
		toFilePath   string
		toFormat     string
		err          error
	}{
		{fromFilePath: "../test_data/jpg/dir1/dir_test1.jpg", toFilePath: "../test_data/jpg/dir1/dir_test1.png", toFormat: "png", err: nil},
		{fromFilePath: "../test_data/jpg/dir1/dir_test1.jpg", toFilePath: "../test_data/jpg/dir1/dir_test1.gif", toFormat: "gif", err: nil},
		{fromFilePath: "../test_data/jpg/dir1/dir_test1.png", toFilePath: "../test_data/jpg/dir1/dir_test1.jpg", toFormat: "jpg", err: nil},
		{fromFilePath: "../test_data/jpg/dir1/dir_test1.png", toFilePath: "../test_data/jpg/dir1/dir_test1.gif", toFormat: "gif", err: nil},
		{fromFilePath: "../test_data/jpg/dir1/dir_test1.gif", toFilePath: "../test_data/jpg/dir1/dir_test1.png", toFormat: "png", err: nil},
		{fromFilePath: "../test_data/jpg/dir1/dir_test1.gif", toFilePath: "../test_data/jpg/dir1/dir_test1.jpg", toFormat: "jpg", err: nil},
		{fromFilePath: "hoge.gif", toFilePath: "hoge.jpg", toFormat: "jpg", err: errors.New("input file open error")},
		{fromFilePath: "../test_data/jpg/dummy.txt", toFilePath: "../test_data/jpg/dir1/dir_test1.png", toFormat: "png", err: errors.New("input file decode error")},
		{fromFilePath: "../test_data/jpg/dir1/dir_test1.jpg", toFilePath: "", toFormat: "png", err: errors.New("output file create error")},
	}

	for _, c := range testCases {
		sic.fromFilePath = c.fromFilePath
		sic.toFilePath = c.toFilePath
		sic.toFormat = c.toFormat
		err := sic.convertTo()
		if err != nil && fmt.Sprintf("%s", err) != fmt.Sprintf("%s", c.err) {
			t.Fatal("fromFilePath: " + c.fromFilePath + " toFilePath: " + c.toFilePath + "err")
		}
	}
}
