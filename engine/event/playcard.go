package event

import (
	"github.com/kcnm/rocky/engine"
)

// playCard is an implementation of PlayCard Event.
type playCard struct {
	player    engine.Player
	cardIndex int
}

// PlayCard returns a new Event, where a player plays one card from his hand.
func PlayCard(player engine.Player, cardIndex int) engine.Event {
	return &playCard{player, cardIndex}
}

func (ev *playCard) Subject() interface{} {
	return ev.player
}

func (ev *playCard) Verb() engine.Verb {
	return engine.PlayCard
}

func (ev *playCard) Object() interface{} {
	return ev.player.Hand()[ev.cardIndex]
}

func (ev *playCard) Trigger(q engine.EventQueue) {
	ev.player.Play(ev.cardIndex)
}
