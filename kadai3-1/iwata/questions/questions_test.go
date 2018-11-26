package questions_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/gopherdojo/dojo4/kadai3-1/iwata/questions"
)

func TestParse(t *testing.T) {
	type args struct {
		txtFile string
	}
	tests := []struct {
		name    string
		args    args
		want    questions.Questions
		wantErr bool
	}{
		{"file not exist", args{"./not-exist.txt"}, nil, true},
		{"empty file", args{"../testdata/parse_empty.txt"}, questions.NewQuestions(), false},
		{"parse successfully", args{"../testdata/parse_questions.txt"}, questions.NewQuestions(
			&questions.Question{"hoge"},
			&questions.Question{"fuga"},
			&questions.Question{"Go Lang"},
			&questions.Question{"I have a pen"},
			&questions.Question{"I have an apple"},
		), false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := questions.Parse(tt.args.txtFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Parse() = %v, want %v, differs: (-got +want):\n%s", got, tt.want, diff)
			}
		})
	}
}

func TestQuestions_Give(t *testing.T) {
	q := &questions.Question{
		Quiz: "test",
	}
	questions := questions.NewQuestions(q)
	got := questions.Give()
	if diff := cmp.Diff(got, q); diff != "" {
		t.Errorf("Give() = %v, want %v, differs: (-got +want):\n%s", got, q, diff)
	}
}

func TestQuestion_Match(t *testing.T) {
	type fields struct {
		Quiz string
	}
	type args struct {
		answer string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"matched", fields{"test"}, args{"test"}, true},
		{"multi words", fields{"test hoge fuga"}, args{"test hoge fuga"}, true},
		{"not matched", fields{"test"}, args{"test2"}, false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			q := questions.Question{
				Quiz: tt.fields.Quiz,
			}
			if got := q.Match(tt.args.answer); got != tt.want {
				t.Errorf("Question.Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
