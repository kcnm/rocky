package event

import (
	"github.com/kcnm/rocky/engine"
)

type combined struct {
	events []engine.Event
}

func Combined(events ...engine.Event) engine.Event {
	return &combined{events}
}

func (ev *combined) Subject() interface{} {
	return ev.events
}

func (ev *combined) Verb() engine.Verb {
	return engine.Combined
}

func (ev *combined) Object() interface{} {
	return nil
}

func (ev *combined) Trigger(q engine.EventQueue) {
	for _, ev := range ev.events {
		ev.Trigger(q)
	}
}
