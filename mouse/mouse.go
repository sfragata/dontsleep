package mouse

import (
	"log"

	"github.com/go-vgo/robotgo"
)

type Options struct {
	Verbose    bool
	MoveOffset int
}

func (op Options) Move() {
	if op.Verbose {
		log.Println("Moving Mouse...")
	}
	x, y := robotgo.GetMousePos()

	if op.Verbose {
		log.Printf("Position: x=%d, y=%d", x, y)
	}
	robotgo.MoveMouseSmooth(x+op.MoveOffset, y+op.MoveOffset, 1.0, 7.0)
	robotgo.MoveMouseSmooth(x-op.MoveOffset, y-op.MoveOffset, 1.0, 7.0)
}
