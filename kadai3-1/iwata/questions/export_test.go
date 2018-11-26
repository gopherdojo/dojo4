package questions

func NewList(qs ...*Item) List {
	return append(list{}, qs...)
}
