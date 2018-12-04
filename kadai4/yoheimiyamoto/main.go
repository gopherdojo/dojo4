package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/YoheiMiyamoto/dojo4/kadai4/yoheimiyamoto/omikuji"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	result := omikuji.Play()
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Fatal(err)
		return
	}
}
