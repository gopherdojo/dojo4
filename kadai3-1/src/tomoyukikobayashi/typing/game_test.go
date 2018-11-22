package typing

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_NewGame(t *testing.T) {
	tests := []struct {
		name   string
		source io.Reader
		hasErr bool
	}{
		{
			name: "valid",
			source: strings.NewReader(`
Level1:
- hoge

Level2:
- difficult

Level3:
- test`),
			hasErr: false,
		},
		{
			name:   "invalid",
			source: strings.NewReader(``),
			hasErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewGame(1, tt.source, os.Stdin)
			if err != nil && !tt.hasErr {
				t.Fatalf("got err = %#v", err)
			}
		})
	}
}

func Test_Run(t *testing.T) {
	tests := []struct {
		name   string
		source io.Reader
		input  io.Reader
		hasErr bool
	}{
		{
			name: "valid",
			source: strings.NewReader(`
Level1:
- hoge

Level2:
- difficult

Level3:
- testtestsete
`),
			input: strings.NewReader(`hoge
hoge
hoge
hoge
hoge
hoge
hoge
hoge
miss
hoge
miss
hoge
difficult
difficult
`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game, err := NewGame(1, tt.source, tt.input)
			if err != nil {
				t.Fatalf("initialize game failed %v", game)
			}

			var want [2]int
			qCh, aCh, rCh := game.Run()
			for {
				q := <-qCh
				if q == "" {
					break
				}
				a := <-aCh
				if a == "" {
					break
				}

				// TOOD 値が微妙に揺れる。。。テスト壊れる
				if q == a {
					want[0] = want[0] + 1
				} else {
					want[1] = want[1] + 1
				}
			}

			got := <-rCh
			if want != got {
				t.Fatalf("failed got = %v, want = %v", got, want)
			}
		})
	}
}
