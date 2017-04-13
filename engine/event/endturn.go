package event

import (
	"github.com/kcnm/rocky/engine"
)

type endTurn struct {
	game engine.Game
}

func EndTurn(game engine.Game) engine.Event {
	return &endTurn{game}
}

func (ev *endTurn) Subject() interface{} {
	return ev.game.CurrentPlayer()
}

func (ev *endTurn) Verb() engine.Verb {
	return engine.EndTurn
}

func (ev *endTurn) Trigger() {
	ev.game.Post(StartTurn(), ev)
}
