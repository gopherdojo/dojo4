package typinggame

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"time"
)
type TypingGame int
//Start
//description:start 1 minute typing game.
//Param
//r:io.Reader example:os.Stdin
func (t *TypingGame)Start(r io.Reader){
	ch1 := input(r)
	ch2 := wait(60)
	var words = []string{"apple","banana","cherry","plum","grape","pineapple",}

	shuffle(words)
	i:=0
	//success count
	s_cnt:=0
	//fail count
	f_cnt :=0
	fmt.Println("try typing.1 minute.")
	fmt.Println(words[i])

TIMEOUT_LABEL:
	for {
		select {
		case msg:=<-ch1:
			if(words[i] == msg){
				if len(words) <= (i + 1){
					i = 0
				}else {
					i++
				}
				s_cnt++
			}else{
				fmt.Println("miss.retyping words.")
				f_cnt++
			}
			fmt.Println(words[i])
		case <-ch2:
			fmt.Println("")
			fmt.Println("time up.success count:",s_cnt," fail count:", f_cnt)

			break TIMEOUT_LABEL
		}
	}
}

func shuffle(a []string){
	for i:=len(a)-1;i >= 0;i-- {
		j:=rand.Intn(i+1)
		a[i],a[j] = a[j],a[i]
	}
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		defer close(ch)
		for s.Scan() {
			ch <- s.Text()
		}
	}()
	return ch
}

func wait(sec int) <-chan bool {
	ch := make(chan bool)
	go func() {
		time.Sleep(time.Duration(sec) * time.Second)
		ch <- true
	}()
	return ch
}
