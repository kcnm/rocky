package event

import (
	"github.com/kcnm/rocky/engine"
)

type startTurn struct {
	game engine.Game
}

func StartTurn(game engine.Game) engine.Event {
	return &startTurn{game}
}

func (ev *startTurn) Subject() interface{} {
	return ev.game.CurrentPlayer()
}

func (ev *startTurn) Verb() engine.Verb {
	return engine.StartTurn
}

func (ev *startTurn) Object() interface{} {
	return ev.game
}

func (ev *startTurn) Trigger(q engine.EventQueue) {
}
