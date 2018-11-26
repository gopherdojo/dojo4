package imconv

import (
	"os"
	"strings"
	"testing"
)

// fixture
var (
	dataRoot                = "../testdata"
	jpgF, jpegF, pngF, txtF = "file1.jpg", "file2.jpeg", "file3.png", "file4.txt"
	invalidJpgF             = "invalid.jpg"
)

// 前提ファイルがあるかチェック
func checkPrecondition(t *testing.T) {
	// testdataの存在確認
	if _, err := os.Stat(dataRoot); err != nil {
		t.Fatal(err)
	}

	// 元になるテストデータの存在確認
	for _, v := range []string{jpgF, jpegF, pngF, txtF, invalidJpgF} {
		if _, err := os.Stat(dataRoot + "/" + v); err != nil {
			t.Fatal(err)
		}
	}

	return
}

func Test_Find(t *testing.T) {
	// 指定した条件でフォルダとファイルを配置する
	checkPrecondition(t)

	tests := []struct {
		name  string
		path  string
		from  string
		to    string
		want  string
		isErr bool
	}{
		// TOOD 毎回全埋めすると長くなるので、composition使ってデフォルト値+上書きにしてしまうと短くできるかも？
		{
			name:  "jpgToPng",
			path:  path(jpgF),
			from:  "jpg",
			to:    "png",
			want:  newPath(jpgF, "png"),
			isErr: false,
		},
		{
			name:  "pngToJpg",
			path:  path(pngF),
			from:  "png",
			to:    "jpg",
			want:  newPath(pngF, "jpg"),
			isErr: false,
		},
		{
			name:  "unSupportedFrom",
			path:  path(jpgF),
			from:  "INVALID",
			to:    "png",
			want:  "",
			isErr: true,
		},
		{
			name:  "unSupportedTo",
			path:  path(jpgF),
			from:  "jpg",
			to:    "INVALID",
			want:  "",
			isErr: true,
		},
		{
			name:  "fromAndToAreEquals",
			path:  path(jpgF),
			from:  "jpg",
			to:    "jpg",
			want:  "",
			isErr: true,
		},
		{
			name:  "failedToOpenFile",
			path:  "notExistFile.jpg",
			from:  "jpg",
			to:    "png",
			want:  "",
			isErr: true,
		},
		{
			name:  "failedToDecodeFile",
			path:  path(invalidJpgF),
			from:  "jpg",
			to:    "png",
			want:  "",
			isErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Convert(tt.path, tt.from, tt.to)
			// TODO error型定義していないので有無の判断だけになっている
			// TOOD ここちょっと長いのでGiven When Then (TearDown)の見通しが悪くなっている
			if tt.isErr {
				if err == nil {
					t.Fatal(err)
				}
				return
			}
			// 終わったらゴミ掃除
			defer func() {
				if err := os.Remove(got); err != nil {
					t.Error(err)
				}
			}()
			if got != tt.want {
				t.Fatalf("want = %#v, got = %#v", tt.want, got)
			}
		})
	}
}

func path(file string) string {
	return dataRoot + "/" + file
}

func newPath(file string, to string) string {
	parts := strings.Split(file, ".")
	newPath := dataRoot + "/" + strings.Join(parts[:len(parts)-1], ".") + "." + to
	return newPath
}
