package event

import (
	"github.com/kcnm/rocky/engine"
)

type destroy struct {
	entity engine.Entity
}

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
