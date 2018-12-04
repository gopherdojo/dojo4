package omikuji

import (
	"testing"
	"time"
)

func TestGetType(t *testing.T) {
	now := time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)
	r := getType(now)
	const expected = "大吉"
	if r != expected {
		t.Fatalf("expected: %s, actual: %s", expected, r)
	}
}
