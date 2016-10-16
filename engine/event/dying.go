package event

import (
	"github.com/kcnm/rocky/engine/base"
)

type dying struct {
	game base.Game
	char base.Character
}

func Dying(game base.Game, char base.Character) base.Event {
	return &dying{game, char}
}

func (ev *dying) Subject() interface{} {
	return ev.char
}

func (ev *dying) Verb() base.Verb {
	return base.Dying
}

func (ev *dying) Trigger() {
	ev.game.Events().Post(
		Destroy(ev.game, ev.char))
}
