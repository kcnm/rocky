package event

import (
	"github.com/kcnm/rocky/engine"
)

type gameOver struct {
}

func GameOver() engine.Event {
	return &gameOver{}
}

func (ev *gameOver) Subject() interface{} {
	return nil
}

func (ev *gameOver) Verb() engine.Verb {
	return engine.GameOver
}

func (ev *gameOver) Trigger() {
}
