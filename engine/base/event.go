package base

import (
	"sort"
)

type Event interface {
	Subject() interface{}
	Verb() Verb

	Trigger()
}

type ListenerID int

type Listener interface {
	Handle(ev Event)
}

type EventBus struct {
	events    []Event
	listeners map[ListenerID]Listener
	idGen     ListenerID
}

func NewEventBus() *EventBus {
	return &EventBus{
		make([]Event, 0),
		make(map[ListenerID]Listener),
		0, // idGen
	}
}

func (bus *EventBus) AddListener(listener Listener) ListenerID {
	id := bus.nextListenerID()
	bus.listeners[id] = listener
	return id
}

func (bus *EventBus) RemoveListener(id ListenerID) bool {
	_, present := bus.listeners[id]
	delete(bus.listeners, id)
	return present
}

func (bus *EventBus) Post(ev Event) {
	if ev.Verb() == GameOver {
		bus.drain()
	} else {
		bus.events = append(bus.events, ev)
	}
}

func (bus *EventBus) PostAndTrigger(ev Event) {
	if len(bus.events) > 0 {
		panic("event bus is not empty")
	}
	bus.Post(ev)
	for len(bus.events) > 0 {
		ev, bus.events = bus.events[0], bus.events[1:]
		ev.Trigger()
		ids := make([]int, 0, len(bus.listeners))
		for id := range bus.listeners {
			ids = append(ids, int(id))
		}
		sort.Ints(ids)
		for _, id := range ids {
			if listener, present := bus.listeners[ListenerID(id)]; present {
				listener.Handle(ev)
			}
		}
	}
}

func (bus *EventBus) drain() {
	bus.events = make([]Event, 0)
	bus.listeners = make(map[ListenerID]Listener)
}

func (bus *EventBus) nextListenerID() ListenerID {
	bus.idGen++
	return bus.idGen
}
