package event

import (
	"github.com/kcnm/rocky/engine"
)

type damage struct {
	game engine.Game
	char engine.Char
	src  engine.Char
	dmg  int
}

func Damage(
	game engine.Game,
	char engine.Char,
	src engine.Char,
	dmg int) engine.Event {
	return &damage{game, char, src, dmg}
}

func (ev *damage) Subject() interface{} {
	return ev.char
}

func (ev *damage) Verb() engine.Verb {
	return engine.Damage
}

func (ev *damage) Trigger() {
	actual, fatal := ev.char.TakeDamage(ev.dmg)
	ev.dmg = actual
	if fatal {
		ev.game.Events().Post(
			Dying(ev.game, ev.char), ev)
	}
}
