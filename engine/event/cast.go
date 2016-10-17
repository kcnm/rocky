package event

import (
	"github.com/kcnm/rocky/engine/base"
)

type cast struct {
	game   base.Game
	player base.Player
	card   base.SpellCard
	tgt    base.Character
}

func Cast(
	game base.Game,
	player base.Player,
	card base.SpellCard,
	tgt base.Character) base.Event {
	return &cast{game, player, card, tgt}
}

func (ev *cast) Subject() interface{} {
	return ev.player
}

func (ev *cast) Verb() base.Verb {
	return base.Cast
}

func (ev *cast) Trigger() {
	for _, effect := range ev.card.Effects() {
		effect.Happen(ev.game, ev, []base.Character{ev.tgt})
	}
}
