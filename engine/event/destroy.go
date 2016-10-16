package event

import (
	"github.com/kcnm/rocky/engine/base"
)

type destroy struct {
	game base.Game
	char base.Character
}

func Destroy(game base.Game, char base.Character) base.Event {
	return &destroy{game, char}
}

func (ev *destroy) Subject() interface{} {
	return ev.char
}

func (ev *destroy) Verb() base.Verb {
	return base.Destroy
}

func (ev *destroy) Trigger() {
}
