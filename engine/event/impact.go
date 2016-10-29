package event

import (
	"github.com/kcnm/rocky/engine"
)

type impact struct {
	game   engine.Game
	target engine.Char
	effect engine.Effect
}

func Impact(
	game engine.Game,
	target engine.Char,
	effect engine.Effect) engine.Event {
	return &impact{game, target, effect}
}

func (ev *impact) Subject() interface{} {
	return ev.effect
}

func (ev *impact) Verb() engine.Verb {
	return engine.Impact
}

func (ev *impact) Trigger() {
	ev.effect.Happen(ev.game, ev, ev.target)
}
