package controllers

import (
	"dojo4/kadai4/yasu/controllers/serializer"
	"dojo4/kadai4/yasu/services"
	"encoding/json"
	"net/http"
)

func DevineFortune(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fortune, err := services.Rottely()
	if err != nil {
		response := serializer.Error{Err: err}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := serializer.RottelyResult{Fortune: fortune}
	json.NewEncoder(w).Encode(response)
}
