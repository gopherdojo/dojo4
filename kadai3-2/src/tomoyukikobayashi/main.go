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

	"golang.org/x/sync/errgroup"
)

// CLIのExitコード
const (
	ExitSuccess     = 0
	ExitError       = 1
	ExitInvalidArgs = 2
)

// HACK 雑に作ったので弱いバリデーション、パラメタ決め打ちと、処理されないエラーがある
// mainに全部押し込んで、テストもない

// CLIツールのエントリーポイント
func main() {
	if len(os.Args) < 2 {
		fmt.Printf("missing URL parameter")
		os.Exit(ExitInvalidArgs)
	}

	url := os.Args[1]
	fmt.Printf("start to download %v\n", url)

	/// Rangeアクセスの準備
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

	/// 処理するプロセス数を決める
	// HACK とりあえず決めうち、本当は環境or引数からProc数決めるのが良い
	procs := 4
	size, _ := strconv.Atoi(hRes.Header["Content-Length"][0])
	fmt.Printf("donwload size %v, parallel %v", size, procs)

	/// 一次ファイルを格納するフォルダの作成

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

	/// goroutine起こして並列ダウンロード
	var eg errgroup.Group
	// bytes のレンジに指定するバイト位置
	from, to := 0, 0
	for i := 0; i < procs; i++ {
		// 各分割リクエストの取得バイト範囲を決める
		to = from + size/procs
		if i == procs-1 {
			to = to + size%procs
		}

		// 平行にDLをする
		i := i + 1
		eg.Go(func() error {
			fn := filepath.Join(dlTmp, strconv.FormatInt(int64(i), 10))

			// ファイルが存在していたらスキップする
			// TODO 本当はファイルサイズがto - fromと一致するかも見た方が良い
			if _, err := os.Stat(fn); err == nil {
				return nil
			}

			// Rangeを指定してGETリクエスト作成
			dReq, err := http.NewRequest("GET", url, nil)
			if err != nil {
				return err
			}
			dReq.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", from, to))

			// DLリクエスト
			dRes, err := http.DefaultClient.Do(dReq)
			if err != nil {
				return err
			}

			// 取得したデータをBufに読み出し
			buf, err := ioutil.ReadAll(dRes.Body)
			if err != nil {
				return err
			}

			// 分割ファイルの保存
			file, err := os.Create(fn)
			if err != nil {
				return err
			}
			defer file.Close()

			// TODO ここで落ちるとファイルが存在するけどDLできていないのにスキップしてしまう
			_, err = file.Write(buf)
			if err != nil {
				return err
			}

			return nil
		})

		from = to + 1
	}
	// 全プロセスがDL終わるまで待つ
	if err := eg.Wait(); err != nil {
		fmt.Printf("error occurred while downloading %v", err)
		os.Exit(ExitError)
	}

	/// DLした部分ファイルを結合する
	// deferをos.Existと同じブロックに置かないように無名関数にしている
	// (でもdefer使わないでシーケンシャルにやるのと変わらんなこれだと)
	err = func() error {
		// 複数ファイルを一気に読み出すReaderを作成
		files := make([]io.Reader, procs)
		for i := 0; i < procs; i++ {
			file, err := os.Open(filepath.Join(dlTmp, strconv.FormatInt(int64(i+1), 10)))
			// 閉じるの失敗しても実害ないので、エラー処理なし
			defer file.Close()
			if err != nil {
				fmt.Printf("error occurred while reading tmp file %v", err)
				os.Exit(ExitError)
			}
			files[i] = file
		}
		reader := io.MultiReader(files...)
		b, err := ioutil.ReadAll(reader)
		if err != nil {
			return err
		}

		/// tmpフォルダのゴミ掃除(失敗しても特に何もしない))
		defer os.RemoveAll(dlTmp)

		// ファイルを結合する
		file, err := os.Create(filepath.Base(url))
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = file.Write(b)
		if err != nil {
			return err
		}

		return nil
	}()

	if err != nil {
		fmt.Printf("error occurred while creating dl file %v", err)
		os.Exit(ExitError)
	}

	os.Exit(ExitSuccess)
}
