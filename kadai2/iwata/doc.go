/*
特定ディレクトリの画像を再帰的に辿って、ファイルフォーマットを変換するCLIです.

How to build

	env GO111MODULE=on go build -o imgconv

Usage

Usage of ./imgconv:
   ./imgconv [OPTIONS] DIR
OPTIONS
  -from string
    Convert from image format (default "jpg")
  -to string
    Convert to image format (default "png")

*/
package main
