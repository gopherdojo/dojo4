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

func TestPlay(t *testing.T) {
	c := &client{
		now(func() time.Time { return time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local) }),
	}
	actual := *c.play()
	expected := Result{"大吉"}
	if actual != expected {
		t.Fatalf("actual: %v, expected: %v", actual, expected)
	}
}
