package event

import (
	"github.com/kcnm/rocky/engine"
)

type event struct {
	instance engine.Event
	cause    engine.Event
}

type queue struct {
	events      []*event
	cached      []*event
	listeners   []engine.Listener
	listenerIDs map[engine.ListenerID]int
	idGen       engine.ListenerID
}

func NewQueue() engine.EventQueue {
	return &queue{
		make([]*event, 0),
		make([]*event, 0),
		make([]engine.Listener, 0),
		make(map[engine.ListenerID]int),
		0, // idGen
	}
}

func (q *queue) AddListener(listener engine.Listener) engine.ListenerID {
	id := q.nextListenerID()
	q.listenerIDs[id] = len(q.listeners)
	q.listeners = append(q.listeners, listener)
	return id
}

func (q *queue) RemoveListener(id engine.ListenerID) bool {
	idx, present := q.listenerIDs[id]
	q.listeners = append(q.listeners[:idx], q.listeners[idx+1:]...)
	delete(q.listenerIDs, id)
	for id, i := range q.listenerIDs {
		if i > idx {
			q.listenerIDs[id]--
		}
	}
	return present
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
	q.listeners = make([]engine.Listener, 0)
	q.listenerIDs = make(map[engine.ListenerID]int)
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
		// Notifies registered listeners.
		for _, listener := range q.listeners {
			listener.Handle(ev)
		}
		// Removes listeners if destroyed.
		if ev.Verb() == engine.Combined {
			for _, ev := range ev.Subject().([]engine.Event) {
				q.maybeDestroyListener(ev)
			}
		} else {
			q.maybeDestroyListener(ev)
		}
	}
}

func (q *queue) maybeDestroyListener(ev engine.Event) {
	if ev.Verb() != engine.Destroy {
		return
	}
	listener, ok := ev.Subject().(engine.Listener)
	if !ok {
		return
	}
	for i, l := range q.listeners {
		if l == listener {
			q.listeners = append(q.listeners[:i], q.listeners[i+1:]...)
		}
	}
}

func (q *queue) nextListenerID() engine.ListenerID {
	q.idGen++
	return q.idGen
}
