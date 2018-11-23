/*
Pacakge main is the entry point of this project.
This mainly provides interaction logics and parameters
used in CLI intrerfaces.
*/
package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// CLIのExitコード
const (
	ExitSuccess     = 0
	ExitError       = 1
	ExitInvalidArgs = 2
)

// HACK 雑に作ったので弱いバリデーション、パラメタ決め打ちとか、処理されないerrがある
// mainに全部押し込んで、テストもない

// CLIツールのエントリーポイント
func main() {
	if len(os.Args) < 2 {
		fmt.Printf("missing URL parameter")
		os.Exit(ExitInvalidArgs)
	}

	url := os.Args[1]
	fmt.Printf("start to download %v\n", url)

	// HEADで取れた情報から、対象のAccept-Ranges対応確認と、ファイルサイズの取得
	hRes, err := http.Head(url)
	if err != nil {
		fmt.Printf("cannot get header %v from %v\n", err, url)
		os.Exit(ExitError)
	}
	if _, ok := hRes.Header["Accept-Ranges"]; !ok {
		fmt.Printf("range access does not supported")
		os.Exit(ExitError)
	}
	// HACK とりあえず決めうち、本当は環境or引数からProc数決めるのが良い
	procs := 3
	size, _ := strconv.Atoi(hRes.Header["Content-Length"][0])
	fmt.Printf("donwload size %v, parallel %v", size, procs)

	// 一時ファイルを格納するディレクトリ名(とりあえずハッシュ)を決める
	h := sha1.Sum([]byte(url))
	hash := fmt.Sprintf("%x", h)

	// DL共通のtmpディレクトリと、個別のDL用のhash名のディレクトリを掘る
	cur, _ := os.Getwd()
	// tmpフォルダがなければ作る
	if _, err := os.Stat(filepath.Join(cur, "tmp")); os.IsNotExist(err) {
		err := os.Mkdir("tmp", 0777)
		if err != nil {
			fmt.Printf("failed to create tmp file dir")
			os.Exit(ExitError)
		}
	}
	// tmp配下にハッシュにマッチするフォルダがなければ作る
	dlTmp := filepath.Join("tmp", hash)
	if _, err := os.Stat(filepath.Join(cur, dlTmp)); os.IsNotExist(err) {
		err := os.Mkdir(dlTmp, 0777)
		if err != nil {
			fmt.Printf("failed to create tmp file dir")
			os.Exit(ExitError)
		}
	}

	// goroutine起こして並列ダウンロード
	// bytes のレンジに指定するバイト位置
	from, to := 0, 0
	for i := 0; i < procs; i++ {
		// 各分割リクエストの取得バイト範囲を決める
		to = from + size/procs
		if i == procs-1 {
			to = to + size%procs
		}

		// TODO ファイル存在確認して、いたらDLしない
		// TOOD waitして最後にファイルを結合
		// TOOD erros使う
		dReq, _ := http.NewRequest("GET", url, nil)
		dReq.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", from, to))
		dRes, _ := http.DefaultClient.Do(dReq)
		buf, _ := ioutil.ReadAll(dRes.Body)
		// TODO 存在とファイルサイズ確認 tmpフォルダ作成
		fn := filepath.Join(dlTmp, strconv.FormatInt(int64(i), 10))
		file, _ := os.Create(fn)
		defer file.Close()

		_, _ = file.Write(buf)

		from = to + 1
	}

	// DLした部分ファイルを結合する
	files := make([]io.Reader, procs)
	for i := 0; i < procs; i++ {
		file, _ := os.Open(filepath.Join(dlTmp, strconv.FormatInt(int64(i), 10)))
		files[i] = file
	}

	reader := io.MultiReader(files...)
	// TOOD urlから作る
	file, _ := os.Create("out.jpg")
	defer file.Close()
	b, _ := ioutil.ReadAll(reader)
	_, _ = file.Write(b)

	// TODO tmpファイルの掃除

	os.Exit(ExitSuccess)
}
