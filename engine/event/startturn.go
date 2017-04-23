package event

import (
	"github.com/kcnm/rocky/engine"
)

// startTurn is an implementation of StartTurn Event.
type startTurn struct {
	game   engine.Game
	player engine.Player
}

// StartTurn returns a new Event, where a player starts his turn.
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
