package event

import (
	"github.com/kcnm/rocky/engine/base"
)

type draw struct {
	game   base.Game
	player base.Player
}

func Draw(game base.Game, player base.Player) base.Event {
	return &draw{game, player}
}

func (ev *draw) Subject() interface{} {
	return ev.player
}

func (ev *draw) Verb() base.Verb {
	return base.Draw
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
