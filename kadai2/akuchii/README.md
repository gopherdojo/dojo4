# io.Readerとio.Writerについて調べてみよう

## 標準パッケージでどのように使われているか
os.File, os.Stdin, os.Stdout, bytes.buffersなどデータの読み書きを行う処理の部分に使われている

## io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる
例えばデータを読み書きする処理があったときに、io.Reader, io.Writerがあることで使う側は内部実装を知らなくても使うことができる。
内部実装を知らずに使えるので、読み書き先をDBからjson, csvなどに切り替えたい時も容易に行うことができる。
