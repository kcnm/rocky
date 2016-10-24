package event

import (
	"github.com/kcnm/rocky/engine"
)

type playCard struct {
	player    engine.Player
	cardIndex int
}

func PlayCard(player engine.Player, cardIndex int) engine.Event {
	return &playCard{player, cardIndex}
}

func (ev *playCard) Subject() interface{} {
	return ev.player
}

func (ev *playCard) Verb() engine.Verb {
	return engine.PlayCard
}

func (ev *playCard) Trigger() {
	ev.player.Play(ev.cardIndex)
}
