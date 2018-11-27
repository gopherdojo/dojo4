package gameplayer_test

import (
	"bytes"
	"context"
	"io"
	"reflect"
	"strings"
	"testing"

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
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			w := &bytes.Buffer{}
			p := gameplayer.NewGame(w, tt.fields.r, &MockQuestionList{tt.fields.q})
			ctx := context.Background()
			got, err := p.Play(ctx, 1)
			if (err != nil) != tt.wantErr {
				t.Errorf("GamePlayer.Play() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			var want *gameplayer.Score
			if tt.correct {
				want = &gameplayer.Score{CorrectNum: 1, InCorrectNum: 0}
			} else {
				want = &gameplayer.Score{CorrectNum: 0, InCorrectNum: 1}
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("GamePlayer.Play() = %v, want %v", got, want)
			}
		})
	}
}
