package questions

import (
	"bufio"
	"os"

	"github.com/pkg/errors"
)

type Question struct {
	Quiz string
}

type Questions = []Question

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
