package controller_test

import (
	"errors"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/dais0n/dojo4/kadai4/dais0n/controller"
)

type OmikujiMock struct {
	MockDo func() (string, error)
}

func (o *OmikujiMock) Do() (string, error) {
	return o.MockDo()
}

func TestOmikujiController_Get(t *testing.T) {
	cases := map[string]struct {
		omikuji controller.Omikuji
		output  string
		status  int
	}{
		"normal": {
			omikuji: &OmikujiMock{
				MockDo: func() (string, error) {
					return "中吉", nil
				},
			},
			output: "{\"result\":\"中吉\"}\n",
			status: 200,
		},
		"error": {
			omikuji: &OmikujiMock{
				MockDo: func() (string, error) {
					return "", errors.New("unexpected error")
				},
			},
			output: "{\"error\":\"unexpected error\"}\n",
			status: 400,
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/v1/omikuji", nil)
			c := controller.NewOmikujiController(tc.omikuji)
			c.Get(w, r)
			rw := w.Result()
			defer rw.Body.Close()
			if rw.StatusCode != tc.status {
				t.Fatalf("unexpected status code want %d, get %d", tc.status, rw.StatusCode)
			}
			body, err := ioutil.ReadAll(rw.Body)
			if err != nil {
				t.Fatal("body read error")
			}
			if string(body) != tc.output {
				t.Errorf("unexpected output want %s, get %s", tc.output, string(body))
			}
		})
	}
}
