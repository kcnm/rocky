package event

import (
	"github.com/kcnm/rocky/engine"
)

type endTurn struct {
	game   engine.Game
	player engine.Player
}

func EndTurn(game engine.Game, player engine.Player) engine.Event {
	return &endTurn{game, player}
}

func (ev *endTurn) Subject() interface{} {
	return ev.player
}

func (ev *endTurn) Verb() engine.Verb {
	return engine.EndTurn
}

func (ev *endTurn) Object() interface{} {
	return ev.game
}

func (ev *endTurn) Trigger(q engine.EventQueue) {
	next := ev.game.Opponent(ev.player)
	q.Post(StartTurn(ev.game, next), ev)
}
