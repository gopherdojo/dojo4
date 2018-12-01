# 画像変換器

指定したディレクトリ配下全ての画像ファイルを別の画像形式に変換するツールです。  
画像形式は入力ファイル、出力ファイルともに、「jpg,png,gif」に対応しています。  

## 使い方
usage:imgconv [option -f] [option -t] targetFilePath  
-f 入力（対象）ファイルの画像形式「jpg,png,gif」が指定できます。省略時はjpgが使用されます。  
-t 出力ファイルの画像形式「jpg,png,gif」が指定できます。省略時はpngが使用されます。  
入力ファイルの画像形式と出力ファイルの画像形式が同じ場合はエラーとなります。  

## 実行例
``` 
$ go run imgconv.go target/file/path
$ go run imgconv.go -f png -t gif target/file/path
$ go run imgconv.go -f jpg -t png target/file/path
```
