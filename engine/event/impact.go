package event

import (
	"github.com/kcnm/rocky/engine"
)

type impact struct {
	game   engine.Game
	you    engine.Player
	target engine.Char
	effect engine.Effect
}

func Impact(
	game engine.Game,
	you engine.Player,
	target engine.Char,
	effect engine.Effect) engine.Event {
	return &impact{game, you, target, effect}
}

func (ev *impact) Subject() interface{} {
	return ev.effect
}

func (ev *impact) Verb() engine.Verb {
	return engine.Impact
}

func (ev *impact) Trigger(q engine.EventQueue) {
	ev.effect.Happen(ev.game, ev.you, ev.target, ev)
}
