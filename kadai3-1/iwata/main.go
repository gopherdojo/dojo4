package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gopherdojo/dojo4/kadai3-1/iwata/cmdparser"
	"github.com/gopherdojo/dojo4/kadai3-1/iwata/questions"
)

//nolint[:gocheknoinits]
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	p := cmdparser.New(os.Stderr)
	c, err := p.Parse(os.Args)
	if err != nil {
		log.Fatalf("Failed to parse command options: %+v", err)
	}

	qs, err := questions.Parse(c.TxtPath)
	if err != nil {
		log.Fatalf("Failed to parse text file: %+v", err)
	}

	q := qs.Give()
	fmt.Println(q.Match("test"))
}
