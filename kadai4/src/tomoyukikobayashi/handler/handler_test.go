package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFortune(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	Fortune(w, r)

	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}
	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("could not read body")
	}
	if !isAnyFourtunes(t, b) {
		t.Fatalf("is not valid fourtune %v", string(b))
	}
}

/*
NOTE 標本増やすためにランダムな時間を与えるのをtest.quickでできないか試してみたが
reflect.Valueでpanicするので、timeはダメっぽい
https://golang.org/src/testing/quick/quick.go?s=1618:1692#L49

func TestFortune_Random(t *testing.T) {
	// テストケースが終わったらClockを元に戻す
	oc := fourtunes.Clock
	defer func() { fourtunes.Clock = oc }()

	f := func(tm time.Time) bool {
		fourtunes.Clock = ClockFunc(func() time.Time {
			return tm
		})

		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)
		Fortune(w, r)

		rw := w.Result()
		defer rw.Body.Close()
		if rw.StatusCode != http.StatusOK {
			t.Fatal("unexpected status code")
		}
		b, err := ioutil.ReadAll(rw.Body)
		if err != nil {
			t.Fatal("could not read body")
		}

		return isAnyFourtunes(t, b)
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
*/

// ClockFunc 時間のモック用
type ClockFunc func() time.Time

// Now 時間のモック用
func (f ClockFunc) Now() time.Time {
	return f()
}

func TestFortune_NewYear(t *testing.T) {
	// テストケースが終わったらClockを元に戻す
	oc := fourtunes.Clock
	defer func() { fourtunes.Clock = oc }()

	tests := []struct {
		name string
		time time.Time
	}{
		{
			name: "new year start",
			time: time.Date(2018, 1, 01, 00, 0, 0, 0, time.Local),
		},
		{
			name: "new year end",
			time: time.Date(2018, 1, 03, 23, 59, 59, 0, time.Local),
		},
	}

	want := fourtunes.data[DAIKICHI]
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ここgiven when then以外の情報が長くなっててビミョ

			// HACK 共有変数を直に書き換えているのは壊れやすくて微妙
			fourtunes.Clock = ClockFunc(func() time.Time {
				return tt.time
			})

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)
			Fortune(w, r)

			rw := w.Result()
			defer rw.Body.Close()
			if rw.StatusCode != http.StatusOK {
				t.Fatal("unexpected status code")
			}
			b, err := ioutil.ReadAll(rw.Body)
			if err != nil {
				t.Fatal("could not read body")
			}

			got := marshal(t, b)
			if want.Luck != got.Luck || want.Message != got.Message {
				t.Fatalf("unexpected fourtune want = %v, got = %v", want, got)
			}
		})
	}
}

func marshal(t *testing.T, body []byte) *Fourtune {
	t.Helper()

	got := &Fourtune{}
	if err := json.Unmarshal(body, got); err != nil {
		fmt.Println(got)
		t.Fatalf("JSON Unmarshal error: %v", err)
	}

	return got
}

func isAnyFourtunes(t *testing.T, body []byte) bool {
	t.Helper()

	got := marshal(t, body)

	// HACK ここで直呼びしてるの良くはない
	for _, f := range fourtunes.data {
		// 本当はequals みたいなの用意した方が良さそう
		if f.Luck == got.Luck && f.Message == got.Message {
			return true
		}
	}

	return false
}
