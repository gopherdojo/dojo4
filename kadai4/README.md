OMIKUJI
=====

# Overview

OMIKUJI DAISUKI

# SetUp

下記のようにコマンドを叩くと、実行形式のserveファイルが生成されます
```
make build
```

# Usage
```
./serve
```
でポート8080で立ち上がるため下記などのようにしてリクエストしてください
```
curl localhost:8080
```
ポートを変えたい場合には、-portオプションを指定してください
```
./serve -port 8090
```
