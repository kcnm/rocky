package event

import (
	"github.com/kcnm/rocky/engine"
)

type damage struct {
	char engine.Char
	src  engine.Char
	dmg  int
}

func Damage(
	char engine.Char,
	src engine.Char,
	dmg int) engine.Event {
	return &damage{char, src, dmg}
}

func (ev *damage) Subject() interface{} {
	return ev.char
}

func (ev *damage) Verb() engine.Verb {
	return engine.Damage
}

func (ev *damage) Trigger(q engine.EventQueue) {
	actual, fatal := ev.char.TakeDamage(ev.dmg)
	ev.dmg = actual
	if fatal {
		q.Post(Dying(ev.char), ev)
	}
}
