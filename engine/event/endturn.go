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

func (ev *endTurn) Object() interface{} {
	return ev.game
}

func (ev *endTurn) Trigger(q engine.EventQueue) {
	q.Post(StartTurn(ev.game), ev)
}
