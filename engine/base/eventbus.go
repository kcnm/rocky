package base

import (
	"sort"
)

type ListenerID int

type Listener interface {
	Handle(ev Event)
}

type event struct {
	instance Event
	cause    Event
}

type EventBus struct {
	events    []*event
	listeners map[ListenerID]Listener
	idGen     ListenerID
}

func NewEventBus() *EventBus {
	return &EventBus{
		make([]*event, 0),
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

func (bus *EventBus) Post(ev Event, cause Event) {
	if ev.Verb() == GameOver {
		bus.drain()
	} else {
		bus.events = append(bus.events, &event{ev, cause})
	}
}

func (bus *EventBus) PostAndTrigger(ev Event) {
	if len(bus.events) > 0 {
		panic("event bus is not empty")
	}
	bus.Post(ev, nil)
	for len(bus.events) > 0 {
		var e *event
		e, bus.events = bus.events[0], bus.events[1:]
		e.instance.Trigger()
		ids := make([]int, 0, len(bus.listeners))
		for id := range bus.listeners {
			ids = append(ids, int(id))
		}
		sort.Ints(ids)
		for _, id := range ids {
			if listener, present := bus.listeners[ListenerID(id)]; present {
				listener.Handle(e.instance)
			}
		}
		if combined, ok := e.instance.(CombinedEvent); ok {
			bus.reduce(combined)
		}
	}
}

func (bus *EventBus) nextListenerID() ListenerID {
	bus.idGen++
	return bus.idGen
}

func (bus *EventBus) drain() {
	bus.events = make([]*event, 0)
	bus.listeners = make(map[ListenerID]Listener)
}

func (bus *EventBus) reduce(ev CombinedEvent) {
	resultIndices := make([]int, 0, len(bus.events))
	for idx, e := range bus.events {
		if ev.HasEvent(e.cause) {
			resultIndices = append(resultIndices, idx)
		}
	}
	if len(resultIndices) > 1 {
		results := make([]Event, len(resultIndices))
		for i, idx := range resultIndices {
			results[i] =
				bus.events[idx-i].instance
			bus.events = append(bus.events[:idx-i], bus.events[idx-i+1:]...)
		}
		idx0 := resultIndices[0]
		bus.events = append(
			bus.events[:idx0],
			append(
				[]*event{&event{Combined(results...), ev}},
				bus.events[idx0:]...)...)
	}
}
