package event

import (
	"github.com/kcnm/rocky/engine"
)

// impact is an implementation of Impact Event.
type impact struct {
	game   engine.Game
	you    engine.Player
	target engine.Char
	effect engine.Effect
}

// Impact returns a new Event, where an effect impacts.
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

func (ev *impact) Object() interface{} {
	return ev.target
}

func (ev *impact) Trigger(q engine.EventQueue) {
	ev.effect.Happen(ev.game, ev.you, ev.target, ev)
}
