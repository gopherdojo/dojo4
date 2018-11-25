package pingpong

import "fmt"

// 正解、不正解のスコアをカウント
type score struct {
	correct   int
	incorrect int
}

// 正解をカウントアップ
func (s *score) addCorrect() {
	s.correct++
}

// 不正解をカウントアップ
func (s *score) addIncorrect() {
	s.incorrect++
}

func (s *score) count() int {
	return s.correct + s.incorrect
}

// 正解率の算出
func (s score) rate() float32 {
	if s.correct == 0 {
		return 0
	}
	sum := float32(s.correct) + float32(s.incorrect)
	return float32(s.correct) / sum * 100
}

// 結果を出力
func (s score) Result() string {
	sum := s.correct + s.incorrect
	return fmt.Sprintf(`正解率: %.2f %%（%d/%d）`, s.rate(), s.correct, sum)
}
