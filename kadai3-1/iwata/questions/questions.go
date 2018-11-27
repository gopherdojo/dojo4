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

type Item struct {
	Quiz string
}

type List interface {
	Give() *Item
}

type list []*Item

func Parse(txtFile string) (List, error) {
	l := list{}

	f, err := os.Open(txtFile)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to open %s", txtFile)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		if s.Text() == "" {
			continue
		}
		l = append(l, &Item{Quiz: s.Text()})
	}

	return l, nil
}

func (l list) Give() *Item {
	return l[rand.Intn(len(l))]
}

func (q *Item) Match(answer string) bool {
	return q.Quiz == answer
}
