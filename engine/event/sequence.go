package event

import (
	"github.com/kcnm/rocky/engine"
)

// sequence is an implementation of Sequence Event.
type sequence struct {
	events []engine.Event
}

// Sequence returns a new Event, which is a sequence of several child events
// logically happening in a sequential order.
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
