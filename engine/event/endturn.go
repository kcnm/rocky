package event

import (
	"github.com/kcnm/rocky/engine/base"
)

type endTurn struct {
	game base.Game
}

func EndTurn(game base.Game) base.Event {
	return &endTurn{game}
}

func (ev *endTurn) Subject() interface{} {
	return ev.game.CurrentPlayer()
}

func (ev *endTurn) Verb() base.Verb {
	return base.EndTurn
}

func (ev *endTurn) Trigger() {
	ev.game.Events().Post(base.StartTurn, ev)
}
