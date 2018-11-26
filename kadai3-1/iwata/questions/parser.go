package questions

import (
	"bufio"
	"math/rand"
	"os"
	"time"

	"github.com/pkg/errors"
)

//nolint[:gocheknoinits]
func init() {
	rand.Seed(time.Now().UnixNano())
}

type Question struct {
	Quiz string
}

type Questions []Question

func Parse(txtFile string) (Questions, error) {
	questions := Questions{}

	f, err := os.Open(txtFile)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to open %s", txtFile)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		questions = append(questions, Question{Quiz: s.Text()})
	}

	return questions, nil
}

func (q Questions) Give() Question {
	return q[rand.Intn(len(q))]
}
