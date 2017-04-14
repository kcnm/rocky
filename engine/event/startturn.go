package event

import (
	"github.com/kcnm/rocky/engine"
)

type startTurn struct {
	game   engine.Game
	player engine.Player
}

func StartTurn(game engine.Game, player engine.Player) engine.Event {
	return &startTurn{game, player}
}

func (ev *startTurn) Subject() interface{} {
	return ev.player
}

func (ev *startTurn) Verb() engine.Verb {
	return engine.StartTurn
}

func (ev *startTurn) Object() interface{} {
	return ev.game
}

func (ev *startTurn) Trigger(q engine.EventQueue) {
}
