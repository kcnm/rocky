package event

import (
	"github.com/kcnm/rocky/engine"
)

type event struct {
	instance engine.Event
	cause    engine.Event
}

type bus struct {
	events      []*event
	listeners   []engine.Listener
	listenerIDs map[engine.ListenerID]int
	idGen       engine.ListenerID
}

func NewBus() engine.EventBus {
	return &bus{
		make([]*event, 0),
		make([]engine.Listener, 0),
		make(map[engine.ListenerID]int),
		0, // idGen
	}
}

func (b *bus) AddListener(listener engine.Listener) engine.ListenerID {
	id := b.nextListenerID()
	b.listenerIDs[id] = len(b.listeners)
	b.listeners = append(b.listeners, listener)
	return id
}

func (b *bus) RemoveListener(id engine.ListenerID) bool {
	idx, present := b.listenerIDs[id]
	b.listeners = append(b.listeners[:idx], b.listeners[idx+1:]...)
	delete(b.listenerIDs, id)
	for id, i := range b.listenerIDs {
		if i > idx {
			b.listenerIDs[id]--
		}
	}
	return present
}

func (b *bus) Post(ev engine.Event, cause engine.Event) {
	b.events = append(b.events, &event{ev, cause})
}

func (b *bus) PostAndTrigger(ev engine.Event) {
	if len(b.events) > 0 {
		panic("event bus is not empty")
	}
	b.Post(ev, nil)
	for len(b.events) > 0 {
		ev, b.events = b.events[0].instance, b.events[1:]
		mark := len(b.events)
		ev.Trigger()
		if ev.Verb() == engine.Combined && len(b.events) > mark {
			comb := &combined{make([]engine.Event, len(b.events)-mark)}
			for i := mark; i < len(b.events); i++ {
				comb.events[i-mark] = b.events[i].instance
			}
			b.events = append(b.events[:mark], &event{comb, ev})
		}
		for _, listener := range b.listeners {
			listener.Handle(ev)
		}
	}
}

func (b *bus) Drain() {
	b.events = make([]*event, 0)
	b.listeners = make([]engine.Listener, 0)
	b.listenerIDs = make(map[engine.ListenerID]int)
}

func (b *bus) nextListenerID() engine.ListenerID {
	b.idGen++
	return b.idGen
}