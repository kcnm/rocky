package event

import (
	"github.com/kcnm/rocky/engine"
)

type destroy struct {
	game engine.Game
	char engine.Character
}

func Destroy(game engine.Game, char engine.Character) engine.Event {
	return &destroy{game, char}
}

func (ev *destroy) Subject() interface{} {
	return ev.char
}

func (ev *destroy) Verb() engine.Verb {
	return engine.Destroy
}

func (ev *destroy) Trigger() {
}
