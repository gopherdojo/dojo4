# 第1回課題　画像変換コマンドの作成
## コマンド実行例
```
go build -o imgconv
./imgconv -dir={ディレクトリパス} -srcFormat=gif -destFormat=png
```
指定したディレクトリ内の `-srcFormat` オプションで指定されたフォーマットのファイルが `-destFormat` オプションで指定されたファイルに変換される。

## コマンドライン引数
### dir
変換元の画像が格納されているディレクトリを指定

### srcFormat
* 任意項目
* 省略した場合は `jpg` となる

### destFormat
* 必須項目
* 今回は jpg, jpeg, png, gif のみ変換対応