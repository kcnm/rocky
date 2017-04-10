package event

import (
	"github.com/kcnm/rocky/engine"
)

type startTurn struct {
}

func StartTurn() engine.Event {
	return &startTurn{}
}

func (ev *startTurn) Subject() interface{} {
	return nil
}

func (ev *startTurn) Verb() engine.Verb {
	return engine.StartTurn
}

func (ev *startTurn) Trigger() {
}
