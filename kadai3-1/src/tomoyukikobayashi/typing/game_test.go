package typing

import (
	"context"
	"io"
	"os"
	"strings"
	"testing"
	"time"
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
			_, err := NewGame(tt.source, os.Stdin)
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
			hasErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game, err := NewGame(tt.source, tt.input)
			if err != nil {
				t.Fatalf("initialize game failed %v", game)
			}

			tc := time.Duration(100) * time.Millisecond
			ctx, cancel := context.WithTimeout(context.Background(), tc)
			defer cancel()

			var want [2]int
			qCh, aCh, rCh := game.Run(ctx)
			var tmpQ, tmpA string
			for {
				select {
				case a := <-aCh:
					t.Errorf("a %v\n", a)
					tmpA = a
					// TODO 固まることがある
					// TOOD 値が微妙に揺れる。。。
					if tmpQ == tmpA {
						want[0] = want[0] + 1
					} else {
						want[1] = want[1] + 1
					}
				case q := <-qCh:
					t.Errorf("q %v\n", q)
					tmpQ = q
				case got := <-rCh:
					if want != got {
						t.Fatalf("failed got = %v, want = %v", got, want)
					}
				default:
					continue
				}
			}
		})
	}
}
