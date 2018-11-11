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

 ## コマンド
 * jpeg, png, jpg, gifに対応

 ## インストール
 ```
 $ go get github.com/Udomomo/dojo4/kadai1/Udomomo/convimg
 $ cd $GOPATH/src/github.com/Udomomo/dojo4/kadai1/Udomomo/convimg
 $ git checkout kadai1-Udomomo
 $ git fetch && git merge
 $ git install
```

## ビルド
 ```
 $ go build -o conv main.go
 ```

 ## コマンド
 ```
 $.conv [options] [directories]
 ```

 ### オプション
 ```
-f string
  	format before conversion (default "jpg")
-t string
   	format after conversion (default "png")
 ```