package fortune_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gopherdojo/dojo4/kadai4/akuchii/fortune"
)

func TestFortune_Handler(t *testing.T) {
	f := fortune.NewFortune(fortune.DefaultClock{})
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	f.Handler(w, r)
	rw := w.Result()
	defer rw.Body.Close()
	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}
	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}

	fr := &fortune.LotResult{}
	if err := json.Unmarshal(b, &fr); err != nil {
		t.Fatal("failed json unmarshal")
	}

	var fortuneList = f.GetFortuneList()
	contain := false
	for _, v := range fortuneList {
		if v == fr.Result {
			contain = true
			break
		}
	}

	if !contain {
		t.Fatalf("unexpected response: %s", string(b))
	}
}

type MockClock struct {
	currentTime time.Time
}

func (mc MockClock) GetCurrentTime() time.Time {
	return mc.currentTime
}

func TestFortune_LotOnNewYearDay(t *testing.T) {
	cases := []struct {
		newYearDay time.Time
	}{
		{time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)},
		{time.Date(2018, 1, 2, 0, 0, 0, 0, time.UTC)},
		{time.Date(2017, 1, 3, 23, 59, 59, 0, time.UTC)},
	}
	for _, c := range cases {
		mc := &MockClock{c.newYearDay}
		f := fortune.NewFortune(mc)
		result := f.Lot()
		if result != "大吉" {
			t.Errorf("unexpected result: %s on %v", result, mc.GetCurrentTime())
		}
	}
}
