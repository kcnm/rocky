package event

import (
	"github.com/kcnm/rocky/engine"
)

type dying struct {
	char engine.Char
}

func Dying(char engine.Char) engine.Event {
	return &dying{char}
}

func (ev *dying) Subject() interface{} {
	return ev.char
}

func (ev *dying) Verb() engine.Verb {
	return engine.Dying
}

func (ev *dying) Trigger(q engine.EventQueue) {
	q.Post(Destroy(ev.char), ev)
}
