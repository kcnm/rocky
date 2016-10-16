package event

import (
	"github.com/kcnm/rocky/engine/base"
)

type playCard struct {
	player    base.Player
	cardIndex int
}

func PlayCard(player base.Player, cardIndex int) base.Event {
	return &playCard{player, cardIndex}
}

func (ev *playCard) Subject() interface{} {
	return ev.player
}

func (ev *playCard) Verb() base.Verb {
	return base.PlayCard
}

func (ev *playCard) Trigger() {
	ev.player.Play(ev.cardIndex)
}
