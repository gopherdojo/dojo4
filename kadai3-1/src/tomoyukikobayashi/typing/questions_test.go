package typing

import (
	"testing"
)

type testData struct {
}

func (td *testData) MaxLevel() int {
	return 3
}
func (td *testData) WordsByLevel(level int) []string {
	switch level {
	case 1:
		return []string{"hoge", "foo", "baz"}
	case 2:
		return []string{"difficult", "anymatch", "haeeee"}
	case 3:
		return []string{"test"}
	}
	return []string{}
}

func Test_GetNewWord(t *testing.T) {
	q := testData{}
	qst := NewQuestioner(&q)

	tests := []struct {
		name   string
		level  int
		wantIn []string
	}{
		{
			name:   "level1",
			level:  1,
			wantIn: q.WordsByLevel(1),
		},
		{
			name:   "level2",
			level:  2,
			wantIn: q.WordsByLevel(2),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := qst.GetNewWord(tt.level)
			for _, v := range tt.wantIn {
				if got == v {
					return
				}
			}
			t.Fatalf("want = %#v, got = %#v", tt.wantIn, got)
		})
	}
}
