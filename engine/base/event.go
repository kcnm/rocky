package base

type Event interface {
	Subject() interface{}
	Verb() Verb

	Trigger()
}

type CombinedEvent interface {
	Event

	HasEvent(ev Event) bool
}

type combined struct {
	events []Event
}

func Combined(events ...Event) CombinedEvent {
	if len(events) <= 0 {
		panic("0 combined events")
	}
	v := events[0].Verb()
	for _, ev := range events {
		if ev.Verb() != v {
			panic("inconsistent verb")
		}
	}
	return &combined{events}
}

func (c *combined) Subject() interface{} {
	subjects := make([]interface{}, len(c.events))
	for i, ev := range c.events {
		subjects[i] = ev.Subject()
	}
	return subjects
}

func (c *combined) Verb() Verb {
	return c.events[0].Verb()
}

func (c *combined) Trigger() {
	for _, ev := range c.events {
		ev.Trigger()
	}
}

func (c *combined) HasEvent(ev Event) bool {
	for _, e := range c.events {
		if e == ev {
			return true
		}
	}
	return false
}
