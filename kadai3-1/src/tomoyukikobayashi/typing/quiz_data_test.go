package typing

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_NewQuizData(t *testing.T) {
	tests := []struct {
		name   string
		reader io.Reader
		want   QuizSource
		hasErr bool
	}{
		{
			name: "valid",
			reader: strings.NewReader(`
Level1:
- hoge

Level2:
- difficult

Level3:
- test`),
			want: QuizSource{
				Level1: []string{"hoge"},
				Level2: []string{"difficult"},
				Level3: []string{"test"},
			},
			hasErr: false,
		},
		// TOOD 適当な文字列入れるとエラーなしでnil帰る
		{
			name: "invalidStruct",
			reader: strings.NewReader(`
LevelA:
- hoge
`),
			want:   QuizSource{},
			hasErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewQuizData(tt.reader)
			if err != nil && !tt.hasErr {
				t.Fatalf("got err = %#v", err)
			}
			if !reflect.DeepEqual(got, &tt.want) {
				t.Fatalf("want = %#v, got = %#v", tt.want, got)
			}
		})
	}
}

func Test_MaxLevel(t *testing.T) {
	q := QuizSource{
		Level1: []string{"hoge"},
		Level2: []string{"difficult"},
		Level3: []string{"test"},
	}

	// HACK 決め打ち
	want := 3
	got := q.MaxLevel()
	if want != got {
		t.Fatalf("want = %d, got = %d", want, got)
	}

}

func Test_WordsByLevel(t *testing.T) {
	q := QuizSource{
		Level1: []string{"hoge"},
		Level2: []string{"difficult"},
		Level3: []string{"test"},
	}

	tests := []struct {
		name  string
		level int
		want  []string
	}{
		{
			name:  "level1",
			level: 1,
			want:  []string{"hoge"},
		},
		{
			name:  "level2",
			level: 2,
			want:  []string{"difficult"},
		},
		{
			name:  "level3",
			level: 3,
			want:  []string{"test"},
		},
		{
			name:  "outOfRange",
			level: 4,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := q.WordsByLevel(tt.level)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("want = %#v, got = %#v", tt.want, got)
			}
		})
	}
}
