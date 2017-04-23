package event

import (
	"github.com/kcnm/rocky/engine"
)

// dying is an implementation of Dying Event.
type dying struct {
	entity engine.Entity
}

// Dying returns a new Event, where an entity is about to be destroyed.
func Dying(entity engine.Entity) engine.Event {
	return &dying{entity}
}

func (ev *dying) Subject() interface{} {
	return ev.entity
}

func (ev *dying) Verb() engine.Verb {
	return engine.Dying
}

func (ev *dying) Object() interface{} {
	return nil
}

func (ev *dying) Trigger(q engine.EventQueue) {
	q.Post(Destroy(ev.entity), ev)
}
