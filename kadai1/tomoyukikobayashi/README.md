GoImageConverter
=====

# Overview

GoImageConverterはGolangで書かれた画像コンバーターです。

# SetUp

下記のようにコマンドを叩くと、実行形式のimconvファイルが生成されます
```
cd kadai1/src/
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

オプションについて知りたい場合は、下記のようにしてヘルプを参照してください
```
./imconv --help
```
# 課題に対するアプローチ
## 次の仕様を満たすコマンドを作って下さい
- [x] ディレクトリを指定する
　main.go
  特に工夫なくflagで指定
- [x] 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
　image_converter.go で以下の流れで実現
  入力ファイルを開く
  入力ファイルをimageパッケージのDecode関数でデコード
  出力ファイルを開く
  出力ファイルに指定したフォーマットのimageサブパッケージのEncode関数でエンコード
  deferで入力、出力ファイルをクローズ
- [x] ディレクトリ以下は再帰的に処理する
  file_finder.go
  ioutil.ReadDir で []os.FileInfo を取得して、FileInfoがdirectoryだったら
  再帰的に処理するようにした
- [x] 変換前と変換後の画像形式を指定できる（オプション）
  image_converter.go
  switch条件で、指定したフォーマットにマッチしたImageサブパッケージでEncodeするようにした
## 以下を満たすように開発してください
- [x] mainパッケージと分離する
  /file 汎用的なファイル操作
  /imconv 画像処理特有のデータと操作
  の2パッケージを作成
- [x] 自作パッケージと標準パッケージと準標準パッケージのみ使う
- [x] 準標準パッケージ：golang.org/x以下のパッケージ  
  標準外パッケージは特に使わなかった
- [x] ユーザ定義型を作ってみる
  file_finder.go
  type structは使う場面がなかったので、[]string 型をラップして、レシーバを定義した
- [x] GoDocを生成してみる
　 Packageとパブリックには全てGoDocコメントをつけて、godocコマンドで表示を確認した