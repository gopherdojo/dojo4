package main

import (
	"net/http"

	"tomoyukikobayashi/handler"
)

func main() {
	http.HandleFunc("/", handler.Fortune)
	http.ListenAndServe(":8080", nil)
}
