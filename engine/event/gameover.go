package event

import (
	"github.com/kcnm/rocky/engine"
)

// gameOver is an implementation of GameOver Event.
type gameOver struct {
	game engine.Game
}

// GameOver returns a new Event, where the game is over.
func GameOver(game engine.Game) engine.Event {
	return &gameOver{game}
}

func (ev *gameOver) Subject() interface{} {
	return ev.game
}

func (ev *gameOver) Verb() engine.Verb {
	return engine.GameOver
}

func (ev *gameOver) Object() interface{} {
	return nil
}

func (ev *gameOver) Trigger(q engine.EventQueue) {
}
