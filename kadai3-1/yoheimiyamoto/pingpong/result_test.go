package pingpong

import (
	"fmt"
	"testing"
)

func TestResult(t *testing.T) {
	var r result
	r.addCorrect()
	r.addCorrect()
	r.addIncorrect()
	fmt.Println(r)
}
