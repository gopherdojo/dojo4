package gameplayer_test

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/gopherdojo/dojo4/kadai3-1/iwata/gameplayer"
	"github.com/gopherdojo/dojo4/kadai3-1/iwata/questions"
)

type MockQuestionList struct {
	q *questions.Item
}

func (m *MockQuestionList) Give() *questions.Item {
	return m.q
}

func TestGamePlayer_Play(t *testing.T) {
	type fields struct {
		r io.Reader
		q *questions.Item
	}
	tests := []struct {
		name    string
		fields  fields
		correct bool
		wantErr bool
	}{
		{"correct", fields{strings.NewReader("CORRECT"), &questions.Item{Quiz: "CORRECT"}}, true, false},
		{"in correct", fields{strings.NewReader("IN CORRECT"), &questions.Item{Quiz: "CORRECT"}}, false, false},
		{"over max size", fields{bytes.NewReader(make([]byte, bufio.MaxScanTokenSize+1)), &questions.Item{}}, false, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			w := &bytes.Buffer{}
			p := gameplayer.NewGame(w, tt.fields.r, &MockQuestionList{tt.fields.q})
			ctx := context.Background()
			got, err := p.Play(ctx, time.Millisecond*100)
			if (err != nil) != tt.wantErr {
				t.Errorf("GamePlayer.Play() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if tt.wantErr {
				return
			}
			var want *gameplayer.Score
			if tt.correct {
				want = &gameplayer.Score{CorrectNum: 1, InCorrectNum: 0}
			} else {
				want = &gameplayer.Score{CorrectNum: 0, InCorrectNum: 1}
			}
			if diff := cmp.Diff(got, want); diff != "" {
				t.Errorf("GamePlayer.Play() = %v, want %v, differs: (-got +want):\n%s", got, want, diff)
			}
		})
	}

	t.Run("timeout game before any scan", func(t *testing.T) {
		w := &bytes.Buffer{}
		p := gameplayer.NewGame(w, strings.NewReader("cancel"), &MockQuestionList{&questions.Item{}})
		ctx := context.Background()
		got, err := p.Play(ctx, 1)
		if err != nil {
			t.Errorf("GamePlayer.Play() error = %v, wantErr false", err)
			return
		}
		want := &gameplayer.Score{}
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("GamePlayer.Play() = %v, want %v, differs: (-got +want):\n%s", got, want, diff)
		}
	})
}
