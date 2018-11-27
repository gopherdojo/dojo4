package gameplayer

import (
	"context"
	"fmt"
	"io"

	"github.com/gopherdojo/dojo4/kadai3-1/iwata/questions"
)

type GamePlayer struct {
	display   io.Writer
	questions questions.List
}

func (p *GamePlayer) play(ctx context.Context) {
	fmt.Println("vim-go")
}
