package omikuji

import (
	"testing"
	"time"
)

func TestGetType(t *testing.T) {
	t.Run("", func(t *testing.T) {
		testGetType(t, time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local), "大吉")
	})
}

// TestGetTypeのテストヘルパー
func testGetType(t *testing.T, in time.Time, expected string) {
	t.Helper()
	r := getType(in)
	if r != expected {
		t.Errorf("expected: %s, actual: %s", expected, r)
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
