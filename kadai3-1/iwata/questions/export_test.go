package questions

func NewQuestions(qs ...*Question) Questions {
	return append(questions{}, qs...)
}
