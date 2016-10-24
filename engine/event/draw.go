package event

import (
	"github.com/kcnm/rocky/engine"
)

type draw struct {
	game   engine.Game
	player engine.Player
}

func Draw(game engine.Game, player engine.Player) engine.Event {
	return &draw{game, player}
}

func (ev *draw) Subject() interface{} {
	return ev.player
}

func (ev *draw) Verb() engine.Verb {
	return engine.Draw
}

func (ev *draw) Trigger() {
	card, fatigue := ev.player.Deck().Draw()
	if fatigue > 0 {
		ev.game.Events().Post(
			Damage(ev.game, ev.player, nil, fatigue), ev)
		return
	}
	if ev.player.HandIsFull() {
		return
	}
	ev.game.Events().Post(
		TakeCard(ev.player, card), ev)
}
