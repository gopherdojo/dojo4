## 課題1
* 次の仕様を満たすコマンドを作って下さい
  - ディレクトリを指定する
  - 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
  - ディレクトリ以下は再帰的に処理する
  - 変換前と変換後の画像形式を指定できる（オプション）

* 以下を満たすように開発してください
  - mainパッケージと分離する
  - 自作パッケージと標準パッケージと準標準パッケージのみ使う
    - 準標準パッケージ：golang.org/x以下のパッケージ
  - ユーザ定義型を作ってみる
  - GoDocを生成してみる

#### 課題2

  - [x] テストのしやすさを考えてリファクタリングしてみる
  - [x] テストのカバレッジを取ってみる
  ```
  go test -coverprofile=coverage.out
  go tool cover -func=coverage.out
  go tool cover -html=coverage.out
  ```
  - [x] テーブル駆動テストを行う
  - [x] テストヘルパーを作ってみる

#### io.Readerとio.Writerを調べる

- 標準パッケージでどのように使われているか
  - ファイルの読み書きなど入出力の基本使います
  - バッファリングして読み書きを行うところを使う
  - 多くの標準パッケージがインタフェースを実装していたり、引数として扱える形でサポートしています。
  ```
  json
  bytes.Buffer
  bufio.Reader
  os.File
  image
  jpeg
  png
  base64
  ```
- io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる 
　- 標準ライブラリに用意されている便利なツールでラップして呼び出すこと
　- モックしやすいでテスト書きやすくなります

## コマンド
* jpeg, png, jpg, gifを対応しました。
* デコード出来ない場合はログを出して、次の処理へ進みます。
* GoDocを生成してみる

## ビルド
```
$ make install
```
## テスト
```
$ make test
```

## コマンド使い方
```
$./bin/kadai2 [options] [directories]
```

### オプション
```
-i string
    Input file type (default "jpg")

-o string
    Output file type (default "png")
```

### 例
```
$./bin/kadai2 -i jpg -o png fixtures
```

## Godoc
```
$godoc -http=:6060
```
以下のURLで読めます。
`http://localhost:6060/pkg/github.com/gopherdojo/dojo4/kadai2/phamvanhung2e123/converter`
