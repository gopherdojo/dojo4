package questions

import (
	"bufio"
	"math/rand"
	"os"

	"github.com/pkg/errors"
)

type Question struct {
	Quiz string
}

type Questions interface {
	Give() *Question
}

type questions []*Question

func Parse(txtFile string) (Questions, error) {
	questions := questions{}

	f, err := os.Open(txtFile)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to open %s", txtFile)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		questions = append(questions, &Question{Quiz: s.Text()})
	}

	return questions, nil
}

func (qs questions) Give() *Question {
	return qs[rand.Intn(len(qs))]
}

func (q *Question) Match(answer string) bool {
	return q.Quiz == answer
}
