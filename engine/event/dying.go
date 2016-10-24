package event

import (
	"github.com/kcnm/rocky/engine"
)

type dying struct {
	game engine.Game
	char engine.Character
}

func Dying(game engine.Game, char engine.Character) engine.Event {
	return &dying{game, char}
}

func (ev *dying) Subject() interface{} {
	return ev.char
}

func (ev *dying) Verb() engine.Verb {
	return engine.Dying
}

func (ev *dying) Trigger() {
	ev.game.Events().Post(
		Destroy(ev.game, ev.char), ev)
}
