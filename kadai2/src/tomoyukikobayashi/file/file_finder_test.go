package file

import (
	"os"
	"reflect"
	"testing"
)

// fixture
var (
	dataRoot            = "../testdata"
	jpg, jpeg, png, txt = "file1.jpg", "file2.jpeg", "file3.png", "file4.txt"
	tmpRoot             = dataRoot + "/root"
	imgDir, txtDir      = tmpRoot + "/dir1/img", tmpRoot + "/dir2/txt"
)

// TODO 愚直にOS使ったが、finderでIF差し替えられるようにしておいて仮想osにやらせてもいいかも
// setup / teardown用のヘルパー関数
func testSetup(t *testing.T) func() {
	// testdataの存在確認
	if _, err := os.Stat(dataRoot); err != nil {
		t.Fatal(err)
	}

	// 元になるテストデータの存在確認
	for _, v := range []string{jpg, jpeg, png, txt} {
		if _, err := os.Stat(dataRoot + "/" + v); err != nil {
			t.Fatal(err)
		}
	}

	// テストデータを格納するフォルダの作成
	for _, v := range []string{imgDir, txtDir} {
		if err := os.MkdirAll(v, 0777); err != nil {
			t.Fatal(err)
		}
	}

	// ファイルの作成
	for _, v := range []string{imgDir + "/" +
		jpg, imgDir + "/" + jpeg, imgDir + "/" + png, txtDir + "/" + txt} {
		file, err := os.OpenFile(v, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			t.Fatal(err)
		}
		defer func() {
			if err := file.Close(); err != nil {
				t.Error(err)
			}
		}()
	}

	return func() {
		if err := os.RemoveAll(tmpRoot); err != nil {
			t.Error(err)
		}
	}
}

func Test_Find(t *testing.T) {
	// 指定した条件でフォルダとファイルを配置する
	tearDown := testSetup(t)

	tests := []struct {
		name  string
		dir   string
		exts  []string
		want  []string
		isErr bool
	}{
		{
			name: "findAll",
			dir:  tmpRoot,
			exts: []string{"jpg", "jpeg", "png", "txt"},
			want: []string{imgDir + "/" + jpg,
				imgDir + "/" + jpeg, imgDir + "/" + png, txtDir + "/" + txt},
			isErr: false,
		},
		{
			name:  "notExistExt",
			dir:   tmpRoot,
			exts:  []string{"notFound"},
			want:  []string{},
			isErr: false,
		},
		{
			name:  "cannotOpenDir",
			dir:   "notExist",
			exts:  []string{"jpg"},
			want:  nil,
			isErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Find(tt.dir, tt.exts)
			// TODO error型定義していないので有無の判断だけになっている
			if tt.isErr {
				if err == nil {
					t.Fatal(err)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("want = %#v, got = %#v", tt.want, got)
			}
		})
	}

	// 事前条件で作った構成を削除
	tearDown()
}

// TODO ザ・テストのためのテストって感じで微妙
func Test_findpaths_errcase(t *testing.T) {
	tearDown := testSetup(t)

	// 書き込みだけ許して、読み込めないdirを作成
	newDir := dataRoot + "/unReadable"
	if err := os.Mkdir(newDir, 0333); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.RemoveAll(newDir); err != nil {
			t.Error(err)
		}
	}()

	got, err := Find(dataRoot, []string{"jpg"})
	if err == nil {
		t.Fatalf("expected to get err, but got=%#v", got)
	}

	tearDown()
}
