package event

import (
	"github.com/kcnm/rocky/engine"
)

// takeCard is an implementation of TakeCard Event.
type takeCard struct {
	player engine.Player
	card   engine.Card
}

// TakeCard returns a new Event, where a player takes a card into his hand.
func TakeCard(player engine.Player, card engine.Card) engine.Event {
	return &takeCard{player, card}
}

func (ev *takeCard) Subject() interface{} {
	return ev.player
}

func (ev *takeCard) Verb() engine.Verb {
	return engine.TakeCard
}

func (ev *takeCard) Object() interface{} {
	return ev.card
}

func (ev *takeCard) Trigger(q engine.EventQueue) {
	ev.player.Take(ev.card)
}
