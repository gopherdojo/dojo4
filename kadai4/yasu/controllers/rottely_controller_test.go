package controllers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func DevineFortuneTest(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/rottely", nil)
	DevineFortune(w, r)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
		t.Fatal("unexpected error")
	}
	if fortune := string(bytes); fortune == "" {
		t.Fatalf("unexpected response: %s", fortune)
	}
}
