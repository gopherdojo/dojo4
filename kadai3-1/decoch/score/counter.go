package score

import "sync"

type Counter struct {
	value int
	m     sync.Mutex
}

func (s *Counter) Add(v int) {
	s.m.Lock()
	defer s.m.Unlock()
	s.value = s.value + v
}

func (s *Counter) Value() int {
	s.m.Lock()
	defer s.m.Unlock()
	return s.value
}
