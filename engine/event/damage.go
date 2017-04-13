package event

import (
	"github.com/kcnm/rocky/engine"
)

type damage struct {
	src engine.Char
	dst engine.Char
	dmg int
}

func Damage(
	src, dst engine.Char,
	dmg int) engine.Event {
	return &damage{src, dst, dmg}
}

func (ev *damage) Subject() interface{} {
	return ev.src
}

func (ev *damage) Verb() engine.Verb {
	return engine.Damage
}

func (ev *damage) Object() interface{} {
	return ev.dst
}

func (ev *damage) Trigger(q engine.EventQueue) {
	actual, fatal := ev.dst.TakeDamage(ev.dmg)
	ev.dmg = actual
	if fatal {
		q.Post(Dying(ev.dst), ev)
	}
}
