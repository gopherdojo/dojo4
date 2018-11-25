package timer

import (
	"time"
)

func NewCh(sec time.Duration) <-chan struct{} {
	timerCh := make(chan struct{})
	go func() {
		time.Sleep(sec * time.Second)
		timerCh <- struct{}{}
		close(timerCh)
	}()
	return timerCh
}
