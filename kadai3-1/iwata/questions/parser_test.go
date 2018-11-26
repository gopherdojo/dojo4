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
		{"empty file", args{"../testdata/parse_empty.txt"}, questions.Questions{}, false},
		{"parse successfully", args{"../testdata/parse_questions.txt"}, questions.Questions{
			questions.Question{"hoge"},
			questions.Question{"fuga"},
			questions.Question{"Go Lang"},
			questions.Question{"I have a pen"},
			questions.Question{"I have an apple"},
		}, false},
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
	q := questions.Question{
		Quiz: "test",
	}
	questions := questions.Questions{q}
	got := questions.Give()
	if diff := cmp.Diff(got, q); diff != "" {
		t.Errorf("Give() = %v, want %v, differs: (-got +want):\n%s", got, q, diff)
	}
}
