package event

import (
	"github.com/kcnm/rocky/engine"
)

type heroPower struct {
	game   engine.Game
	player engine.Player
	tgt    engine.Char
}

func HeroPower(
	game engine.Game,
	player engine.Player,
	tgt engine.Char) engine.Event {
	return &heroPower{game, player, tgt}
}

func (ev *heroPower) Subject() interface{} {
	return ev.player
}

func (ev *heroPower) Verb() engine.Verb {
	return engine.HeroPower
}

func (ev *heroPower) Trigger() {
	for _, e := range ev.player.HeroPower() {
		e.Happen(ev.game, ev, []engine.Char{ev.tgt})
	}
}
