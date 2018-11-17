GoImageConverter
=====

# Overview

GoImageConverterはGolangで書かれた画像コンバーターです。

# SetUp

下記のようにコマンドを叩くと、実行形式のimconvファイルが生成されます
```
cd kadai2
make build
```

# Usage

実行時の引数は下記の通りです。
オプションの後に、ディレクトリ名を引数として与えてください
```
./imconv [OPTIONS] dirname
  -f string
    	input image file format (default "jpg")
  -t string
    	output image file format (default "png")
Supported formats are [jpg jpeg png]
```

# 課題に対するアプローチ
## 1回目の宿題のテストを作ってみて下さい
テーブル駆動テストを行う
基本的にほぼテーブル駆動でテストは作成

テストヘルパーを作ってみる
あまりヘルパーにしたいところが見当たらなかった
setup/teardown系は*testing.Tのヘルパー関数群を利用している

テストのしやすさを考えてリファクタリングしてみる
https://mattn.kaoriya.net/software/lang/go/20161025113154.htm
https://deeeet.com/writing/2014/12/18/golang-cli-test/
CLIのインターフェイス部がテストしづらくなっていたので、上記を参考に、リファクタリングを行った
主にやったことは
- Exit処理をテストから上書き可能にするために、エイリアスを与え、テストではExit処理をラップして、エラーコードを取れるようにした
- CLIの出力先に任意のio.Writerを与えられるように構造体を定義し、実行環境ではos.Std、テストではbytes.Bufferに出力するように変更している

テストのカバレッジを取ってみる
make test > coverage がみれる
make cover > プロファイルで通っていないところがみれる
というコマンドを定義して確認した
概ね全体をカバーしたが、いくつか通しづらいところがあったので、こちらはプルリクでコメント補足する

## io.Readerとio.Writerについて調べてみよう
標準パッケージでどのように使われているか
io.Readerとio.Writerがあることで
どういう利点があるのか具体例を挙げて考えてみる

https://github.com/golang/go/tree/master/src
をio.Reader、io.Writerで検索してみると

binary、csv、http、バッファ、XML、アーカイブファイル、標準入出力、etc
多様なIOフォーマットの読み書きで、基底となる共通のIFとして使われている

例としてpngのreaderを取り出して見てみると、

image/png/reader.go

// Decode reads a PNG image from r and returns it as an image.Image.
func Decode(r io.Reader) (image.Image, error) {
PNGを扱うライブラリを実装しているにも関わらず、入力をPNG特有のなにかとして扱っていない

// 実際に処理しているところ d.r がio.Reader
n, err := io.ReadFull(d.r, d.tmp[:8])
読み出しは、同じ規約(io.Reader)を持った標準ライブラリ関数に処理させている

IOに関して固有のロジックを実装せずにすむ (読み出した後に、PNG特有のデータ構造として考える)
同じ規約を持ったライブラリに、まるっとIO処理を委譲できる(取り回しがしやすい)
そもそもここ、どういう風に引数作ったらいいんだろう？を悩む必要がない
-> 作る側として便利

image/png/reader_test.go

// PNGのテストだが、stringから読み出させている
const (
  ihdr = "\x00\x00\x00\x0dIHDR\x00\x00\x00\x01\x00\x00\x00\x02\x08\x00\x00\x00\x00\xbc\xea\xe9\xfb"
  idat = "\x00\x00\x00\x0eIDAT\x78\x9c\x62\x62\x00\x04\x00\x00\xff\xff\x00\x06\x00\x03\xfa\xd0\x59\xae"
  iend = "\x00\x00\x00\x00IEND\xae\x42\x60\x82"
)
Decode(strings.NewReader(pngHeader + ihdr + idatWhite + idatZero + iend))

// PNGのテストだが、バイト列から読み出させている
_, err := Decode(bytes.NewReader([]byte{
  0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d, 0x49, 0x48, 0x44, 0x52,
  0x7f, 0xff, 0xff, 0xfe, 0x7f, 0xff, 0xff, 0xfe, 0x08, 0x06, 0x00, 0x00, 0x00, 0x30, 0x57, 0xb3,
  0xfd, 0x00, 0x00, 0x00, 0x15, 0x49, 0x44, 0x41, 0x54, 0x78, 0x9c, 0x62, 0x62, 0x20, 0x12, 0x8c,
  0x2a, 0xa4, 0xb3, 0x42, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x13, 0x38, 0x00, 0x15, 0x2d, 0xef,
  0x5f, 0x0f, 0x00, 0x00, 0x00, 0x00, 0x49, 0x45, 0x4e, 0x44, 0xae, 0x42, 0x60, 0x82,
}))

Readerインターフェイスを実装していればいいので、IO元/先をPNGファイルにこだわらずにすむ (テストしやすい)
IOする必要があるライブラリを使おうとするたびに個別の規約を覚える必要がない
同じIO IFを提供しているライブラリに読み出しを任せられる(取り回しがしやすい)
-> 使う側として便利

何かを読んでデータをp []byteに書き込むという振る舞いが満たされていることを共通の規約にすることで、
実装の詳細がどんな性質であるかを意識することなく、上記にあげたような諸々の便利さを手に入れることができている
func Read(p []byte) (n int, err error)
