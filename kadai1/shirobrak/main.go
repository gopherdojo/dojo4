package main

import (
	"log"

	"github.com/gopherdojo/dojo4/kadai1/shirobrak/imgconverter"
)

func main() {

	// 画像コンバータのインスタンス作成
	imgConverter := new(imgconverter.ImageConverter)

	// 初期化
	imgConverter.Init()

	// 画像のコンバート処理
	if err := imgConverter.Run(); err != nil {
		log.Fatal(err)
	}
}
