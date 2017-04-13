package event

import (
	"github.com/kcnm/rocky/engine"
)

type destroy struct {
	char engine.Char
}

func Destroy(char engine.Char) engine.Event {
	return &destroy{char}
}

func (ev *destroy) Subject() interface{} {
	return ev.char
}

func (ev *destroy) Verb() engine.Verb {
	return engine.Destroy
}

func (ev *destroy) Trigger(q engine.EventQueue) {
}
