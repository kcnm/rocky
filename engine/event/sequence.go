package event

import (
	"github.com/kcnm/rocky/engine"
)

type sequence struct {
	events []engine.Event
}

func Sequence(events ...engine.Event) engine.Event {
	return &sequence{events}
}

func (ev *sequence) Subject() interface{} {
	return ev.events
}

func (ev *sequence) Verb() engine.Verb {
	return engine.Sequence
}

func (ev *sequence) Object() interface{} {
	return nil
}

func (ev *sequence) Trigger(q engine.EventQueue) {
	for _, x := range ev.events {
		q.Cache(x, ev)
	}
}
