package event

import (
	"github.com/kcnm/rocky/engine/base"
)

type damage struct {
	game base.Game
	char base.Character
	src  base.Character
	dmg  int
}

func Damage(
	game base.Game,
	char base.Character,
	src base.Character,
	dmg int) base.Event {
	return &damage{game, char, src, dmg}
}

func (ev *damage) Subject() interface{} {
	return ev.char
}

func (ev *damage) Verb() base.Verb {
	return base.Damage
}

func (ev *damage) Trigger() {
	actual, fatal := ev.char.TakeDamage(ev.dmg)
	ev.dmg = actual
	if fatal {
		ev.game.Events().Post(
			Dying(ev.game, ev.char), ev)
	}
}
