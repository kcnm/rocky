package event

import (
	"github.com/kcnm/rocky/engine"
)

type event struct {
	instance engine.Event
	cause    engine.Event
}

type queue struct {
	events    []*event
	cached    []*event
	entities  []engine.Entity
	entityIDs map[engine.EntityID]int
}

func NewQueue() engine.EventQueue {
	return &queue{
		make([]*event, 0),
		make([]*event, 0),
		make([]engine.Entity, 0),
		make(map[engine.EntityID]int),
	}
}

func (q *queue) Join(e engine.Entity) {
	q.entityIDs[e.ID()] = len(q.entities)
	q.entities = append(q.entities, e)
}

func (q *queue) Exit(e engine.Entity) bool {
	i, present := q.entityIDs[e.ID()]
	if !present {
		return false
	}
	q.entities = append(q.entities[:i], q.entities[i+1:]...)
	delete(q.entityIDs, e.ID())
	for id, j := range q.entityIDs {
		if j > i {
			q.entityIDs[id]--
		}
	}
	return true
}

func (q *queue) Fire(ev engine.Event) {
	if len(q.events) > 0 {
		panic("event queue is not empty")
	}
	q.Post(ev, nil)
	q.settle()
	for len(q.cached) > 0 {
		q.Post(q.cached[0].instance, q.cached[0].cause)
		q.cached = q.cached[1:]
		q.settle()
	}
}

func (q *queue) Post(ev engine.Event, cause engine.Event) {
	q.events = append(q.events, &event{ev, cause})
}

func (q *queue) Cache(ev engine.Event, cause engine.Event) {
	q.cached = append(q.cached, &event{ev, cause})
}

func (q *queue) Drain() {
	q.events = make([]*event, 0)
	q.cached = make([]*event, 0)
	q.entities = make([]engine.Entity, 0)
	q.entityIDs = make(map[engine.EntityID]int)
}

func (q *queue) settle() {
	for len(q.events) > 0 {
		ev := q.events[0].instance
		q.events = q.events[1:]
		mark := len(q.events)
		ev.Trigger(q)
		// Merges direct results for combined event.
		if ev.Verb() == engine.Combined && len(q.events) > mark {
			comb := &combined{make([]engine.Event, len(q.events)-mark)}
			for i := mark; i < len(q.events); i++ {
				comb.events[i-mark] = q.events[i].instance
			}
			q.events = append(q.events[:mark], &event{comb, ev})
		}
		// Notifies joined entities.
		for _, entity := range q.entities {
			entity.React(ev)
		}
		// Removes entities if destroyed.
		if ev.Verb() == engine.Combined {
			for _, ev := range ev.Subject().([]engine.Event) {
				q.maybeDestroyEntity(ev)
			}
		} else {
			q.maybeDestroyEntity(ev)
		}
	}
}

func (q *queue) maybeDestroyEntity(ev engine.Event) {
	if e, ok := ev.Subject().(engine.Entity); ok && ev.Verb() == engine.Destroy {
		q.Exit(e)
	}
}
