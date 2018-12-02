package main

import (
	"net/http"

	"github.com/dais0n/dojo4/kadai4/dais0n/omikuji"

	"github.com/dais0n/dojo4/kadai4/dais0n/controller"
)

func main() {
	mux := http.NewServeMux()
	omikuji := omikuji.Omikuji{}
	c := controller.NewOmikujiController(&omikuji)
	mux.Handle("/v1/omikuji", http.HandlerFunc(c.Get))
	http.ListenAndServe(":8080", mux)
}
