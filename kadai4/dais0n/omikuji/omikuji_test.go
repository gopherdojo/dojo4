package omikuji_test

import (
	"testing"
	"time"

	"github.com/dais0n/dojo4/kadai4/dais0n/omikuji"
)

func mockClock(t *testing.T, v string) omikuji.Clock {
	t.Helper()
	now, err := time.Parse("2006/01/02 15:04:05", v)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	return omikuji.ClockFunc(func() time.Time {
		return now
	})
}

func mockRandom(t *testing.T, v int) omikuji.Random {
	t.Helper()
	return omikuji.RandomFunc(func() int {
		return v
	})
}

func TestOmikuji_Do(t *testing.T) {
	cases := map[string]struct {
		clock       omikuji.Clock
		random      omikuji.Random
		result      string
		expectError bool
	}{
		"luckyDay_1_1": {
			clock:       mockClock(t, "2019/01/01 0:00:00"),
			random:      nil,
			result:      "大吉",
			expectError: false,
		},
		"luckyDay_1_3": {
			clock:       mockClock(t, "2019/01/03 23:59:59"),
			random:      nil,
			result:      "大吉",
			expectError: false,
		},
		"normalDay": {
			clock:       nil,
			random:      mockRandom(t, 2),
			result:      "中吉",
			expectError: false,
		},
		"unluckyDay": {
			clock:       nil,
			random:      mockRandom(t, 4),
			result:      "凶",
			expectError: false,
		},
		"unexpectedRandomNum": {
			clock:       nil,
			random:      mockRandom(t, 9999),
			result:      "",
			expectError: true,
		},
	}
	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			omikuji := omikuji.Omikuji{
				Clock:  tc.clock,
				Random: tc.random,
			}
			result, err := omikuji.Do()
			if err == nil && tc.expectError {
				t.Error("expected error did not occur")
			}
			if err != nil && !tc.expectError {
				t.Error("expected error did not occur")
			}
			if err != nil && result != tc.result {
				t.Errorf("omikuji do result get %s want %s", tc.result, result)
			}
		})
	}
}
