package main

import (
	"net/http"

	"github.com/gopherdojo/dojo4/kadai4/akuchii/fortune"
)

func main() {
	f := fortune.NewFortune(fortune.DefaultClock{})
	http.HandleFunc("/", f.Handler)
	http.ListenAndServe(":8080", nil)
}
