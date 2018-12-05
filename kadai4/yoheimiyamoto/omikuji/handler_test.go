package omikuji

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// どうにかして、現在時刻を三が日に変更する必要がある

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	Handler(w, r)
	rw := w.Result()
	defer rw.Body.Close()
	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}
	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}
	var actual Result
	err = json.Unmarshal(b, &actual)
	if err != nil {
		t.Fatal(err)
	}
	expected := Result{"大吉"}
	if actual != expected {
		t.Fatalf("actual: %v, expected: %v", actual, expected)
	}
}
