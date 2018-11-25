package pingpong

import "fmt"

type result struct {
	correct   int
	incorrect int
}

func (r *result) addCorrect() {
	r.correct++
}

func (r *result) addIncorrect() {
	r.incorrect++
}

func (r result) rate() float32 {
	if r.correct == 0 {
		return 0
	}
	sum := float32(r.correct) + float32(r.incorrect)
	return float32(r.correct) / sum * 100
}

func (r result) String() string {
	sum := r.correct + r.incorrect
	return fmt.Sprintf(`正解率: %.2f %%（%d/%d）`, r.rate(), r.correct, sum)
}
