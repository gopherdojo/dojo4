package fortune_test

import (
	"bytes"
	"net/http"
	"testing"
	"time"

	"github.com/gopherdojo/dojo4/kadai4/uobikiemukot/fortune"
)

// dummyWriter implement http.ResponseWriter
type dummyWriter struct {
	buf bytes.Buffer
}

func (d *dummyWriter) Header() http.Header {
	return nil
}

func (d *dummyWriter) Write(b []byte) (int, error) {
	return d.buf.Write(b)
}

func (d *dummyWriter) WriteHeader(code int) {
}

// mockWriter creates dummyWriter from bytes.Buffer
func mockWriter(t *testing.T, b bytes.Buffer) dummyWriter {
	t.Helper()
	return dummyWriter{buf: b}
}

// mockClock creates fortune.ClockFunc from date string
func mockClock(t *testing.T, s string) fortune.ClockFunc {
	t.Helper()

	now, err := time.Parse("2006/01/02", s)
	if err != nil {
		t.Fatal(err)
	}

	return fortune.ClockFunc(func() time.Time {
		return now
	})
}

func parse(t *testing.T, s string) string {
	t.Helper()

	r, err := fortune.Decode(s)
	if err != nil {
		t.Fatal(err)
	}

	return r.Fortune
}

func TestHandler_Success(t *testing.T) {
	w := mockWriter(t, bytes.Buffer{})

	cases := map[string]struct {
		writer *dummyWriter
		clock  fortune.Clock
	}{
		"1900/01/01": {
			writer: &w,
			clock:  mockClock(t, "1900/01/01"),
		},
		"2000/02/29": {
			writer: &w,
			clock:  mockClock(t, "2000/02/29"),
		},
		"2020/12/19": {
			writer: &w,
			clock:  mockClock(t, "2020/12/19"),
		},
	}

	for _, c := range cases {
		c.writer.buf.Reset()

		f := &fortune.Fortune{Clock: c.clock}
		f.Handler(c.writer, nil)

		if f.Err != nil {
			t.Errorf("error occurred: %s", f.Err)
		}
	}
}

func TestHandler_SpecialDate(t *testing.T) {
	w := mockWriter(t, bytes.Buffer{})

	cases := map[string]struct {
		writer *dummyWriter
		clock  fortune.Clock
		msg    string
	}{
		"2019/01/01": {
			writer: &w,
			clock:  mockClock(t, "2019/01/01"),
			msg:    "大吉",
		},
		"2020/01/02": {
			writer: &w,
			clock:  mockClock(t, "2020/01/02"),
			msg:    "大吉",
		},
		"2021/01/03": {
			writer: &w,
			clock:  mockClock(t, "2021/01/03"),
			msg:    "大吉",
		},
	}

	for _, c := range cases {
		c.writer.buf.Reset()

		f := &fortune.Fortune{Clock: c.clock}
		f.Handler(c.writer, nil)

		r := parse(t, c.writer.buf.String())
		if r != c.msg {
			t.Errorf("want:%s but got:%s\n", c.msg, r)
		}
	}
}
