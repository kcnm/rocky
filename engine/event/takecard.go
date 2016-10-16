package event

import (
	"github.com/kcnm/rocky/engine/base"
)

type takeCard struct {
	player base.Player
	card   base.Card
}

func TakeCard(player base.Player, card base.Card) base.Event {
	return &takeCard{player, card}
}

func (ev *takeCard) Subject() interface{} {
	return ev.player
}

func (ev *takeCard) Verb() base.Verb {
	return base.TakeCard
}

func (ev *takeCard) Trigger() {
	ev.player.Take(ev.card)
}
