package omikuji

import (
	"encoding/json"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	result := play()
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Fatal(err)
		return
	}
}
