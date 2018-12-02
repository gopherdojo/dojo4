package omikuji

type Random interface {
	Intn() int
}

type RandomFunc func() int

func (f RandomFunc) Intn() int {
	return f()
}
