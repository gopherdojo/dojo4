# 第1回課題　画像変換コマンドの作成
## コマンド実行例
```
./main -dir={パス} -srcFormat=jpg -destFormat=png
```
imagesディレクトリ内の -f オプションで指定されたフォーマットのファイルが -t オプションで指定されたファイルに変換される。

## 補足
* コマンドライン引数 destFormat は必須項目
* 今回は jpg, jpeg, png, gif のみ変換対応