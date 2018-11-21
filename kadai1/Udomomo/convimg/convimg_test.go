package convimg

import (
	"image"
	"os"
	"testing"
)

func TestSearchFile(t *testing.T) {
	t.Helper()
	expProcessingPath := []string{
		"/Users/naoya/Golang_practice/dojo4/kadai1/Udomomo/testdata/test.jpg",
		"/Users/naoya/Golang_practice/dojo4/kadai1/Udomomo/testdata/test2/test2.jpg",
	}

	processingPath := searchFile("/Users/naoya/Golang_practice/dojo4/kadai1/Udomomo/testdata", ".jpg", ".png", make([]string, 0))
	for _, p := range expProcessingPath {
		if contains(processingPath, p) == false {
			t.Errorf("searchFile failed: %s is not searched", p)
		}
	}

	if len(processingPath) != len(expProcessingPath) {
		t.Errorf("searchFile failed: result is longer than expected: %v", processingPath)
	}
}

func contains(su []string, fl string) bool {
	for _, s := range su {
		if s == fl {
			return true
		}
	}
	return false
}

func TestValidateIfConvNeeded(t *testing.T) {
	t.Helper()
	jpgPath := "/Users/naoya/Golang_practice/dojo4/kadai1/Udomomo/testdata/test.jpg"
	if validateIfConvNeeded(jpgPath, ".jpg", ".gif") == false {
		t.Error("validate failed: file should be converted")
	}
	if validateIfConvNeeded(jpgPath, ".png", ".gif") == true {
		t.Error("validate failed: fromExt difference is ignored")
	}
	if validateIfConvNeeded(jpgPath, ".jpg", ".jpg") == true {
		t.Error("validate failed: from and to are equal")
	}
}

func TestConvFile(t *testing.T) {
	t.Helper()
	var testCase = []struct {
		path    string
		newPath string
		toExt   string
		expFmt  string
		wantErr bool
	}{
		//正常ケース png->jpg
		{"/Users/naoya/Golang_practice/dojo4/kadai1/Udomomo/testdata/test.png",
			"/Users/naoya/Golang_practice/dojo4/kadai1/Udomomo/output/test.jpg",
			".jpg",
			"jpeg",
			false,
		},
		//正常ケース jpg->gif
		{"/Users/naoya/Golang_practice/dojo4/kadai1/Udomomo/testdata/test.jpg",
			"/Users/naoya/Golang_practice/dojo4/kadai1/Udomomo/output/test.gif",
			".gif",
			"gif",
			false,
		},
		//正常ケース gif->png
		{"/Users/naoya/Golang_practice/dojo4/kadai1/Udomomo/testdata/test.gif",
			"/Users/naoya/Golang_practice/dojo4/kadai1/Udomomo/output/test.png",
			".png",
			"png",
			false,
		},
		//異常ケース 存在しないファイルパスの場合
		{"/Users/naoya/Golang_practice/dojo4/kadai1/Udomomo/testdata/hogehoge.gif",
			"/Users/naoya/Golang_practice/dojo4/kadai1/Udomomo/output/test.png",
			".png",
			"png",
			true,
		},
		//異常ケース ファイルが壊れている場合
		{"/Users/naoya/Golang_practice/dojo4/kadai1/Udomomo/testdata/broken.gif",
			"/Users/naoya/Golang_practice/dojo4/kadai1/Udomomo/output/test.png",
			".png",
			"png",
			true,
		},
		//異常ケース 書き込み先の権限がない場合
		{"/Users/naoya/Golang_practice/dojo4/kadai1/Udomomo/testdata/test.gif",
			"/Users/naoya/Golang_practice/dojo4/kadai1/Udomomo/output/permission/test.png",
			".png",
			"png",
			true,
		},
		//異常ケース 変換先の拡張子が不正な場合
		{"/Users/naoya/Golang_practice/dojo4/kadai1/Udomomo/testdata/test.gif",
			"/Users/naoya/Golang_practice/dojo4/kadai1/Udomomo/output/test.pmg",
			".pmg",
			"png",
			true,
		},
	}

	if err := os.Mkdir("/Users/naoya/Golang_practice/dojo4/kadai1/Udomomo/output", 0755); err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll("/Users/naoya/Golang_practice/dojo4/kadai1/Udomomo/output")

	//書き込み先に権限がない場合のテスト用にディレクトリを作っておく
	if err := os.Mkdir("/Users/naoya/Golang_practice/dojo4/kadai1/Udomomo/output/permission", 0555); err != nil {
		t.Fatal(err)
	}

	for _, tc := range testCase {
		processedPath, err := convFile(tc.path, tc.newPath, tc.toExt)
		if err != nil && tc.wantErr == false {
			t.Errorf("Case %s should succeed but returned error: %#v", tc.path, err)
		}

		if err == nil && tc.wantErr == true {
			t.Errorf("Case %s should fail but succeeded", tc.path)
		}

		if err == nil {
			file, err := os.Open(processedPath)
			if err != nil {
				t.Errorf("processed file can't open. Path: %s", processedPath)
			}
			defer file.Close()

			if _, fmt, _ := image.Decode(file); fmt != tc.expFmt {
				t.Errorf("fmt is wrong: expected %s but actually %s", tc.expFmt, fmt)
			}
		}
	}

}
