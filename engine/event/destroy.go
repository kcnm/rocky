package event

import (
	"github.com/kcnm/rocky/engine"
)

// destroy is an implementation of Destroy Event.
type destroy struct {
	entity engine.Entity
}

// Destroy returns a new Event, where an entity is destroyed.
func Destroy(entity engine.Entity) engine.Event {
	return &destroy{entity}
}

func (ev *destroy) Subject() interface{} {
	return ev.entity
}

func (ev *destroy) Verb() engine.Verb {
	return engine.Destroy
}

func (ev *destroy) Object() interface{} {
	return nil
}

func (ev *destroy) Trigger(q engine.EventQueue) {
}
