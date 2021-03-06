package event

import (
	"github.com/kcnm/rocky/engine"
)

type takeCard struct {
	player engine.Player
	card   engine.Card
}

func TakeCard(player engine.Player, card engine.Card) engine.Event {
	return &takeCard{player, card}
}

func (ev *takeCard) Subject() interface{} {
	return ev.player
}

func (ev *takeCard) Verb() engine.Verb {
	return engine.TakeCard
}

func (ev *takeCard) Trigger() {
	ev.player.Take(ev.card)
}
