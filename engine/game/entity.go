package game

import (
	"github.com/kcnm/rocky/engine"
)

type entity struct {
	id       engine.EntityID
	reactors []engine.Reactor
}

func newEntity(id engine.EntityID) engine.Entity {
	return &entity{id, make([]engine.Reactor, 0)}
}

func (e *entity) ID() engine.EntityID {
	return e.id
}

func (e *entity) React(ev engine.Event) {
	for _, r := range e.reactors {
		if ev.Verb() == engine.Combined {
			for _, ev := range ev.Subject().([]engine.Event) {
				r(ev)
			}
		} else {
			r(ev)
		}
	}
}

func (e *entity) AppendReactor(r engine.Reactor) {
	e.reactors = append(e.reactors, r)
}

func (e *entity) Reset() {
	e.reactors = make([]engine.Reactor, 0)
}
