package fortune

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// index of values
const (
	excellent = iota
	great
	good
	bad
)

var values = []string{
	"大吉", "中吉", "小吉", "凶",
}

// Clock return current time
type Clock interface {
	Now() time.Time
}

// ClockFunc is an adapter to implement Clock interface
type ClockFunc func() time.Time

// Now calls f()
func (f ClockFunc) Now() time.Time {
	return f()
}

// Fortune implement Clock interfate
type Fortune struct {
	Clock Clock
	Err error
}

// Result is type for JSON encode/decode
type Result struct {
	Fortune string `json:"fortune"`
}

// Encode converts Result into JSON string
func Encode(r Result) (string, error) {
	b := bytes.Buffer{}
	e := json.NewEncoder(&b)

	err := e.Encode(&r)
	if err != nil {
		return "", err
	}

	return string(b.Bytes()), nil
}

// Decode converts JSON string into Result
func Decode(s string) (Result, error) {
	b := bytes.NewBufferString(s)
	d := json.NewDecoder(b)

	var r Result
	err := d.Decode(&r)
	if err != nil {
		return r, err
	}

	return r, nil
}

// Handler http.HandleFunc implementation
func (f *Fortune) Handler(w http.ResponseWriter, r *http.Request) {
	var res Result
	now := f.Clock.Now()

	if now.Month() == time.January && (1 <= now.Day() && now.Day() <= 3) {
		// always return excellent
		res.Fortune = values[excellent]
	} else {
		// random pick up
		rand.Seed(now.Unix())
		i := rand.Int() % len(values)
		res.Fortune = values[i]
	}

	json, err := Encode(res)
	if err != nil {
		f.Err = err
	} else {
		fmt.Fprintf(w, json)
	}
}
