package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Omikuji interface {
	Do() (string, error)
}

type OmikujiController struct {
	Omikuji Omikuji
}

func NewOmikujiController(omikuji Omikuji) *OmikujiController {
	return &OmikujiController{Omikuji: omikuji}
}

func (o OmikujiController) Get(w http.ResponseWriter, r *http.Request) {
	result, err := o.Omikuji.Do()
	w.Header().Add("Content-Type", "application/json")
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	// error response
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resultFormat := struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		}
		enc.Encode(resultFormat)
		fmt.Fprint(w, buf.String())
		return
	}
	// normal response
	w.WriteHeader(http.StatusOK)
	resultFormat := struct {
		Result string `json:"result"`
	}{
		Result: result,
	}
	enc.Encode(resultFormat)
	fmt.Fprint(w, buf.String())
}
