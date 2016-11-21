package event

import (
	"github.com/kcnm/rocky/engine"
)

type summon struct {
	game     engine.Game
	player   engine.Player
	card     engine.MinionCard
	position int
}

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

func (ev *summon) Trigger() {
	ev.game.Summon(ev.card, ev.player, ev.position)
}
