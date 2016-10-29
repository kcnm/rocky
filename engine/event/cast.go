package event

import (
	"github.com/kcnm/rocky/engine"
)

type cast struct {
	game   engine.Game
	player engine.Player
	card   engine.SpellCard
	tgt    engine.Char
}

func Cast(
	game engine.Game,
	player engine.Player,
	card engine.SpellCard,
	tgt engine.Char) engine.Event {
	return &cast{game, player, card, tgt}
}

func (ev *cast) Subject() interface{} {
	return ev.player
}

func (ev *cast) Verb() engine.Verb {
	return engine.Cast
}

func (ev *cast) Trigger() {
	ev.game.Events().Post(
		Impact(ev.game, ev.tgt, ev.card.Effect()), ev)
}
