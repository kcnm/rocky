package event

import (
	"github.com/kcnm/rocky/engine"
)

type sequence struct {
	game   engine.Game
	events []engine.Event
}

func Sequence(game engine.Game, events ...engine.Event) engine.Event {
	return &sequence{game, events}
}

func (ev *sequence) Subject() interface{} {
	return ev.events
}

func (ev *sequence) Verb() engine.Verb {
	return engine.Sequence
}

func (ev *sequence) Trigger() {
	for _, x := range ev.events {
		ev.game.Cache(x, ev)
	}
}
