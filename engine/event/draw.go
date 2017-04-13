package event

import (
	"github.com/kcnm/rocky/engine"
)

type draw struct {
	player engine.Player
}

func Draw(player engine.Player) engine.Event {
	return &draw{player}
}

func (ev *draw) Subject() interface{} {
	return ev.player
}

func (ev *draw) Verb() engine.Verb {
	return engine.Draw
}

func (ev *draw) Object() interface{} {
	return ev.player.Deck()
}

func (ev *draw) Trigger(q engine.EventQueue) {
	card, fatigue := ev.player.Deck().Draw()
	if fatigue > 0 {
		q.Post(Damage(nil, ev.player, fatigue), ev)
		return
	}
	if ev.player.HandIsFull() {
		return
	}
	q.Post(TakeCard(ev.player, card), ev)
}
