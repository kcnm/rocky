package event

import (
	"github.com/kcnm/rocky/engine"
)

type cast struct {
	game   engine.Game
	player engine.Player
	card   engine.SpellCard
	target engine.Char
}

func Cast(
	game engine.Game,
	player engine.Player,
	card engine.SpellCard,
	target engine.Char) engine.Event {
	return &cast{game, player, card, target}
}

func (ev *cast) Subject() interface{} {
	return ev.player
}

func (ev *cast) Verb() engine.Verb {
	return engine.Cast
}

func (ev *cast) Object() interface{} {
	return []interface{}{ev.card, ev.target}
}

func (ev *cast) Trigger(q engine.EventQueue) {
	q.Post(Impact(ev.game, ev.player, ev.target, ev.card.Effect()), ev)
}
