package event

import (
	"github.com/kcnm/rocky/engine"
)

type gameOver struct {
	game engine.Game
}

func GameOver(game engine.Game) engine.Event {
	return &gameOver{game}
}

func (ev *gameOver) Subject() interface{} {
	return ev.game
}

func (ev *gameOver) Verb() engine.Verb {
	return engine.GameOver
}

func (ev *gameOver) Trigger(q engine.EventQueue) {
}
