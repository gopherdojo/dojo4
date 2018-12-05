package main

import (
	"net/http"

	"github.com/YoheiMiyamoto/dojo4/kadai4/yoheimiyamoto/omikuji"
)

func main() {
	http.HandleFunc("/", omikuji.Handler)
	http.ListenAndServe(":8080", nil)
}
