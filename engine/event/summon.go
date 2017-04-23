package event

import (
	"github.com/kcnm/rocky/engine"
)

// summon is an implementation of Summon Event.
type summon struct {
	game     engine.Game
	player   engine.Player
	card     engine.MinionCard
	position int
}

// Summon returns a new Event, where a minion is summoned for a player.
func Summon(
	game engine.Game,
	player engine.Player,
	card engine.MinionCard,
	position int) engine.Event {
	return &summon{game, player, card, position}
}

func (ev *summon) Subject() interface{} {
	return ev.player
}

func (ev *summon) Verb() engine.Verb {
	return engine.Summon
}

func (ev *summon) Object() interface{} {
	return ev.card
}

func (ev *summon) Trigger(q engine.EventQueue) {
	ev.game.Summon(ev.player, ev.card, ev.position)
}
